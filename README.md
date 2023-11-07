# Leostream-admin-cli

This repository contains a command line interface (CLI) for the Leostream Connection Broker REST API.  The CLI is written in Go and is compiled to a single binary that can be run on any platform that supports Go. Currently it can only fetch
- centers
- pools
- gateways

## Installation

To install the CLI, you must have Go installed on your system.  Once Go is installed, you can install the CLI by running the following command:

```bash
go install github.com/joustie/leostream-admin-cli
```

## Usage

### Add environment variables

Your user in Leostream needs permission to access the REST API.  You can create a user with the appropriate permissions in the Leostream Connection Broker web interface.  Once you have created a user, you can set the following environment variables:

``` 
export LEOSTREAM_API_USERNAME=<username>
export LEOSTREAM_API_PASSWORD=<password>
export LEOSTREAM_API_HOST=<host>
```

### List centers

```
leostream-admin-cli list-centers
```

### List pools

```
leostream-admin-cli list-pools
```

### List gateways

```
leostream-admin-cli list-gateways
```
