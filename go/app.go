package main

import (
    "bufio"
	"io"
	"net/http"
	"os"
	"bytes"
	"strings"
	"fmt"
	"github.com/pubnative/mysqldriver-go"
)

db := mysqldriver.NewDB("root@tcp(localhost:3306)/mysql", 2)

func strHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func fileHandler(w http.ResponseWriter, r *http.Request) {
	fh, ferr := os.Open("./testfile.txt")
    if ferr != nil {
		fmt.Println(ferr)
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
	// obtain connection from the pool
	conn, err := db.GetConn()
	if err != nil {
		panic(err)
	}
	rows, err := conn.Query(`select count(*) from mysql.user`)
	if err != nil {
		panic(err)
	}
	if err := db.PutConn(conn); err != nil {
		panic(err)
	}
	io.WriteString(w, "Hello World")
}
func redisHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/1", strHandler)
	mux.HandleFunc("/2", fileHandler)
	mux.HandleFunc("/3", dbHandler)
	mux.HandleFunc("/4", redisHandler)

	http.ListenAndServe(":3000", mux)
}