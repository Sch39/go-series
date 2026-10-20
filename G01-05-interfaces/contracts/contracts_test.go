package contracts_test

import (
	"testing"

	"github.com/sch39/go-series/g0105/interfaces/contracts"
)

type recordingNotifier struct {
	message string
}

func (r *recordingNotifier) Notify(message string) string {
	r.message = message
	return "recorded: " + message
}

func TestImplicitImplementationAndConsumerInterface(t *testing.T) {
	service := contracts.AlertService{}
	email := contracts.EmailNotifier{Address: "dev@example.com"}

	gotEmail := service.Send(email, "build failed")
	recorder := &recordingNotifier{}
	gotRecord := contracts.Deliver(recorder, "deploy ready")

	t.Logf("email=%q", gotEmail)
	t.Logf("record=%q captured=%q", gotRecord, recorder.message)

	if gotEmail != "email to dev@example.com: build failed" {
		t.Errorf("Send() = %q, want email notification", gotEmail)
	}
	if recorder.message != "deploy ready" {
		t.Errorf("Deliver() captured %q, want %q", recorder.message, "deploy ready")
	}
}

func TestPointerReceiverMethodSet(t *testing.T) {
	service := contracts.AlertService{}
	push := &contracts.PushNotifier{Token: "device-42"}

	got := service.Send(push, "maintenance")
	t.Logf("notification=%q concrete=%T", got, push)

	if got != "push to device-42: maintenance" {
		t.Errorf("Send() = %q, want push notification", got)
	}
}

func TestNilInterface(t *testing.T) {
	var err error
	okErr := contracts.LoadStatus(true)
	failedErr := contracts.LoadStatus(false)

	t.Logf("zero interface: value=%v type=%T nil=%t", err, err, err == nil)
	t.Logf("success: value=%v type=%T nil=%t", okErr, okErr, okErr == nil)
	t.Logf("failure: value=%v type=%T nil=%t", failedErr, failedErr, failedErr == nil)

	if err != nil || okErr != nil {
		t.Errorf("nil errors must compare equal to nil: zero=%v success=%v", err, okErr)
	}
	if failedErr == nil {
		t.Error("LoadStatus(false) returned nil error")
	}
}
