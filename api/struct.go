package api

type CreateTokenRequest struct {
	ClientID     string `json:"client_id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	ClientSecret string `json:"client_secret"`
}

type CreatePaymentURLRequest struct {
	OrderType   string  `json:"vnp_OrderType,omitempty"`
	Amount      float64 `json:"vnp_Amount,omitempty"`
	BankCode    string  `json:"vnp_BankCode,omitempty"`
	Description string  `json:"vnp_OrderInfo,omitempty"` // order info
	Locale      string  `json:"vnp_Locale,omitempty"`
}

type CreatePayemntURLResponse struct {
	RedirectURL string `json:"redirect_url"`
}

type GetIPNUrlRequest struct {
	Amount            int    `json:"vnp_Amount,omitempty"`
	BankCode          string `json:"vnp_BankCode,omitempty"`
	BankTranNo        string `json:"vnp_BankTranNo"`
	CardType          string `json:"vnp_CardType"` // Value: ATM or QRCODE
	OrderInfo         string `json:"vnp_OrderInfo,omitempty"`
	PayDate           string `json:"vnp_PayDate,omitempty"`
	ResponseCode      string `json:"vnp_ResponseCode"`
	TmnCode           string `json:"vnp_TmnCode,omitempty"`
	TransactionNo     string `json:"vnp_TransactionNo"`
	TransactionStatus string `json:"vnp_TransactionStatus,omitempty"`
	TxnRef            string `json:"vnp_TxnRef"`
	SecureHash        string `json:"vpn_SecureHash"`
}

type GetIPNUrlResponse struct {
	RspCode string `json:"rsp_code"`
	Message string `json:"message"`
}

type GetVNPayReturnRequest struct {
	Amount            int    `json:"vnp_Amount,omitempty"`
	BankCode          string `json:"vnp_BankCode,omitempty"`
	BankTranNo        string `json:"vnp_BankTranNo"`
	CardType          string `json:"vnp_CardType"`
	OrderInfo         string `json:"vnp_OrderInfo,omitempty"`
	PayDate           string `json:"vnp_PayDate,omitempty"`
	ResponseCode      string `json:"vnp_ResponseCode"`
	TransactionNo     string `json:"vnp_TransactionNo"`
	TransactionStatus string `json:"vnp_TransactionStatus,omitempty"`
	TxnRef            string `json:"vnp_TxnRef"`
	SecureHash        string `json:"vpn_SecureHash"`
}

type GetVNPayReturnResponse struct {
	RspCode string `json:"rsp_code"`
	Message string `json:"message"`
}
