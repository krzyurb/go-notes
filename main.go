package main

import (
	"notesapp/loaders"
)

func main() {
	config := loaders.BuildConfig()
	loaders.DBConnect(config)
	loaders.BuildServer(loaders.BuildRouter()).Run()
}