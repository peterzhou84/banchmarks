# benchmarks
Performance benchmarks for common languages, PHP, NODEJS, GO and so on

# Concerns

* We don't know how big the performance difference is between common languages
* We don't know weather the common operation effects the performance or not


# main functional senarios

All request uses get method.

1. Only respond 'hello world' for every request
2. respond the content in a local file, read the file content to the memory and respond
3. respond the content in the mysql database
4. respond the content in the redis cache

# basic design

## request path

`/n`refer to the different senario. For example, `/1` will call the first functional senario.

## performance testing tool

JMeter on the local machine

## testing senarios

1. 100 concurrent users
2. 200 concurrent users
3. 300 concurrent users
4. 400 concurrent users
5. 500 concurrent users

## data recording

1. Data recorded by JMeter
2. CPU 
3. MEM

# Environment Preparation

## db

    grant select on mysql.user to 'banchagent'@'%' identified by 'banchagent';
    flush privileges;