package main

import (
	"fmt"
	"log"
)

type Logger interface {
	Info(args ...any)
	Fatal(args ...any)
}

type Log struct {
}

func (m *Log) Info(args ...any) {
	log.Printf(fmt.Sprintf("%v", args[0]), args[1:]...)
}

func (m *Log) Fatal(args ...any) {
	log.Fatalf(fmt.Sprintf("%v", args[0]), args[1:]...)
}

type MockLog struct {
}

func (m *MockLog) Info(args ...any) {
	// do nothing
}

func (m *MockLog) Fatal(args ...any) {
	// do nothing
}

type UserUC struct {
	logger Logger
}

func NewUserUC(logger Logger) *UserUC {
	return &UserUC{
		logger: logger,
	}
}

func (u *UserUC) GetUsername() string {
	u.logger.Info("initialize")

	return "Mamazo"
}

type Sender interface {
	Send(msg string) bool
}

type ApplePushNotif struct {
}

func (a *ApplePushNotif) Send(msg string) bool {
	fmt.Printf("message %v is sent to Apple Push notif\n", msg)
	return true
}

type FirebasePushNotif struct {
}

func (fb *FirebasePushNotif) Send(msg string) bool {
	fmt.Printf("message %v is sent to Firebase push notif\n", msg)
	return true
}

type SMSNotif struct {
}

func (fb *SMSNotif) Send(msg string) bool {
	fmt.Printf("message %v is sent via SMS\n", msg)
	return true
}

func main() {
	// Use Case 1: Dependency Injection

	mockedLogUser := NewUserUC(&MockLog{})
	// should print nothing
	// use mock in testing environment
	mockedLogUser.GetUsername()

	userUC := NewUserUC(&Log{})
	userUC.GetUsername() // should log

	// Use Case 2: Polymorphism
	// send message through many Push Notif service
	message := "Silakan check promo menarik ini!"
	pushNotifServices := []Sender{
		&ApplePushNotif{},
		&FirebasePushNotif{},
		&SMSNotif{},
	}

	for _, pushService := range pushNotifServices {
		pushService.Send(message)
	}
}
