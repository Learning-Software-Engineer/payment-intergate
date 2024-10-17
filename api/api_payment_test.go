package api

import (
	"crypto/hmac"
	"crypto/sha512"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreate(t *testing.T) {
	ctx, _ := gin.CreateTestContext(nil)
	ctx.Request = &http.Request{
		Header: make(http.Header),
	}
	ctx.Request.Header.Set("X-Forwarded-For", "127.0.0.1")

	request := &CreatePaymentURLRequest{
		Amount:      1000,
		Locale:      "vn",
		Description: "description",
		OrderType:   "other",
		BankCode:    "NCB",
	}

	response, err := CreatePaymentUrl(ctx, request)

	assert.NoError(t, err)
	assert.NotNil(t, response)

	t.Log("Redirect URL:", response.RedirectURL)

	parsedURL, err := url.Parse(response.RedirectURL)
	assert.NoError(t, err)

	vnpParams := parsedURL.Query()

	assert.Equal(t, "2.1.0", vnpParams.Get("vnp_Version"))
	assert.Equal(t, "pay", vnpParams.Get("vnp_Command"))
	assert.Equal(t, VnpTmnCode, vnpParams.Get("vnp_TmnCode"))
	assert.Equal(t, "vn", vnpParams.Get("vnp_Locale"))
	assert.Equal(t, "VND", vnpParams.Get("vnp_CurrCode"))
	assert.Equal(t, "description", vnpParams.Get("vnp_OrderInfo"))
	assert.Equal(t, "other", vnpParams.Get("vnp_OrderType"))
	assert.Equal(t, strconv.Itoa(1000*100), vnpParams.Get("vnp_Amount"))
	assert.Equal(t, VnpReturnURL, vnpParams.Get("vnp_ReturnUrl"))

	sortedKeys := make([]string, 0, len(vnpParams))
	for key := range vnpParams {
		if key != "vnp_SecureHash" {
			sortedKeys = append(sortedKeys, key)
		}
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
	// expectedSignature := hex.EncodeToString(h.Sum(nil))

	//assert.Equal(t, expectedSignature, vnpParams.Get("vnp_SecureHash"))
}
