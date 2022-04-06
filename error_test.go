package midtrans

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestErrorStruct(t *testing.T) {
	err := &Error{
		Message:        "Error Test Message",
		StatusCode:     200,
		RawError:       errors.New("TEST FROM GO ERROR"),
		RawApiResponse: nil,
	}
	var midError *Error
	assert.Error(t, err)
	assert.True(t, true, errors.Is(err, err))
	assert.True(t, true, errors.As(err, &midError))
	assert.Equal(t, `Error Test Message`, err.GetMessage())
	assert.Equal(t, 200, err.GetStatusCode())
	assert.Equal(t, "Error Test Message: TEST FROM GO ERROR", err.Error())
}

func TestErrorResponse(t *testing.T) {
	serverKey := "dummy"
	c := GetHttpClient(Environment)
	jsonReq, _ := json.Marshal("{\"transaction_details\": {\"order_id\": \"TEST-1648108994111\", \"gross_amount\": 10000}}")
	err := c.Call(http.MethodPost, "https://app.midtrans.com/snap/v1/transactions", &serverKey, nil, bytes.NewBuffer(jsonReq), nil)

	var midError *Error
	assert.True(t, true, errors.Is(err, err))
	assert.True(t, true, errors.As(err, &midError))
	assert.Error(t, err)
	assert.Equal(t, 401, err.StatusCode)
	assert.Equal(t, "app.midtrans.com", err.RawApiResponse.Request.Host)
}