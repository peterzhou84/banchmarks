var express = require('express')
var app = express()
var fs = require('fs')
var mysql      = require('mysql');

var pool = mysql.createPool({
    host: "localhost",
    user: "benchagent",
    password: "benchagent1Q",
    database: "mysql",
    connectionLimit: 2
});

var redis = require("redis"),
client = redis.createClient();

client.on("error", function (err) {
    console.log("Error " + err);
});
 
app.get('/1', function (req, res) {
  res.send('Hello World')
})

app.get('/2', function (req, res) {
    fs.readFile('testfile.txt', (err, data) => {
        if (err) {res.status(501).end();return;}
            res.send(data + "")
        });
})
  
app.get('/3', function (req, res) {
    pool.query("select count(*) from mysql.user",function(err, rst){
        if (err) {res.status(501).end();return;}
        res.send('Hello World')
    });
})

app.get('/4', function (req, res) {
    client.get('key1',function(err, rst){
        if (err) {res.status(501).end();return;}
        res.send('Hello World')
    });
})
 
app.listen(3000)