package main

import (
	"fmt"
	"net/http"

	"github.com/lestrrat/go-xslate"
)

func main() {
	http.HandleFunc("/hello", hello)

	http.ListenAndServe("127.0.0.1:8080", nil)
}

func hello(w http.ResponseWriter, req *http.Request) {
	tx, _ := xslate.New(xslate.Args{
		"Loader": xslate.Args{
			"LoadPaths": []string{"./templates"},
		},
		"Parser": xslate.Args{"Syntax": "TTerse"},
	})

	queryParams := req.URL.Query()

	req.ParseForm()
	bodyParams := req.PostForm
	fmt.Println(bodyParams)

	tx.RenderInto(w, "hello.tt", xslate.Vars{
		"name": queryParams["name"][0],
	})
}
