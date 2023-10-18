package payment

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/fsandov/api-flow-go/pkg/utils"
)

const (
	getStatusURL             = "/payment/getStatus"
	getStatusByCommerceIDURL = "/payment/getStatusByCommerceId"
)

// GetStatus gets the status of an order in Flow.cl. It returns a GetStatusResponse and an error if any.
func (c *Client) GetStatus(ctx context.Context, token string) (*Status, error) {
	if token == "" {
		return nil, fmt.Errorf("token is required")
	}

	dataString := utils.SortFields(StatusURIData{APIKey: c.APIKey, Token: token})
	signature, err := utils.GenerateHMAC(dataString, c.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("error generating hmac. err: %v", err)
	}
	encode := url.Values{}
	encode.Set("apiKey", c.APIKey)
	encode.Set("token", token)
	encodeString := encode.Encode()

	URI := fmt.Sprintf("%s%s?%s&s=%s", c.BaseURL, getStatusURL, encodeString, signature)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, URI, nil)
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
	var response *Status
	if err = json.Unmarshal(resBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}

// GetStatusByCommerceID gets the status of an order in Flow.cl by commerceID. It returns a GetStatusResponse and an error if any.
func (c *Client) GetStatusByCommerceID(ctx context.Context, commerceID string) (*Status, error) {
	if commerceID == "" {
		return nil, fmt.Errorf("token is required")
	}

	dataString := utils.SortFields(StatusURIData{APIKey: c.APIKey, CommerceID: commerceID})
	signature, err := utils.GenerateHMAC(dataString, c.SecretKey)
	if err != nil {
		return nil, fmt.Errorf("error generating hmac. err: %v", err)
	}
	encode := url.Values{}
	encode.Set("apiKey", c.APIKey)
	encode.Set("commerceId", commerceID)
	encodeString := encode.Encode()

	URI := fmt.Sprintf("%s%s?%s&s=%s", c.BaseURL, getStatusByCommerceIDURL, encodeString, signature)
	req, _ := http.NewRequestWithContext(ctx, http.MethodGet, URI, nil)
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
	var response *Status
	if err = json.Unmarshal(resBody, &response); err != nil {
		return nil, err
	}

	return response, nil
}
