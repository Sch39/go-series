package contracts

import "fmt"

type Notifier interface {
	Notify(message string) string
}

type EmailNotifier struct {
	Address string
}

func (e EmailNotifier) Notify(message string) string {
	return fmt.Sprintf("email to %s: %s", e.Address, message)
}

var _ Notifier = EmailNotifier{}

type AlertService struct{}

func (AlertService) Send(notifier Notifier, message string) string {
	return notifier.Notify(message)
}

type MessageSender interface {
	Notify(message string) string
}

func Deliver(sender MessageSender, message string) string {
	return sender.Notify(message)
}

type PushNotifier struct {
	Token string
}

func (p *PushNotifier) Notify(message string) string {
	return fmt.Sprintf("push to %s: %s", p.Token, message)
}

var _ Notifier = (*PushNotifier)(nil)

type StatusError struct {
	Code int
}

func (e *StatusError) Error() string {
	return fmt.Sprintf("status %d", e.Code)
}

func LoadStatus(ok bool) error {
	if ok {
		return nil
	}
	return &StatusError{Code: 503}
}
