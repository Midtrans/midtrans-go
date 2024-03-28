package midtrans

import (
	"context"
	"net/http"
	"strings"
	"time"
)

// EnvironmentType is global config Environment for Midtrans api
type EnvironmentType int8

const (
	_ EnvironmentType = iota

	//Sandbox : represent sandbox environment
	Sandbox

	//Production : represent production environment
	Production

	//libraryVersion : midtrans go library version
	libraryVersion = "v1.3.8"
)

// ServerKey is config payment API key for global use
var ServerKey string

// ClientKey is config payment public API key for global use
var ClientKey string

// PaymentOverrideNotification opt to change or add custom notification urls globally on every transaction.
var PaymentOverrideNotification *string

// PaymentAppendNotification opt to change or set custom notification urls globally on every transaction.
var PaymentAppendNotification *string

// SetPaymentOverrideNotification opt to change or set custom notification urls globally on every transaction.
// To use new notification url(s) disregarding the settings on Midtrans dashboard, only receive up to maximum of 3 urls.
func SetPaymentOverrideNotification(val string) {
	PaymentOverrideNotification = &val
}

// SetPaymentAppendNotification opt to change or add custom notification urls globally on every transaction.
// To use new notification url(s) disregarding the settings on Midtrans dashboard, only receive up to maximum of 3 urls.
func SetPaymentAppendNotification(val string) {
	PaymentAppendNotification = &val
}

var (
	//Environment default Environment for Midtrans API
	Environment = Sandbox

	//DefaultHttpTimeout default timeout for go HTTP HttpClient
	DefaultHttpTimeout = 80 * time.Second

	//DefaultGoHttpClient default Go HTTP Client for Midtrans HttpClient API
	DefaultGoHttpClient = &http.Client{Timeout: DefaultHttpTimeout}

	//DefaultLoggerLevel logging level that will be used for config globally by Midtrans logger
	DefaultLoggerLevel = &LoggerImplementation{LogLevel: LogError}

	//defaultHttpClientImplementation
	defaultHttpClientImplementation = &HttpClientImplementation{
		HttpClient: DefaultGoHttpClient,
		Logger:     GetDefaultLogger(Environment),
	}
)

// GetDefaultLogger the default logger that the library will use to log errors, debug, and informational messages.
func GetDefaultLogger(env EnvironmentType) LoggerInterface {
	if env == Sandbox {
		return &LoggerImplementation{LogLevel: LogDebug}
	} else {
		return DefaultLoggerLevel
	}
}

// GetHttpClient : get HttpClient implementation
func GetHttpClient(Env EnvironmentType) *HttpClientImplementation {
	return &HttpClientImplementation{
		HttpClient: DefaultGoHttpClient,
		Logger:     GetDefaultLogger(Env),
	}
}

var typeString = map[EnvironmentType]string{
	Sandbox:    "https://api.sandbox.midtrans.com",
	Production: "https://api.midtrans.com",
}

// BaseUrl To get Midtrans Base URL
func (e EnvironmentType) BaseUrl() string {
	for k, v := range typeString {
		if k == e {
			return v
		}
	}
	return "undefined"
}

// SnapURL : To get Snap environment API URL
func (e EnvironmentType) SnapURL() string {
	return strings.Replace(e.BaseUrl(), "api.", "app.", 1)
}

// IrisURL : To get Iris environment API URL
func (e EnvironmentType) IrisURL() string {
	return strings.Replace(e.BaseUrl(), "api.", "app.", 1) + "/iris"
}

// ConfigOptions : is used to configure some feature before request to Midtrans API
// via `coreapi.Gateway` `snap.Gateway` and `iris.Gateway`
type ConfigOptions struct {
	PaymentIdempotencyKey       *string
	PaymentOverrideNotification *string
	PaymentAppendNotification   *string
	IrisIdempotencyKey          *string
	Ctx                         context.Context
}

// SetPaymentIdempotencyKey : options to change or add unique idempotency-key on header on Midtrans Payment API request with key maximum length is 36.
// To safely handle retry request without performing the same operation twice. This is helpful for cases where merchant didn't
// receive the response because of network issue or other unexpected error.
func (o *ConfigOptions) SetPaymentIdempotencyKey(val string) {
	o.PaymentIdempotencyKey = &val
}

// SetIrisIdempotencyKey : options to change or add unique idempotency-key on header on Iris API request with key maximum length is 100.
// To safely handle retry request without performing the same operation twice. This is helpful for cases where merchant didn't
// receive the response because of network issue or other unexpected error.
func (o *ConfigOptions) SetIrisIdempotencyKey(val string) {
	o.IrisIdempotencyKey = &val
}

// SetPaymentOverrideNotification : options to change or add custom notification urls on every transaction.
// To use new notification url(s) disregarding the settings on Midtrans dashboard, only receive up to maximum of 3 urls
func (o *ConfigOptions) SetPaymentOverrideNotification(val string) {
	o.PaymentOverrideNotification = &val
}

// SetPaymentAppendNotification : options to change or add custom notification urls on every transaction.
// To use new notification url(s) disregarding the settings on Midtrans dashboard, only receive up to maximum of 3 urls
func (o *ConfigOptions) SetPaymentAppendNotification(val string) {
	o.PaymentAppendNotification = &val
}

// SetContext : options to change or add Context for each API request
func (o *ConfigOptions) SetContext(ctx context.Context) {
	o.Ctx = ctx
}
