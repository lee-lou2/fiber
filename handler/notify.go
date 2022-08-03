package handler

import (
	"car/conf"
	"fmt"
	twilioApi "github.com/kevinburke/twilio-go"
	"gopkg.in/gomail.v2"
	"log"
	"net/url"
	"strconv"
)

func NotifySendEmail(toEmail, subject, message string) {
	// 이메일 발송
	go notifySendSMTPEmail(toEmail, subject, message)
}

func NotifySendSMS(toNumber, message string) {
	// 문자 발송
	go notifySendTwilioSMS(toNumber, message)
}

func NotifySendCall(toNumber string, callUrl *url.URL) {
	// 전화 발신
	go notifySendTwilioCall(toNumber, callUrl)
}

func notifySendSMTPEmail(toEmail, subject, message string) {
	// 기본 변수
	fromEmail := conf.Config("DEFAULT_FROM_EMAIL")
	emailUser := conf.Config("EMAIL_HOST_USER")
	emailPassword := conf.Config("EMAIL_HOST_PASSWORD")
	emailSMTPHost := conf.Config("EMAIL_HOST")
	emailSMTPPortString := conf.Config("EMAIL_PORT")
	emailSMTPPort, _ := strconv.Atoi(emailSMTPPortString)

	// 이메일 발송
	m := gomail.NewMessage()
	m.SetHeader("From", fromEmail)
	m.SetHeader("To", toEmail)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", message)

	d := gomail.NewDialer(emailSMTPHost, emailSMTPPort, emailUser, emailPassword)

	if err := d.DialAndSend(m); err != nil {
		log.Printf("이메일 발송 실패 : %s\n", toEmail)
		panic(err)
	}
	log.Printf("이메일 발송 성공 : %s\n", toEmail)
}

func notifySendTwilioSMS(toNumber string, message string) {
	/*
		문자 발송
	*/
	fromNumber := conf.Config("TWILIO_API_FROM_NUMBER")
	sid := conf.Config("TWILIO_API_SID")
	token := conf.Config("TWILIO_API_AUTH_TOKEN")

	client := twilioApi.NewClient(sid, token, nil)

	msg, err := client.Messages.SendMessage(
		fromNumber,
		fmt.Sprintf("+82%s", toNumber[1:]),
		message,
		nil,
	)
	if err != nil {
		log.Printf("문자 발송 실패 : %s", err)
		panic(err)
	}
	log.Printf("문자 발송 성공 : %s, %s", msg.Sid, msg.FriendlyPrice())
}

func notifySendTwilioCall(toNumber string, callUrl *url.URL) {
	/*
		전화 발신
	*/
	fromNumber := conf.Config("TWILIO_API_FROM_NUMBER")
	sid := conf.Config("TWILIO_API_SID")
	token := conf.Config("TWILIO_API_AUTH_TOKEN")

	client := twilioApi.NewClient(sid, token, nil)

	call, err := client.Calls.MakeCall(fromNumber, toNumber, callUrl)

	if err != nil {
		log.Printf("전화 발신 실패 : %s", err)
		panic(err)
	}
	log.Printf("전화 발신 성공 : %s, %s", call.Sid, call.FriendlyPrice())
}
