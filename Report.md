# Description

To compare the performance among different program languages, we designed the following functionality scenarios.

1. Only respond 'hello world' for every request
2. respond the content in a local file, read the file content to the memory and respond
3. respond the content in the mysql database
4. respond the content in the redis cache

The performance test is carried out with Jmeter on two seperated machines.

Applications written in different languages are deployed on one machine. The JMeter is running on another.

The machines have 2vCPUs and 4G memory on cloud.

# Performance result Data

![](http://otn252ndm.bkt.clouddn.com/17-9-18/68105195.jpg)

The 1st column is the scenario name, combined with {language}_{concurrnetUsers}_{functionalScenario}

# Summary

1. Go's throughput is almost 3 times better than Nodejs.
2. The throughput is decreased when the concurrent user increased. Nodejs is decreasing slower than Go.
3. When the concurrent users comes to 400, the max response time is more than 15s for Nodejs. 7s when it hits 500 concurrent users. This means that, Go is much more stable than Nodejs in the high concurrency situation.
4. In scenario 3 (use mysql) and 4 (use redis), the system throuphput droped significantly. This maybe coused by the process capability of mysql and redis themselves. Go is still more stable and 