package test

import (
	"fmt"
	"net/http"

	"code.google.com/p/buildtagstest/pkg"
)

var isGo11 bool

func init() {
	http.HandleFunc("/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "<p>Go11 is ", isGo11)
			fmt.Fprintln(w, "<p>pkg.Go11 is ", pkg.IsGo11)
		})
}
