package api

import "fmt"

const (
	// Status code for VPN TransactionStatus
	SuccessTransaction             = "00"
	NotCompleteTransaction         = "01"
	ErrorTransaction               = "02"
	ReverseTransaction             = "04" // Customer is subtracted from bank but transaction not success at VNPAY
	ProcessingTransaction          = "05" // Transaction refund money
	RequestRefundMoney             = "06" // VNPay request refund to bank
	SuspectedFraudulentTransaction = "07"
	RefusedRefundTransaction       = "09"
)

type HTTPError struct {
	StatusCode string
	Message    string
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("code: %v, message: %v", e.StatusCode, e.Message)
}

func NewHTTPError(statusCode, message string) *HTTPError {
	return &HTTPError{
		StatusCode: statusCode,
		Message:    message,
	}
}
