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

func Data(values ...interface{}) (map[interface{}]interface{}, error) {
	if len(values)%2 != 1 {
		return nil, errors.New("invalid data call")
	}
	data := make(map[interface{}]interface{}, len(values)/2)
	for i := 0; i < len(values)-1; i += 2 {
		if key, ok := values[i].(string); ok {
			data[key] = values[i+1]
		} else {
			return nil, errors.New("data keys must be strings")
		}
	}

	last := values[len(values)-1]
	if reflect.ValueOf(last).Kind() == reflect.Map {
		for k, v := range last.(map[interface{}]interface{}) {
			data[k] = v
		}
	} else {
		return nil, errors.New("data last elem must be map[interface{}]interface{}")
	}

	return data, nil
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
