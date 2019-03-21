package main

import (
	"fmt"
	"net/http"
	"os"
)

func HandleEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, os.Getenv("SMTP_HOST"))
}
