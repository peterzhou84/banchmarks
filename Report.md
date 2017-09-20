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

Addtionally, compared with PHP benchmark data [PHP Benchmarks](https://github.com/kenjis/php-framework-benchmark), Nodejs is 2.55 times faster than PHP siler. If we compare Nodejs with Yii2, it's 12.89 times faster. PHP without template engine should be faster than this.

# Analysis

During the performance execution, the CPU cost is almost even, abount 50%. The difference is that, system core's cpu usage is about 16% for Go, but for Nodejs, it's 5%. I think Go is more optimized to use system core. On the other hand, Nodejs is mainly running on the user mode. 

After looking at the performance data between different languages on different operations, Go and Nodejs have their own merits. For example, Nodejs is good at base64, json proccesing, while Go is good at matmul processing. Here is the original [performance data](https://github.com/kostya/benchmarks).

Moreover, for engineering purpose, nodejs has complete solution of how a project could be hold in a private repository and it's very easy to manage the project dependencies. 

On the other hand, Go project could be hold in a private git repository too, and the dependency management is easy.

With the powerful authority management provided by gitlab, go projects could have a very flexible structure. You can define that some project could be used by specific user group, or by logined users, even anonymous users.

On the other hand, Nodejs private project managed by sinopia doesn't support user group and the configuration is based on config file.

So, I think Go is a better choice for big oganizations from engineering perspective.

