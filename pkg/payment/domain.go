package payment

type MethodType int64

const (
	// WebPay method type for WebPay only
	WebPay MethodType = 1
	// Klap method type for Klap only
	Klap MethodType = 3
	// OnePay method type for OnePay only
	OnePay MethodType = 5
	// AllMethods method type for all methods available
	AllMethods MethodType = 9
	// Mach method type for Mach only
	Mach MethodType = 15
	// Khipu method type for Khipu only
	Khipu MethodType = 22
	// Chek method type for Chek only
	Chek MethodType = 25
	// Fpay method type for Fpay only
	Fpay MethodType = 110
)

// CreateOrderRequest is the request body for CreateOrder method.
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

// CreateOrderResponse is the response body for CreateOrder method.
type CreateOrderResponse struct {
	URL       string `json:"url"`
	Token     string `json:"token"`
	FlowOrder int64  `json:"flowOrder"`
}

// Status is the response body for GetStatus method.
type Status struct {
	FlowOrder     int64             `json:"flowOrder"`
	CommerceOrder string            `json:"commerceOrder"`
	RequestDate   string            `json:"requestDate"`
	Status        int               `json:"status"`
	Subject       string            `json:"subject"`
	Currency      string            `json:"currency"`
	Amount        string            `json:"amount"`
	Payer         string            `json:"payer"`
	Optional      map[string]string `json:"optional,omitempty"`
	PendingInfo   PendingInfo       `json:"pending_info,omitempty"`
	PaymentData   StatusData        `json:"paymentData,omitempty"`
	MerchantId    string            `json:"merchantId"`
}

// StatusData is the part of response body for GetStatus method.
type StatusData struct {
	Date           string  `json:"date"`
	Media          string  `json:"media"`
	ConversionDate string  `json:"conversionDate"`
	ConversionRate float64 `json:"conversionRate"`
	Amount         string  `json:"amount"`
	Currency       string  `json:"currency"`
	Fee            string  `json:"fee"`
	Balance        int     `json:"balance"`
	TransferDate   string  `json:"transferDate"`
}

// PendingInfo is the part of response body for GetStatus method.
type PendingInfo struct {
	Media string `json:"media"`
	Date  string `json:"date"`
}

// StatusURIData is the data to be encoded for GetStatus method.
type StatusURIData struct {
	APIKey string `json:"apiKey"`
	Token  string `json:"token"`
	S      string `json:"s"`
}
