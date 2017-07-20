# Cntich
Cnitch is a framework for monitoring processes inside of Docker to identify any processes which are running as route. Cnitch will monitor processes running under the docker engine every 5s and will alert when a process in a container is found to be running as root.


Currently WIP.

## How to run
* Set environment variable DOCKER_HOST pointed at your docker engine
* execute `go run main.go`
* run some containers

# Roadmap:
* Library which can be implemented as a plugin into applications
* Docker container
* Pluggable alerting framework
