package main

import (
	"notesapp/http"
)

func main() {
	http.BuildServer(8080, http.BuildRouter()).Run()
}
