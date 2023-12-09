# Leostream-admin-cli

This repository contains a command line interface (CLI) for the Leostream Connection Broker REST API.  The CLI is written in Go and is compiled to a single binary that can be run on any platform that supports Go. Currently it can fetch and update the following Leostream resources:

- pools
- centers
- gateways
- pool assignments

Additionally, it can create the following resources:
- pool assignments

*This is a work in progress*

## TODO
- add create command for
        - pools
        - gateways
- add delete command for all resources
- add more tests
- add more error handling
- add more documentation  how to use the CLI
- add more examples
- add more output formats


## Installation

To install the CLI, you must have Go installed on your system. Please see the [Go installation instructions](https://golang.org/doc/install) for more information.

Next, you must install the Leostream client library:
```
go get github.com/joustie/leostream-client-go@latest
```

Once the library is installed, you can install the CLI by running the following command:

```bash
go install github.com/joustie/leostream-admin-cli
```

The binary will then be installed in your GOPATH/bin/ directory.  You can add this directory to your PATH to make the CLI available from anywhere.

In my case, I have the following in my ~/.bash_profile:

```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## Building

### Install go-task (MacOS)

```bash
% brew install go-task/tap/go-task
```

### Build the binary using task
```bash
% task build
task: [build] GOFLAGS=-mod=mod go build -o $GOPATH/bin/leostream-admin-cli main.go
```

## Usage

### Add environment variables

Your user in Leostream needs permission to access the REST API.  You can create a user with the appropriate permissions in the Leostream Connection Broker web interface.  Once you have created a user, you can set the following environment variables:

``` 
export LEOSTREAM_API_USERNAME=<username>
export LEOSTREAM_API_PASSWORD=<password>
export LEOSTREAM_API_API_URL=https://<api host>/rest/v1
```

### Execute a command

Coming soon...
