# docker-training
Docker Training golang project

## Prerequisites
1. Docker
1. golang
1. `GOPATH` environment variable set to the directory in which this repo is cloned.

## Local Build
Assuming the repo is cloned to /home/dinu/docker-training
Set the `GOPATH` environment variable.
```
WINDOWS : set GOPATH=/home/dinu/docker-training
LINUX : export GOPATH=/home/dinu/docker-training
```
**To build the project locally**
1. Open terminal/command prompt in `$GOPATH/src/traning.com/server/`
1. Execute the below command to generate the executable.
   ```
   go build .
   ```
1. Run the executable. The application will be available at `http://localhost:8080/`

## Docker Build
Open terminal/command prompt in `$GOPATH/src/traning.com/server/`
**Build docker image**
```
docker build -f Dockerfile . -t dmathai/training
```
**Run docker image**
```
docker run -d -p 7777:8080 --name training dmathai/training
```
TThe application will be available at `http://localhost:7777/`

