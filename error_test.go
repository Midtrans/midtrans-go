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
	assert.Error(t, err)
	assert.Equal(t, `Error Test Message`, err.GetMessage())
	assert.Equal(t, 200, err.GetStatusCode())
	assert.Equal(t, `TEST FROM GO ERROR`, err.Error())
}

func TestErrorResponse(t *testing.T) {
	serverKey := "dummy"
	c := GetHttpClient(Environment)
	jsonReq, _ := json.Marshal("{\"transaction_details\": {\"order_id\": \"TEST-1648108994111\", \"gross_amount\": 10000}}")
	err := c.Call(http.MethodPost, "https://app.midtrans.com/snap/v1/transactions", &serverKey, nil, bytes.NewBuffer(jsonReq), nil)

	assert.Error(t, err)
	assert.Equal(t, 401, err.StatusCode)
	assert.Equal(t, "app.midtrans.com", err.RawApiResponse.Request.Host)
}