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
- add delete command
        - pools
        - gateways
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

### Step 1: Add environment variables

Your user in Leostream needs permission to access the REST API.  You can create a user with the appropriate permissions in the Leostream Connection Broker web interface.  Once you have created a user, you can set the following environment variables:

``` 
export LEOSTREAM_API_USERNAME=<username>
export LEOSTREAM_API_PASSWORD=<password>
export LEOSTREAM_API_API_URL=https://<api host>/rest/v1
```

### Step 2: Execute a command

To execute a command, you must specify the resource type and the command like so:
```bash
leostream-admin-cli <resource type> <command> <options>
```

Help is available for each command by running the command with the --help flag. For example, to get help for the pool get command, you would run the following command:


```bash
%  leostream-admin-cli pool get --help 
Get Leostream pool by ID. For example:
		leostream-admin-cli pool get --pool_id 1

Usage:
  leostream-admin-cli pool get [flags]

Flags:
  -h, --help   help for get

Global Flags:
      --pool_id string   Pool_id to get pool for
```



Not all commands have flags. For example, to list all pools, you would run the following command:

```bash 
% leostream-admin-cli pool list                                              
{
	"id": 1,
	"name": "All Desktops"
}
{
	"id": 4,
	"name": "All Linux Desktops"
}
{
	"id": 3,
	"name": "All Windows Desktops"
}
{
	"id": 11,
	"name": "Linux Advanced Desktop Pool"
}
{
	"id": 103,
	"name": "My Pool2"
}
....
```

To get a specific pool, you would run the following command:

```leostream-admin-cli pool get --pool_id <pool id>```

E.g.
```bash
% leostream-admin-cli pool get --pool_id 103
leostream-admin-cli pool get --pool_id 103
{
	"id": 103,
	"name": "My Pool2",
	"display_name": "My Pool",
	"pool_definition": {
		"restrict_by": "C",
		"never_rogue": 0,
		"server_ids": [
			1
		],
		"use_vmotion": 0
	},
	"provision": {
		"provision_on_off": 0,
		"provision_threshold": 1,
		"provision_max": 1,
		"center": {}
	}
}

```

Updating a pool is similar to getting a pool, except you must add the appropriate flags. The help function can be used to see what flags are available for each command. For example, to get help for the pool update command, you would run the following command:

```bash 
% leostream-admin-cli pool update --help
Update Leostream pool by ID. For example:
		leostream-admin-cli pool update --pool_id 1 --name "My Pool" --display_name "My Pool" --provision_onoff 1 --provision_threshold 1 --provision_vm_name test1 --provision_vm_id 1 --provision_name_next_value 1 --provision_max 1 --center_id 1 --center_type "vmware"

Usage:
  leostream-admin-cli pool update [flags]

Flags:
      --center_id string                   ID of center to which this pool belongs
      --center_type string                 Type of center to which this pool belongs (amazon, azure, vmware, etc.
      --display_name string                Pool display name in the UI
  -h, --help                               help for update
      --name string                        Pool name
      --provision_max string               Maximum number of provisioned desktops (default "1")
      --provision_name_next_value string   Next value to use for naming provisioned desktops (default "1")
      --provision_onoff string             Desktop provisioning enabled or not
      --provision_threshold string         Threshold at which to provision desktops
      --provision_vm_id string             ID of the desktop image to use for provisioning
      --provision_vm_name string           Name of the desktop instance to be provisioned (default "1")

Global Flags:
      --pool_id string   Pool_id to get pool for
```

To update a pool for instance to change the display name, you would run the following command:

```bash
%  leostream-admin-cli pool update --pool_id 103 --name "My Pool" --display_name "My new Pool" --provision_onoff 0 --provision_threshold 1 --provision_vm_name test1 --provision_vm_id 1 --provision_name_next_value 1 --provision_max 1
{
	"stored_data": {
		"id": 103,
		"name": "My Pool",
		"display_name": "My new Pool",
		"pool_definition": {
			"restrict_by": "C",
			"never_rogue": 0,
			"server_ids": [
				1
			],
			"use_vmotion": 0
		},
		"provision": {
			"provision_on_off": 0,
			"provision_threshold": 1,
			"provision_max": 1,
			"center": {}
		}
	}
}
```
The response contains the updated pool exactly as it is stored in the Leostream Connection Broker database.
