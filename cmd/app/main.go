package main

import (
	"main/server"
	"main/server/common/globals"
)


func main() {
	globals.SetupEnvironmentVariables()
	server.Run()
}