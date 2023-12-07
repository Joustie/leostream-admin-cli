# Leostream-admin-cli

This repository contains a command line interface (CLI) for the Leostream Connection Broker REST API.  The CLI is written in Go and is compiled to a single binary that can be run on any platform that supports Go. Currently it can only fetch
- pool assignments


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
task: [build] GOFLAGS=-mod=mod go build -o bin/leostream-admin-cli main.go
```

## Usage

### Add environment variables

Your user in Leostream needs permission to access the REST API.  You can create a user with the appropriate permissions in the Leostream Connection Broker web interface.  Once you have created a user, you can set the following environment variables:

``` 
export LEOSTREAM_API_USERNAME=<username>
export LEOSTREAM_API_PASSWORD=<password>
export LEOSTREAM_API_API_URL=https://<api host>/rest/v1
```



### List pools-assignments

```bash
% leostream-admin-cli poolassignments list --policy 8
{
        "id": 29,
        "pool_name": "dev-pool",
        "pool_id": 13
},
{
        "id": 30,
        "pool_name": "acc-pool",
        "pool_id": 14
},
{
        "id": 31,
        "pool_name": "prd-pool",
        "pool_id": 15
}
....
```
