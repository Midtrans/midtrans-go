package midtrans

import "fmt"

type Error struct {
	Message        string
	StatusCode     int
	RawError       error
	RawApiResponse *ApiResponse
}

// Error returns error message.
// To comply midtrans.Error with Go error interface.
func (e *Error) Error() string {
	if e.RawError != nil {
		return fmt.Sprintf("%s: %s", e.Message, e.RawError.Error())
	}
	return e.Message
}

// Unwrap method that returns its contained error
// if there is RawError supplied during error creation, return RawError. Else, will return nil
func (e *Error) Unwrap() error {
	return e.RawError
}

// GetMessage this get general message error when call api
func (e *Error) GetMessage() string {
	return e.Message
}

// GetStatusCode this get api response status code coming from midtrans backend
func (e *Error) GetStatusCode() int {
	return e.StatusCode
}

// GetRawApiResponse this get api raw response from midtrans backend
func (e *Error) GetRawApiResponse() *ApiResponse {
	return e.RawApiResponse
}

// GetRawError GetRawApiResponse this get api raw response from midtrans backend
func (e *Error) GetRawError() error {
	return e.RawError
}
