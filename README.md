# cnitch

[CircleCI](https://circleci.com/gh/nicholasjackson/cnitch)
[GoDoc](https://godoc.org/github.com/nicholasjackson/cnitch)
[Docker Repository on Quay](https://quay.io/repository/nicholasjackson/cnitch)

cnitch (snitch or container snitch) is a simple framework and command line tool for monitoring Docker containers to identify any processes which are running as root.  

Why is this a bad thing?  If you have not already been to [can I haz non-privileged containers? by mhausenblas](http://canihaznonprivilegedcontainers.info) then I recommend you head over there now to get all the info.


## How it works
cnitch connects to the Docker Engine using the API and queries the currently running containers,  it then inspects the processes running inside this container and identifies any which are running as the root user.  
When a root process is found this information is sent to the configurable reporting modules allowing you to audit or take action on this information.

```bash
2017/07/29 16:04:27 Starting Cnitch: Monitoring Docker Processes at: tcp://172.16.255.128:2376
2017/07/29 16:04:27 Checking for root processes every: 10s
2017/07/29 16:05:08 Checking image: ubuntu, id: 7bd489560a310343c39186500daa680290289c27f7a730524a31355a3aaf0430
2017/07/29 16:05:08 >> WARNING: found process running as root: tail -f /dev/null pid: 365
```

## Reporting Modules
At present cnitch has the capability of reporting to **StatsD** and **StdOut**.  Reporting backends are extensible to make it easy to support any backend, for example it would be a fairly trivial process to build a backend to support log stash or another log file aggregation tool.

### StatsD
The exceptions are sent to the statsD endpoint as a count using the `cnitch.exception.root_process` metric.  The metrics are also tagged with the `host` name of the cnitch instance and the `container` name.

### StdOut
The StdOut logger is a simple output logger which sends the reported exceptions to StdOut. 

## How to run
Wether you run cnitch in a Docker container or if you run it as a binary it needs access to the Docker api by setting the URL of the server or the path to the socket with the environment variable **DOCKER_HOST**

### Flags
* `--hostname=[hostname]` the name or ip address to be used for metric aggregation
* `--statsd-server=[hostname:port]` the URI of the statsd collector, if omitted statsd reporting will be disabled
* `--check=[duration  e.g. 10s (10 seconds), 1m (1 minute)]` , the check frequency that snitch will scan for root processes

### Command Line
Set environment variable **DOCKER_HOST** to your docker engine API then run snitch with the required flags.

```bash
$ cnitch --hostname=myhost --statsd-server=127.0.0.1:8125 --check=10s
```

### Docker
cnitch runs in a non privileged container and if you wish to use the Docker sock for access to the API you need to add the cnitch user to the `docker` group.   This can be achieved through the flag `--group-add` , set this to the group id for the docker user group.  
 For example:

 `--group-add=$(stat -f "%g" /var/run/docker.sock`

*Example using the Docker sock file for API access*

```bash
$ docker run -i -t --rm \
  -v /var/run/docker.sock:/var/run/docker.sock \
  --group-add=$(stat -f "%g" /var/run/docker.sock) \
  -e "DOCKER_HOST:unix:///var/run/docker.sock" \
  quay.io/nicholasjackson/cnitch [options]
```

If you are running on a mac and using Docker Machine the Docker sock is inside the VM which means you can not use the `stat` command to discover the group id.

## Example
There is an example Docker Compose stack inside the [./example](/example) folder to show how cnitch exports data to statsd.  To run this example start the example run:

```
$ cd ./example
$ docker-compose up
``` 

Once everything has started running open `http://[docker host ip]:3000` in your web browser and you should see the Grafana login screen.

![grafana login](./screen1.png)

Log in to Grafana using the following credentials:
* user: admin
* password: admin

Then select the cnitch dashboard.  This dashboard shows the current running root processes.  

![root processes chart](./screen2.png)

If you are not using `/var/run/docker.sock` to communicate with your Docker host then you will need to change some of the settings inside the file `./example/docker-compose.yml` to match your settings.

## Roadmap
[ ] Implement TLS client certificates for Docker API
