package api

import "github.com/gin-gonic/gin"

const (
	VnpVersion    = "2.1.0"
	VnpCommand    = "pay"
	VNLocale      = "vn"
	VnpURL        = "https://sandbox.vnpayment.vn/paymentv2/vpcpay.html"
	VnpTmnCode    = "7ONWQPBC"
	VnpReturnURL  = "https://domain.vn/VnPayReturn"
	VnpHashSecret = "NETBTJKFPMLLVDHFWGGHATPAHKWPLSRU"
	TimeZone      = "Asia/Ho_Chi_Minh"

	// Account test
	BankAccountNumber = "9704198526191432198"
	BankType          = "NCB"
	AccountName       = "NGUYEN VAN A"
	PasswordOTP       = "123456"
	RealeaseDate      = "7/15"
)

type (
	CreatePaymentUrl func(ctx *gin.Context, request *CreatePaymentURLRequest) (*CreatePayemntURLResponse, error)
	GetIPNUrl        func(ctx *gin.Context, request *GetIPNUrlRequest) *GetIPNUrlResponse
	GetVNPayReturn   func(ctx *gin.Context, request *GetVNPayReturnRequest) *GetIPNUrlResponse
)
