package payment

import (
	"fmt"
	"net/http"
	"time"
)

type BaseURL string

const (
	production     BaseURL = "https://www.flow.cl/api"
	sandbox        BaseURL = "https://sandbox.flow.cl/api"
	defaultTimeout         = 2000
)

var (
	errAPIKeyRequired    = fmt.Errorf("api key is required")
	errSecretKeyRequired = fmt.Errorf("secret key is required")
)

type (
	clientOptions struct {
		Timeout int64
		BaseURL BaseURL
	}

	ClientOption func(*clientOptions)
)

// WithTimeOut sets the timeout for the http client in milliseconds. Default value is 2000. This is optional.
func WithTimeOut(timeout int64) ClientOption {
	return func(o *clientOptions) {
		if timeout > 0 {
			o.Timeout = timeout
		} else {
			o.Timeout = defaultTimeout
		}
	}
}

// WithProductionURL sets the BaseURL to Production. This is optional.
func WithProductionURL() ClientOption {
	return func(o *clientOptions) {
		o.BaseURL = production
	}
}

func (o *clientOptions) apply(opts ...ClientOption) {
	for _, opt := range opts {
		opt(o)
	}
}

// Client is the client for the Flow.cl API. It contains the http client, the BaseURL, the APIKey and the SecretKey.
type Client struct {
	HTTPClient *http.Client
	BaseURL    BaseURL
	APIKey     string
	SecretKey  string
}

// NewClient returns a new Client with the given APIKey and SecretKey. The BaseURL is set to Sandbox by default.
// This can be changed by passing WithProductionURL as an option. The timeout for the http client is set to 2000. This
// can be changed by passing WithTimeOut as an option. The timeout is set in milliseconds.
func NewClient(APIKey, SecretKey string, opts ...ClientOption) (*Client, error) {
	if APIKey == "" {
		return nil, errAPIKeyRequired
	}
	if SecretKey == "" {
		return nil, errSecretKeyRequired
	}

	options := &clientOptions{}
	options.apply(opts...)

	if options.BaseURL == "" {
		options.BaseURL = sandbox
	}

	return &Client{
		HTTPClient: &http.Client{
			Timeout: time.Duration(options.Timeout) * time.Millisecond,
		},
		BaseURL:   options.BaseURL,
		APIKey:    APIKey,
		SecretKey: SecretKey,
	}, nil
}
