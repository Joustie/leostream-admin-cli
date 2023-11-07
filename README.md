# Leostream-admin-cli

This repository contains a command line interface (CLI) for the Leostream Connection Broker REST API.  The CLI is written in Go and is compiled to a single binary that can be run on any platform that supports Go. Currently it can only fetch
- centers
- pools
- gateways

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
% leostream-admin-cli list-centers
{
        "id": 1,
        "name": "Vsphere",
        "flavor": "I",
        "status": 2,
        "status_label": "Offline",
        "type": "vcenter",
        "type_label": "VMware vSphere and vCenter Server"
}
...
```

### List pools

```
% leostream-admin-cli list-pools
{
        "id": 1,
        "name": "All Desktops",
        "pool_definition": {
                "never_rogue": 0,
                "use_vmotion": 0
        },
        "provision": {
                "provision_on_off": 0,
                "provision_threshold": 0,
                "provision_max": 0,
                "center": {}
        }
}
....
```

### List gateways

```
% leostream-admin-cli list-gateways
{
        "id": 1,
        "name": "gateway_1",
        "address": "gateway.hocmodo.nl",
        "address_private": "10.0.0.1",
        "load_balancer_id": 0,
        "use_src_ip": 0,
        "notes": ""
}
```
