package couchbasecloud

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const couchbaseCloudApiVersion = "v2"
const couchbaseCloudBaseUrl = "https://cloudapi.cloud.couchbase.com"
const requestTimeoutMinutes = 5

type CouchbaseCloudClient struct {
	apiAccessKey string
	apiSecretKey string
	BaseURL      string
	HTTPClient   *http.Client
}

type Cursor struct {
	Hrefs struct {
		First string `json:"first"`
		Last  string `json:"last"`
	} `json:"hrefs"`
	Pages struct {
		Last       int `json:"last"`
		Page       int `json:"page"`
		PerPage    int `json:"perPage"`
		TotalItems int `json:"totalItems"`
	} `json:"pages"`
}

type errorResponse struct {
	Code      int    `json:"code"`
	ErrorType string `json:"errorType"`
	Message   string `json:"message"`
}

func NewClient(apiAccessKey string, apiSecretKey string) *CouchbaseCloudClient {
	return &CouchbaseCloudClient{
		apiAccessKey: apiAccessKey,
		apiSecretKey: apiSecretKey,
		HTTPClient: &http.Client{
			Timeout: requestTimeoutMinutes * time.Minute,
		},
		BaseURL: couchbaseCloudBaseUrl,
	}
}

func (client *CouchbaseCloudClient) do(r *request, response interface{}) error {
	timestamp := getEpochTimestampMs()

	var ctx context.Context
	if r.timeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), r.timeout)
		defer cancel()
	} else {
		ctx = context.Background()
	}

	req, err := http.NewRequestWithContext(ctx, r.method, client.BaseURL+client.getApiEndpoint(r.endpoint),
		bytes.NewReader(r.body))
	if err != nil {
		return err
	}

	if len(r.queryParameters) != 0 {
		req.URL.RawQuery = r.queryParameters.Encode()
	}

	bearerToken := getBearerToken(client.apiAccessKey, client.apiSecretKey, req.Method, req.URL.RequestURI(), timestamp)

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken))
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Couchbase-Timestamp", timestamp)

	if r.contentType != "" {
		req.Header.Set("Content-Type", contentTypeJSON)
	}

	res, err := client.HTTPClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		var errRes errorResponse
		if err = json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	return json.NewDecoder(res.Body).Decode(&response)
}

func (client *CouchbaseCloudClient) getApiEndpoint(endpoint string) string {
	return fmt.Sprintf("/%s%s", couchbaseCloudApiVersion, endpoint)
}

func getBearerToken(accessKey, secretKey, requestMethod string, requestUri string, timestamp string) string {
	signature := getAPISignature(secretKey, requestMethod, requestUri, timestamp)
	return accessKey + ":" + signature
}

func getAPISignature(secretKey string, requestMethod string, requestUri string, timestamp string) string {
	message := strings.Join([]string{requestMethod, requestUri, timestamp}, "\n")
	return computeHmac256(secretKey, message)
}

func getEpochTimestampMs() string {
	return strconv.FormatInt(time.Now().UnixNano()/int64(time.Millisecond), 10)
}

func computeHmac256(secret string, message string) string {
	key := []byte(secret)
	hash := hmac.New(sha256.New, key)
	hash.Write([]byte(message))
	return base64.StdEncoding.EncodeToString(hash.Sum(nil))
}
