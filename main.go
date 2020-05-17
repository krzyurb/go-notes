package main

import (
	"notesapp/http"
)

func main() {
	http.BuildServer(http.BuildRouter()).Run()
}