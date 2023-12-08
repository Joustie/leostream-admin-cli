/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	 "github.com/joustie/go-leostream-admin-cli/cmd"
	 _ "github.com/joustie/go-leostream-admin-cli/cmd/gateway"
	 _ "github.com/joustie/go-leostream-admin-cli/cmd/center"
	 _ "github.com/joustie/go-leostream-admin-cli/cmd/pool"
	 _ "github.com/joustie/go-leostream-admin-cli/cmd/poolassignment"
)

func main() {
	cmd.Execute()
}
