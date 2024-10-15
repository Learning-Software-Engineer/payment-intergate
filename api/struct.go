package api

type CreatePaymentURLRequest struct {
	OrderType   string  `json:"order_type,omitempty"`
	Amount      float64 `json:"amount,omitempty"`
	BankCode    string  `json:"bank_code,omitempty"`
	Description string  `json:"description,omitempty"` // order info
	Locale      string  `json:"locale,omitempty"`
}

type CreatePayemntURLResponse struct {
	RedirectURL string `json:"redirect_url"`
}

type GetIPNUrlRequest struct {
	SecureHash     string `json:"vpn_SecureHash"`
	SecureHashType string `json:"vpn_SecureHashType,omitempty"`
	TxnRef         string `json:"vnp_TxnRef"`
	ResponseCode   string `json:"vnp_ResponseCode"`
}

type GetIPNUrlResponse struct {
	RspCode string `json:"rsp_code"`
	Message string `json:"message"`
}

type GetVNPayReturnRequest struct {
	SecureHash     string `json:"vpn_SecureHash"`
	SecureHashType string `json:"vpn_SecureHashType,omitempty"`
	ResponseCode   string `json:"vnp_ResponseCode"`
}

type GetVNPayReturnResponse struct {
	RspCode string `json:"rsp_code"`
	Message string `json:"message"`
}
