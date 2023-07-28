package clients

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const (
	urlLogin    = "/api/auth/login"
	urlLogout   = "/api/auth/logout"
	urlRegister = "/api/auth/register"
)

type HTTPClient interface {
	Login(request LoginRequest) (*LoginResponseWrapper, error)
	Logout(token string) (*Response, error)
	Register(request RegisterRequest) (*RegisterResponseWrapper, error)
}

type HTTPClientOption func(c *httpClient)

func WithClient(client *http.Client) HTTPClientOption {
	return func(c *httpClient) {
		c.client = client
	}
}

func applyHTTPClientOptions(c *httpClient, options ...HTTPClientOption) {
	for _, opt := range options {
		opt(c)
	}
}

type httpClient struct {
	host   string
	client *http.Client
}

func (c *httpClient) json(method string, addr string, req, rsp any) error {
	buf := bytes.NewBuffer([]byte{})
	if err := json.NewEncoder(buf).Encode(req); err != nil {
		return err
	}
	request, err := http.NewRequest(method, addr, buf)
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	response, err := c.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		//return
	}
	return json.NewDecoder(response.Body).Decode(rsp)
}

func NewHTTPClient(host string, options ...HTTPClientOption) HTTPClient {
	host = strings.TrimRight(host, "/")
	if !strings.HasPrefix(host, "http") {
		host = "https://" + host
	}
	client := &httpClient{host: host, client: &http.Client{Timeout: time.Second * 5}}
	applyHTTPClientOptions(client, options...)
	return client
}

func (c *httpClient) Login(request LoginRequest) (*LoginResponseWrapper, error) {
	addr := fmt.Sprintf("%s%s", c.host, urlLogin)
	rsp := new(LoginResponseWrapper)
	err := c.json(http.MethodPost, addr, request, rsp)
	return rsp, err
}

func (c *httpClient) Logout(token string) (*Response, error) {
	addr := fmt.Sprintf("%s%s", c.host, urlLogout)
	request, err := http.NewRequest(http.MethodPost, addr, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Set("Authorization", "Bearer "+token)
	response, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		//return
	}
	rsp := new(Response)
	err = json.NewDecoder(response.Body).Decode(rsp)
	return rsp, err
}

func (c *httpClient) Register(request RegisterRequest) (*RegisterResponseWrapper, error) {
	addr := fmt.Sprintf("%s%s", c.host, urlRegister)
	rsp := new(RegisterResponseWrapper)
	err := c.json(http.MethodPost, addr, request, rsp)
	return rsp, err
}
