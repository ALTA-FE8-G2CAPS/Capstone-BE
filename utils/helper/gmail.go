package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/gomail.v2"
)

func SendGmailNotif(email, field, venue, cost, qty, amount, total, totalnotax, tax string) {
	template, _ := filepath.Abs("./utils/helper/templates/email-notif.html")
	subject := "Payment Notification"
	templateData := BodyEmail{
		FIELD:      field,
		VENUE:      venue,
		COST:       cost,
		QTY:        qty,
		AMOUNT:     amount,
		TOTAL:      total,
		TOTALNOTAX: totalnotax,
		TAX:        tax,
	}
	result, errParse := ParseTemplate(template, templateData)
	fmt.Println(errParse)

	runtime.GOMAXPROCS(1)
	go SendEmail(email, subject, result)
}

func SendEmail(to string, subject string, result string) error {
	const CONFIG_SMTP_HOST = "smtp.gmail.com"
	const CONFIG_SMTP_PORT = 587
	const CONFIG_SENDER_NAME = "Segoro App <muhammadadityogunawan@gmail.com>"
	CONFIG_AUTH_EMAIL := os.Getenv("EMAIL")
	CONFIG_AUTH_PASSWORD := os.Getenv("EMAIL_PASSWORD")
	m := gomail.NewMessage()
	m.SetHeader("From", CONFIG_SENDER_NAME)
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", result)

	d := gomail.NewDialer(
		CONFIG_SMTP_HOST, CONFIG_SMTP_PORT, CONFIG_AUTH_EMAIL, CONFIG_AUTH_PASSWORD)
	err := d.DialAndSend(m)
	if err != nil {
		panic(err)
	}
	return nil
}

func ParseTemplate(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type BodyEmail struct {
	FIELD      string
	VENUE      string
	COST       string
	QTY        string
	AMOUNT     string
	TOTAL      string
	TOTALNOTAX string
	TAX        string
}
