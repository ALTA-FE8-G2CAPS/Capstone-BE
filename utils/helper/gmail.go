package helper

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"runtime"
	"time"

	// "go.starlark.net/lib/time"
	"gopkg.in/gomail.v2"
)

func SendGmailNotif(email, user, field, invoice, venue string, amount uint, qty, total int, cost uint) error {

	template, _ := filepath.Abs("./utils/helper/templates/notif-email.html")
	subject := "Payment Notification"
	templateData := BodyEmail{
		NAMA_USER: user,
		FIELD:     field,
		VENUE:     venue,
		INVOICE:   invoice,
		COST:      cost,
		QTY:       qty,
		AMOUNT:    amount,
		DATE:      time.Now().Format("2006-01-02"),
		TOTAL:     total,
	}
	result, errParse := ParseTemplate(template, templateData)
	fmt.Println(errParse)

	runtime.GOMAXPROCS(1)
	go SendEmail(email, subject, result)
	return nil
}

func SendEmail(to string, subject string, result string) error {
	const CONFIG_SMTP_HOST = "smtp.gmail.com"
	const CONFIG_SMTP_PORT = 587
	const CONFIG_SENDER_NAME = "Segoro App <segoroapp@gmail.com>"
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
	NAMA_USER string
	FIELD     string
	VENUE     string
	INVOICE   string
	ORDERID   int
	COST      uint
	QTY       int
	AMOUNT    uint
	TOTAL     int
	DATE      string
}
