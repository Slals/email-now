package main

import (
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/haisum/recaptcha"
	gomail "gopkg.in/gomail.v2"
)

type Message struct {
	EmailAddress string `json:"email"`
	Message      string `json:"message"`
	GreToken     string `json:"greToken"`
}

func HandleEmail(w http.ResponseWriter, r *http.Request) {
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}

	re := recaptcha.R{
		Secret: os.Getenv("RE_SECRET"),
	}

	var mess Message
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&mess); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid message"))
		return
	}

	if !re.VerifyResponse(mess.GreToken) {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid Re-Captcha"))
		return
	}

	m := gomail.NewMessage()
	m.SetHeader("From", mess.EmailAddress)
	m.SetHeader("To", os.Getenv("EMAIL"))
	m.SetHeader("Subject", "New message")
	m.SetBody("text/plain", mess.Message)

	port, _ := strconv.ParseInt(os.Getenv("SMTP_PORT"), 10, 32)
	d := gomail.NewDialer(os.Getenv("SMTP_HOST"), int(port), os.Getenv("SMTP_LOGIN"), os.Getenv("SMTP_PASSWD"))

	if err := d.DialAndSend(m); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Failed to send"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Successfuly sent"))
}
