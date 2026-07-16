package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

type EnigmaProxyClient struct {
	APIKey     string
	BaseURL    string
	HTTPClient *http.Client
}

// New creates a new client. An optional baseURL can be passed to override
// the default "https://enigmaproxy.net".
func New(apiKey string, baseURL ...string) *EnigmaProxyClient {

	url := "https://enigmaproxy.net"

	if len(baseURL) > 0 && baseURL[0] != "" {

		url = baseURL[0]

	}

	return &EnigmaProxyClient{

		APIKey: apiKey,

		BaseURL: strings.TrimRight(url, "/"),

		HTTPClient: &http.Client{

			Timeout: 30 * time.Second,
		},
	}
}

func (c *EnigmaProxyClient) request(

	method string,

	endpoint string,

	body interface{},

) (any, error) {

	var reader io.Reader

	if body != nil {

		data, err := json.Marshal(body)

		if err != nil {

			return nil, err

		}

		reader = bytes.NewBuffer(data)

	}

	req, err := http.NewRequest(

		method,

		c.BaseURL+endpoint,

		reader,
	)

	if err != nil {

		return nil, err

	}

	req.Header.Set(

		"Authorization",

		"Bearer "+c.APIKey,
	)

	req.Header.Set(

		"Content-Type",

		"application/json",
	)

	resp, err := c.HTTPClient.Do(req)

	if err != nil {

		return nil, err

	}

	defer resp.Body.Close()

	data, _ := io.ReadAll(resp.Body)

	if resp.StatusCode >= 400 {

		return nil,

			fmt.Errorf(

				"request failed (%d): %s",

				resp.StatusCode,

				string(data),
			)

	}

	var result any

	err = json.Unmarshal(

		data,

		&result,
	)

	return result, err

}

// ----------------------------
// Packages
// ----------------------------

func (c *EnigmaProxyClient) GetPackages() (any, error) {

	return c.request(

		"GET",

		"/api/customer/packages",

		nil,
	)

}

func (c *EnigmaProxyClient) GetPackage(

	packageID string,

) (any, error) {

	return c.request(

		"GET",

		"/api/customer/packages/"+packageID,

		nil,
	)

}

// ----------------------------
// Proxy Generator
// ----------------------------

type GenerateProxyOptions struct {
	PackageID string

	Protocol string

	Format string

	Country string

	State string

	City string

	Session bool

	SessionTime *int

	Qty int

	Lifetime *int

	FastMode bool

	HTTP3 bool
}

func (c *EnigmaProxyClient) GenerateProxy(

	opts GenerateProxyOptions,

) (any, error) {

	protocol := opts.Protocol

	if protocol == "" {

		protocol = "http"

	}

	format := opts.Format

	if format == "" {

		format = "host:port:username:password"

	}

	qty := opts.Qty

	if qty == 0 {

		qty = 1

	}

	payload := map[string]interface{}{

		"packageId": opts.PackageID,

		"protocol": protocol,

		"format": format,

		"qty": qty,

		"session": opts.Session,

		"fastMode": opts.FastMode,

		"http3": opts.HTTP3,
	}

	if opts.Country != "" {

		payload["country"] = opts.Country

	}

	if opts.State != "" {

		payload["state"] = opts.State

	}

	if opts.City != "" {

		payload["city"] = opts.City

	}

	if opts.SessionTime != nil {

		payload["sessionTime"] = *opts.SessionTime

	}

	if opts.Lifetime != nil {

		payload["lifetime"] = *opts.Lifetime

	}

	return c.request(

		"POST",

		"/api/customer/proxy",

		payload,
	)

}
