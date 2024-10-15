package api

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	VnpVersion    = "2.1.0"
	VnpCommand    = "pay"
	VNLocale      = "vn"
	VnpURL        = "https://sandbox.vnpayment.vn/paymentv2/vpcpay.html"
	VnpTmnCode    = "2QXUI4J4"
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

func CreatePaymentUrl(ctx *gin.Context, request *CreatePaymentURLRequest) (*CreatePayemntURLResponse, error) {
	ipAddr := ctx.ClientIP()

	date := time.Now()
	createDate := date.Format("20060102150405")
	orderID := date.Format("150405")

	amount := int(request.Amount * 100)

	locale := request.Locale
	if locale == "" {
		locale = "vn"
	}

	vnpParams := url.Values{}
	vnpParams.Set("vnp_Version", "2.1.0")
	vnpParams.Set("vnp_Command", "pay")
	vnpParams.Set("vnp_TmnCode", VnpTmnCode)
	vnpParams.Set("vnp_Locale", locale)
	vnpParams.Set("vnp_CurrCode", "VND")
	vnpParams.Set("vnp_TxnRef", orderID)
	vnpParams.Set("vnp_OrderInfo", request.Description)
	vnpParams.Set("vnp_OrderType", "other")
	vnpParams.Set("vnp_Amount", strconv.Itoa(amount))
	vnpParams.Set("vnp_ReturnUrl", VnpReturnURL)
	vnpParams.Set("vnp_IpAddr", ipAddr)
	vnpParams.Set("vnp_CreateDate", createDate)

	if request.BankCode != "" {
		vnpParams.Set("vnp_BankCode", request.BankCode)
	}

	sortedKeys := make([]string, 0, len(vnpParams))
	for key := range vnpParams {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	var queryStringBuilder strings.Builder
	for _, key := range sortedKeys {
		queryStringBuilder.WriteString(key)
		queryStringBuilder.WriteString("=")
		queryStringBuilder.WriteString(vnpParams.Get(key))
		queryStringBuilder.WriteString("&")
	}

	queryString := queryStringBuilder.String()
	queryString = queryString[:len(queryString)-1]

	h := hmac.New(sha512.New, []byte(VnpHashSecret))
	h.Write([]byte(queryString))
	signedData := hex.EncodeToString(h.Sum(nil))
	vnpParams.Set("vnp_SecureHash", signedData)

	paymentURL := VnpURL + "?" + vnpParams.Encode()

	return &CreatePayemntURLResponse{
		RedirectURL: paymentURL,
	}, nil
}

func GetIPNUrl(ctx *gin.Context, request *GetIPNUrlRequest) *GetIPNUrlResponse {
	return &GetIPNUrlResponse{}
}

func GetVNPayReturn(ctx *gin.Context, request *GetVNPayReturnRequest) *GetVNPayReturnResponse {
	return &GetVNPayReturnResponse{}
}
