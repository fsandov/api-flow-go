package payment

type MethodType int64

const (
	WebPay     MethodType = 1
	Klap       MethodType = 3
	OnePay     MethodType = 5
	AllMethods MethodType = 9
	Mach       MethodType = 15
	Khipu      MethodType = 22
	Chek       MethodType = 25
	Fpay       MethodType = 110
)

type CreateOrderRequest struct {
	APIKey          string            `json:"apiKey"`
	CommerceOrder   string            `json:"commerceOrder"`
	Subject         string            `json:"subject"`
	Currency        string            `json:"currency"`
	Amount          int64             `json:"amount"`
	Email           string            `json:"email"`
	PaymentMethod   MethodType        `json:"paymentMethod"`
	URLConfirmation string            `json:"urlConfirmation"`
	URLReturn       string            `json:"urlReturn"`
	Optional        map[string]string `json:"optional"`
	Timeout         int64             `json:"timeout"`
	MerchantID      string            `json:"merchantId"`
	PaymentCurrency string            `json:"payment_currency"`
	S               string            `json:"s"`
}

type CreateOrderResponse struct {
	URL       string `json:"url"`
	Token     string `json:"token"`
	FlowOrder int64  `json:"flowOrder"`
}
