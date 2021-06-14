package midtrans

// TransactionDetails : Represent transaction details
type TransactionDetails struct {
	OrderID  string `json:"order_id"`
	GrossAmt int64  `json:"gross_amount"`
}

// ItemDetails : Represent the transaction details
type ItemDetails struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name"`
	Price        int64  `json:"price"`
	Qty          int32  `json:"quantity"`
	Brand        string `json:"brand,omitempty"`
	Category     string `json:"category,omitempty"`
	MerchantName string `json:"merchant_name,omitempty"`
}

// CustomerAddress : Represent the customer address
type CustomerAddress struct {
	FName       string `json:"first_name,omitempty"`
	LName       string `json:"last_name,omitempty"`
	Phone       string `json:"phone,omitempty"`
	Address     string `json:"address,omitempty"`
	City        string `json:"city,omitempty"`
	Postcode    string `json:"postal_code,omitempty"`
	CountryCode string `json:"country_code,omitempty"`
}

// CustomerDetails : Represent the customer detail
type CustomerDetails struct {
	// first name
	FName string `json:"first_name,omitempty"`

	// last name
	LName string `json:"last_name,omitempty"`

	Email    string           `json:"email,omitempty"`
	Phone    string           `json:"phone,omitempty"`
	BillAddr *CustomerAddress `json:"billing_address,omitempty"`
	ShipAddr *CustomerAddress `json:"customer_address,omitempty"`
}
