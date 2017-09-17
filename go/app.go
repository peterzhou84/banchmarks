package main
/**
 * mysql链接的时候一开始采用github.com/pubnative/mysqldriver-go库。
 * 但是，执行的时候总是报库内部的越界错误，虽然查询结构正确，但是还是弃用了。
 */
import (
    "bufio"
	"io"
	"net/http"
	"os"
	"bytes"
	"strings"
	"time"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/garyburd/redigo/redis"
)


var db *sqlx.DB;
var redispool *redis.Pool

func newRedisPool(addr string) *redis.Pool {
	return &redis.Pool{
	  MaxIdle: 3,
	  IdleTimeout: 240 * time.Second,
	  Dial: func () (redis.Conn, error) { return redis.Dial("tcp", addr) },
	}
  }
  

func strHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	fh, ferr := os.Open("./testfile.txt")
    if ferr != nil {
		fmt.Println(ferr)
		w.WriteHeader(501)
        return 
    }
	defer fh.Close()
	var buffer bytes.Buffer
    inputread := bufio.NewReader(fh)
    for {
        input, ferr := inputread.ReadString('\n')
		buffer.WriteString(strings.TrimSpace(input))
        if ferr == io.EOF {
            break
		}
	}
	io.WriteString(w, buffer.String())
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`select count(*) from mysql.user`)
	defer rows.Close();
	rows.Next()
	var count int
	rows.Scan((&count))
	// fmt.Println(count)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(501)
		return
	}

	io.WriteString(w, "Hello World")
}
func redisHandler(w http.ResponseWriter, r *http.Request) {
	conn := redispool.Get()
	defer conn.Close()
	n, err := conn.Do("GET", "key1")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(501)
        return 
	}
	if n != nil {;}
	// fmt.Println(n)
	io.WriteString(w, "Hello World")
}


func init(){
	// 初始化数据库连接池
	database, err := sqlx.Open("mysql", "benchagent:benchagent1Q#@tcp(localhost:3306)/mysql")
	database.SetMaxIdleConns(2)
	database.SetMaxOpenConns(2)
	if err != nil {
        fmt.Println("open mysql failed,", err)
        return
    }

	db = database
	
	// 初始化redis连接池
	redispool = newRedisPool("localhost:6379")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/1", strHandler)
	mux.HandleFunc("/2", fileHandler)
	mux.HandleFunc("/3", dbHandler)
	mux.HandleFunc("/4", redisHandler)

	http.ListenAndServe(":3000", mux)
}