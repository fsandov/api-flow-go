package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/fsandov/api-flow-go/pkg/utils"
)

const (
	createOrderURL = "/payment/create"
)

// CreateOrder creates a new order in Flow.cl. It returns a CreateOrderResponse and an error if any.
func (c *Client) CreateOrder(ctx context.Context, request CreateOrderRequest) (*CreateOrderResponse, error) {
	if err := c.verifyRequiredCreateOrderRequestFields(request); err != nil {
		return nil, err
	}
	request.APIKey = c.APIKey
	dataString := utils.SortFields(request)
	signature, err := utils.GenerateHMAC(dataString, c.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("error generating hmac. err: %v", err)
	}
	request.S = signature
	form := utils.EncodeForm(request)

	URI := fmt.Sprintf("%s%s", c.BaseURL, createOrderURL)
	req, _ := http.NewRequestWithContext(ctx, http.MethodPost, URI, strings.NewReader(form))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request. err: %v", err)
	}
	defer func() {
		_ = res.Body.Close()
	}()

	if res.StatusCode == http.StatusBadRequest {
		return nil, fmt.Errorf("bad request, code: %d", res.StatusCode)
	}

	if res.StatusCode == http.StatusUnauthorized {
		return nil, fmt.Errorf("unauthorized, code: %d", res.StatusCode)
	}

	resBody, _ := io.ReadAll(res.Body)
	var response *CreateOrderResponse
	if err = json.Unmarshal(resBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) verifyRequiredCreateOrderRequestFields(req CreateOrderRequest) error {
	if c.APIKey == "" {
		return fmt.Errorf("api key is required")
	}
	if c.SecretKey == "" {
		return fmt.Errorf("secret key is required")
	}
	if req.CommerceOrder == "" {
		return fmt.Errorf("commerce order is required")
	}
	if req.Subject == "" {
		return fmt.Errorf("subject is required")
	}
	if req.Amount == 0 {
		return fmt.Errorf("amount is required")
	}
	if req.Email == "" {
		return fmt.Errorf("email is required")
	}
	if req.URLConfirmation == "" {
		return fmt.Errorf("url confirmation is required")
	}
	if req.URLReturn == "" {
		return fmt.Errorf("url return is required")
	}

	return nil
}
