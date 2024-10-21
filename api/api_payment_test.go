package api

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func TestCreate(t *testing.T) {
	ctx, _ := gin.CreateTestContext(nil)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Request.Header.Set("X-Forwarded-For", "127.0.0.1")

	request := &CreatePaymentURLRequest{
		Amount:      1000,
		Description: "description",
	}

	date := time.Now()
	createDate := date.Format("20060102150405")
	orderID := date.Format("150405")

	amount := int(request.Amount * 100)

	vnpParams := url.Values{}
	vnpParams.Set("vnp_Version", "2.1.0")
	vnpParams.Set("vnp_Command", "pay")
	vnpParams.Set("vnp_TmnCode", "MA3RBGJO")
	vnpParams.Set("vnp_Locale", "vn")
	vnpParams.Set("vnp_TxnRef", orderID)
	vnpParams.Set("vnp_OrderInfo", request.Description)
	vnpParams.Set("vnp_OrderType", "other")
	vnpParams.Set("vnp_Amount", strconv.Itoa(amount))
	vnpParams.Set("vnp_ReturnUrl", "https://domainmerchant.vn/ReturnUrl")
	vnpParams.Set("vnp_IpAddr", "127.0.0.1")
	vnpParams.Set("vnp_CreateDate", createDate)
	vnpParams.Set("vnp_CurrCode", "VND")
	if request.BankCode != "" {
		vnpParams.Set("vnp_BankCode", request.BankCode)
	}

	signed := generateHMACSHA512(createSignData(vnpParams), VnpHashSecret)
	vnpParams.Set("vnp_SecureHash", signed)

	fmt.Println("signed::::", signed)

	finalUrl := fmt.Sprintf("%s?%s", VnpURL, vnpParams.Encode())

	fmt.Println("Payment URL:", finalUrl)
}

func createSignData(params url.Values) string {
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	var signData []string
	for _, key := range keys {
		signData = append(signData, fmt.Sprintf("%s=%s", key, params.Get(key)))
	}
	res := strings.Join(signData, "&")
	fmt.Println("res", res)
	return res
}

func encodeParams(params map[string]string) string {
	var keys []string
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var encodedParams []string
	for _, key := range keys {
		encodedParams = append(encodedParams, fmt.Sprintf("%s=%s", url.QueryEscape(key), url.QueryEscape(params[key])))
	}
	return strings.Join(encodedParams, "&")
}

func generateHMACSHA512(data, secret string) string {
	h := hmac.New(sha512.New, []byte(secret))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
