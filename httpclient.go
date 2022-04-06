package midtrans

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type HttpClient interface {
	Call(method string, url string, apiKey *string, options *ConfigOptions, body io.Reader, result interface{}) *Error
}

// HttpClientImplementation : this is for midtrans HttpClient Implementation
type HttpClientImplementation struct {
	HttpClient *http.Client
	Logger     LoggerInterface
}

// Call the Midtrans API at specific `path` using the specified HTTP `method`. The result will be
// given to `result` if there is no error. If any error occurred, the return of this function is the `midtrans.Error`
// itself, otherwise nil.
func (c *HttpClientImplementation) Call(method string, url string, apiKey *string, options *ConfigOptions, body io.Reader, result interface{}) *Error {
	// NewRequest is used by Call to generate an http.Request.
	req, err := http.NewRequest(method, url, body)

	if err != nil {
		c.Logger.Error("Cannot create Midtrans request: %v", err)
		return &Error{
			Message:  fmt.Sprintf("Error Request creation failed: %s", err.Error()),
			RawError: err,
		}
	}

	if options != nil {
		if options.Ctx != nil {
			req.WithContext(options.Ctx)
		}

		if options.IrisIdempotencyKey != nil {
			req.Header.Add("X-Idempotency-Key", *options.IrisIdempotencyKey)
		}

		if options.PaymentIdempotencyKey != nil {
			req.Header.Add("Idempotency-Key", *options.PaymentIdempotencyKey)
		}

		if options.PaymentOverrideNotification != nil {
			req.Header.Add("X-Override-Notification", *options.PaymentOverrideNotification)
		}

		if options.PaymentAppendNotification != nil {
			req.Header.Add("X-Append-Notification", *options.PaymentAppendNotification)
		}
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "Midtrans-Go_"+libraryVersion)
	if apiKey != nil {
		key := *apiKey
		if key == "" {
			err := &Error{
				Message: "The API Key (ServerKey/IrisApiKey) is invalid, as it is an empty string. Please double-check your API key. " +
					"You can check from the Midtrans Dashboard. " +
					"See https://docs.midtrans.com/en/midtrans-account/overview?id=retrieving-api-access-keys " +
					"for the details or please contact us via https://midtrans.com/contact-us. ",
			}
			c.Logger.Error("Authentication: ", err.GetMessage())
			return err
		} else if strings.Contains(key, " ") {
			err := &Error{
				Message:  "The API Key (ServerKey/IrisApiKey) contains white-space. Please double-check your API key. " +
					"You can check the ServerKey from the Midtrans Dashboard. " +
					"See https://docs.midtrans.com/en/midtrans-account/overview?id=retrieving-api-access-keys " +
					"for the details or please contact us via https://midtrans.com/contact-us. ",
			}
			c.Logger.Error("Authentication: ", err.GetMessage())
			return err
		} else {
			req.SetBasicAuth(key, "")
		}
	}

	c.Logger.Info("================ Request ================")
	c.Logger.Info("%v Request %v %v", req.Method, req.URL, req.Proto)
	logHttpHeaders(c.Logger, req.Header, true)
	return c.DoRequest(req, result)
}

// DoRequest : is used by Call to execute an API request using HTTP client and parse the response into `result`.
func (c *HttpClientImplementation) DoRequest(req *http.Request, result interface{}) *Error {
	start := time.Now()
	res, err := c.HttpClient.Do(req)
	if err != nil {
		c.Logger.Error("Cannot send request: %v", err.Error())
		var statusCode int

		if res != nil {
			statusCode = res.StatusCode
		} else if strings.Contains(err.Error(), "timeout") {
			statusCode = 408
		} else {
			statusCode = 0
		}

		return &Error{
			Message:    fmt.Sprintf("Error when request via HttpClient, Cannot send request with error: %s", err.Error()),
			StatusCode: statusCode,
			RawError:   err,
		}
	}

	if res != nil {
		defer res.Body.Close()

		c.Logger.Info("================== END ==================")
		c.Logger.Info("Request completed in %v ", time.Since(start))

		resBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			c.Logger.Error("Request failed: %v", err)
			return &Error{
				Message:    "Cannot read response body: " + err.Error(),
				StatusCode: res.StatusCode,
				RawError: err,
			}
		}

		rawResponse := newHTTPResponse(res, resBody)
		c.Logger.Debug("=============== Response ===============")
		// Loop through headers to perform log
		logHttpHeaders(c.Logger, rawResponse.Header, false)
		c.Logger.Debug("Response Body: %v", string(rawResponse.RawBody))

		if result != nil {
			if err = json.Unmarshal(resBody, &result); err != nil {
				return &Error{
					Message:        fmt.Sprintf("Invalid body response, parse error during API request to Midtrans with message: %s", err.Error()),
					StatusCode:     res.StatusCode,
					RawError:       err,
					RawApiResponse: rawResponse,
				}
			}
		}

		// Check status_code from Midtrans response body
		if found, data := HasOwnProperty("status_code", resBody); found {
			statusCode, _ := strconv.Atoi(data["status_code"].(string))
			if statusCode >= 401 && statusCode != 407 {
				errMessage := fmt.Sprintf("Midtrans API is returning API error. HTTP status code: %s API response: %s", strconv.Itoa(statusCode), string(resBody))
				return &Error{
					Message:        errMessage,
					StatusCode:     statusCode,
					RawError: 		errors.New(errMessage),
					RawApiResponse: rawResponse,
				}
			}
		}

		// Check StatusCode from Midtrans HTTP response api StatusCode
		if res.StatusCode >= 400 {
			errMessage := fmt.Sprintf("Midtrans API is returning API error. HTTP status code: %s  API response: %s", strconv.Itoa(res.StatusCode), string(resBody))
			return &Error{
				Message:        errMessage,
				StatusCode:     res.StatusCode,
				RawError:       errors.New(errMessage),
				RawApiResponse: rawResponse,
			}
		}
	}
	return nil
}

// ApiResponse : is a structs that may come from Midtrans API endpoints
type ApiResponse struct {
	Status     string // e.g. "200 OK"
	StatusCode int    // e.g. 200
	Proto      string // e.g. "HTTP/1.0"

	// response Header contain a map of all HTTP header keys to values.
	Header http.Header
	// response body
	RawBody []byte
	// request that was sent to obtain the response
	Request *http.Request
}

// newHTTPResponse : internal function to set HTTP Raw response return to ApiResponse
func newHTTPResponse(res *http.Response, responseBody []byte) *ApiResponse {
	return &ApiResponse{
		Status:     res.Status,
		StatusCode: res.StatusCode,
		Proto:      res.Proto,
		Header:     res.Header,
		RawBody:    responseBody,
		Request:    res.Request,
	}
}

// logHttpHeaders : internal function to perform log from headers
func logHttpHeaders(log LoggerInterface, header http.Header, isReq bool) {
	// Loop through headers to perform log
	for name, headers := range header {
		name = strings.ToLower(name)
		for _, h := range headers {
			if name == "authorization" {
				log.Debug("%v: %v", name, h)
			} else {
				if isReq {
					log.Info("%v: %v", name, h)
				} else {
					log.Debug("%v: %v", name, h)
				}
			}
		}
	}
}

//HasOwnProperty : Convert HTTP raw response body to map and check if the body has own field
func HasOwnProperty(key string, body []byte) (bool, map[string]interface{}) {
	d := make(map[string]interface{})
	_ = json.Unmarshal(body, &d)
	if _, found := d[key].(string); found {
		return found, d
	} else {
		return found, d
	}
}
