package main

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	EmailAddress string `json:"email"`
	Message      string `json:"message"`
}

func HandleEmail(w http.ResponseWriter, r *http.Request) {
	var mess Message
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&mess); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// smtpHost := os.Getenv("SMTP_HOST")
	//
	// emailAuth := smtp.PlainAuth("", os.Getenv("SMTP_LOGIN"), os.Getenv("SMTP_PASSWD"), smtpHost)
	//
	// to := []string{os.Getenv("EMAIL")}
	// msg := []byte(fmt.Sprintf("%s sent\r\n\r\n%s", mess.EmailAddress, mess.Message))

	// if err := smtp.SendMail(smtpHost+os.Getenv("SMTP_PORT"), emailAuth, mess.EmailAddress, to, msg); err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("success"))
}
