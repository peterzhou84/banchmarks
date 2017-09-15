# We use jmeter to do the performance testing

# steps

## install jmeter

download jmeter from apache http://jmeter.apache.org/download_jmeter.cgi

Unzip to your local directory and it's done.

> note: jmeter could be added to the path of the machine

## install the java jvm

The version of java required to be later than 1.8.0.

## run test

    jmeter -n -t performance.jmx -Jhost=192.168.1.3 -Jport=3000 -Juses=100 -Jpath=1 -Jsecs=600 -Jlanguage=nodejs

Here is an explaination of the parameters:

1. `host` `port` are the hostname and the port of the service to be invoked. Default as `192.168.8.3` and `3000`
2. `users` is the concurrent user number. Default as 100.
3. `path` is the get path to be requested, not include the root `/`. In this banchmark project, it's 1, 2, 3 or 4. Default as `1`
4. `secs` for how many seconds this test could last, the default is `600`.
5. `language` is the coding language used for the service. Default as `nodejs`.

## result collection

For each test execution, there should be one data collection file generated on the execution directory. The name of the file should be ${language}_${users}_${path}, like `nodejs_100_1.log`

Later, you can import the data into the jmeter 'Aggregate Graph' to get the statistic data and graph.

All the different parapeter should generate different log file. If they are combined in one file, the graph should be working fine too.

## Tips

1. Please make sure your machine time is up to date
2. If you come to unknown error, please add `-LDEBUG` the the command line parameter, this will add debug information to the log file.