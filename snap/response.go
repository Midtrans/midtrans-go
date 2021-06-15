package snap

// Response : Snap response after calling the Snap API
type Response struct {
	Token         string   `json:"token"`
	RedirectURL   string   `json:"redirect_url"`
	StatusCode    string   `json:"status_code,omitempty"`
	ErrorMessages []string `json:"error_messages,omitempty"`
}

// ResponseWithMap : Snap response with map after calling the Snap API
type ResponseWithMap map[string]interface{}