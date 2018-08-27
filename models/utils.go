package models

import (
	"errors"
	"net/smtp"
	"os"
	"reflect"
)

func HasElem(slice interface{}, elem interface{}) bool {
	data := reflect.ValueOf(slice)

	if data.Kind() == reflect.Slice {
		for i := 0; i < data.Len(); i++ {
			if data.Index(i).Interface() == elem {
				return true
			}
		}
	}

	return false
}

func LookupEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func SendEmail(from string, to string, subject string, body string) error {
	if len(body) == 0 {
		return errors.New("body must not be empty")
	}

	if from == "" {
		from = "unknown@gmail.com"
	}

	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n\n" + body

	address := LookupEnv("smtp.address", "smtp.gmail.com:587")
	host := LookupEnv("smtp.host", "smtp.gmail.com")

	user := LookupEnv("smtp.user", "username")
	password := LookupEnv("smtp.password", "password")

	return smtp.SendMail(address, smtp.PlainAuth("", user, password, host), from, []string{to}, []byte(msg))
}
