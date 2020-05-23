package main

import (
	"notesapp/db"
	"notesapp/loaders"
)

func main() {
	config := loaders.BuildConfig()
	db.BuildConnection(config)
	loaders.BuildServer(loaders.BuildRouter()).Run()
}