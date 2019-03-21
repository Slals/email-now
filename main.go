package main

import (
	"fmt"
	"net/http"
)

func HandleEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hey")
}
