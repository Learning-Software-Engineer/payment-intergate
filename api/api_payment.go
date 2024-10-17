package api

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	VnpVersion   = "2.1.0"
	VnpCommand   = "pay"
	VNLocale     = "vn"
	VnpURL       = "https://sandbox.vnpayment.vn/paymentv2/vpcpay.html"
	VnpTmnCode   = "1OSFY9NV"
	VnpReturnURL = "https://domain.vn/VnPayReturn"
	TimeZone     = "Asia/Ho_Chi_Minh"
	RestfulAPI   = "https://sandbox.vnpayment.vn/isp-svc/oauth/authenticat"
	Domain       = "http://sandbox.vnpayment.vn"

	// Information for testing
	McCode        = "20241017160637"
	CheckSum      = "b04f2b1b90f98239be38b57429731eec"
	Email         = "theflash28012002@gmail.com"
	TmnCode       = "1OSFY9NV"
	VnpHashSecret = "S67CZNUZFUSVXKV5XA84EZTJJI29OL16"

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
	vnpParams.Set("vnp_Amount", strconv.Itoa(amount))
	vnpParams.Set("vnp_Command", "pay")
	vnpParams.Set("vnp_CreateDate", createDate)
	vnpParams.Set("vnp_CurrCode", "VND")
	vnpParams.Set("vnp_IpAddr", ipAddr)
	vnpParams.Set("vnp_Locale", locale)
	vnpParams.Set("vnp_OrderInfo", request.Description)
	vnpParams.Set("vnp_OrderType", "other")
	vnpParams.Set("vnp_ReturnUrl", VnpReturnURL)
	vnpParams.Set("vnp_TmnCode", VnpTmnCode)
	vnpParams.Set("vnp_TxnRef", orderID)
	vnpParams.Set("vnp_Version", "2.1.0")

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
	urlAPI := fmt.Sprintf("https://%s/IPN", Domain)

	query := url.Values{}
	query.Set("vnp_Amount", strconv.Itoa(int(request.Amount)))
	query.Set("vnp_BankCode", BankType)
	query.Set("vnp_BankTranNo", request.BankTranNo)
	query.Set("vnp_CardType", request.CardType)
	query.Set("vnp_OrderInfo", request.OrderInfo)
	query.Set("vnp_PayDate", request.PayDate)
	query.Set("vnp_ResponseCode", request.ResponseCode)
	query.Set("vnp_TmnCode", request.TmnCode)
	query.Set("vnp_TransactionNo", request.TransactionNo)
	query.Set("vnp_TransactionStatus", request.TransactionStatus)
	query.Set("vnp_TxnRef", request.TxnRef)

	sortedKeys := make([]string, 0, len(query))
	for key := range query {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)
	signData := strings.Join(sortedKeys, "&")

	h := hmac.New(sha512.New, []byte(VnpHashSecret))
	h.Write([]byte(signData))
	signed := hex.EncodeToString(h.Sum(nil))

	if request.SecureHash == signed {
		finalUrl := fmt.Sprintf("%s?%s", urlAPI, query.Encode())
		fmt.Println(finalUrl)
		return &GetIPNUrlResponse{
			RspCode: "00",
			Message: "success",
		}
	}

	return &GetIPNUrlResponse{
		RspCode: "97",
		Message: "Fail checksum",
	}
}

func GetVNPayReturn(ctx *gin.Context, request *GetVNPayReturnRequest) *GetVNPayReturnResponse {
	return &GetVNPayReturnResponse{}
}
