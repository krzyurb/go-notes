package main

import (
	"notesapp/loaders"
)

func main() {
	loaders.DBConnect()
	loaders.BuildServer(loaders.BuildRouter()).Run()
}