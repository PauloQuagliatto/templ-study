package main

import (
	"fmt"
	"net/http"

  "github.com/a-h/templ"
	"github.com/PauloQuagliatto/templ-study/hello_templ"
)

func main() {
	component := headerTemplate("macaco")
	
	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	http.ListenAndServe(":3000", nil)
}
