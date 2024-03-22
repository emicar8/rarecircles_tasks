package task1

import "fmt"

var Cap *Capture

// Task2 implementation of the “adapter” structural design pattern
// This task simulates a use case where a payment is usually created based on a authoriation. Even though that is the normal logic
// there are border cases where the payments must be created from a capture. To cover that border case, an adapter is implemented to allow
// a capture to be passed to the process payment function.
func Task1() string {
	client := &Client{}

	capAdapter := &CaptureAdapter{Capture: Cap}

	return client.ProcessPayment(capAdapter)
}

type Client struct{}

func (c *Client) ProcessPayment(payment Payment) string {
	return "approved amount: " + payment.GeneratePayment()
}

type Payment interface {
	GeneratePayment() string
}

type Authorization struct {
	AuthorizedAmount float64
}

func (auth *Authorization) GeneratePayment() string {
	return fmt.Sprintf("%f", auth.AuthorizedAmount)
}

type Capture struct {
	CaptureAmount float64
}

type CaptureAdapter struct {
	Capture *Capture
}

func (captureAdapter *CaptureAdapter) GeneratePayment() string {
	return fmt.Sprintf("%f", captureAdapter.Capture.CaptureAmount)
}
