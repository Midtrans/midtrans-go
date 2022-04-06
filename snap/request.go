package snap

import (
	"github.com/midtrans/midtrans-go"
)

// RequestParamWithMap SnapReqWithMap : Represent snap request with map payload
type RequestParamWithMap map[string]interface{}

// Request : Represent SNAP API request payload that are used in Create Snap Token parameter.
// https://snap-docs.midtrans.com/#json-objects
type Request struct {
	TransactionDetails midtrans.TransactionDetails `json:"transaction_details"`
	Items              *[]midtrans.ItemDetails     `json:"item_details,omitempty"`
	CustomerDetail     *midtrans.CustomerDetails   `json:"customer_details,omitempty"`
	EnabledPayments    []SnapPaymentType           `json:"enabled_payments,omitempty"`
	CreditCard         *CreditCardDetails          `json:"credit_card,omitempty"`
	BcaVa              *BcaVa                      `json:"bca_va,omitempty"`
	BniVa              *BniVa                      `json:"bni_va,omitempty"`
	BriVa              *BriVa                      `json:"bri_va,omitempty"`
	PermataVa          *PermataVa                  `json:"permata_va,omitempty"`
	Gopay              *GopayDetails               `json:"gopay,omitempty"`
	ShopeePay          *ShopeePayDetails           `json:"shopeepay,omitempty"`
	Callbacks          *Callbacks                  `json:"callbacks,omitempty"`
	Expiry             *ExpiryDetails              `json:"expiry,omitempty"`
	UserId             string                      `json:"user_id,omitempty"`
	Cstore             *Cstore                     `json:"cstore,omitempty"`
	CustomField1       string                      `json:"custom_field1,omitempty"`
	CustomField2       string                      `json:"custom_field2,omitempty"`
	CustomField3       string                      `json:"custom_field3,omitempty"`
	Metadata           interface{}                 `json:"metadata,omitempty"`
}

// CreditCardDetails : Represent credit card detail for PaymentTypeCreditCard payment type
type CreditCardDetails struct {
	// indicate if generated token should be saved for next charge
	SaveCard bool `json:"save_card,omitempty"`

	// Use 3D-Secure authentication when using credit card. Default: false
	Secure bool `json:"secure,omitempty"`

	// Acquiring bank. Valid values: `midtrans.BankBca` `midtrans.BankMandiri`, `midtrans.BankBni`,
	//`midtrans.BankCimb`, `midtrans.BankMaybank`, and `midtrans.BankBri`
	Bank string `json:"bank,omitempty"`

	// Acquiring channel. Options: migs
	Channel string `json:"channel,omitempty"`

	// Credit card transaction type. Options: authorize, authorize_capture. Default: “authorize_capture”
	Type string `json:"type,omitempty"`

	// Snap for installment detail
	Installment *InstallmentDetail `json:"installment,omitempty"`

	// Allowed credit card BIN numbers.
	// The bin value can be either a prefix(upto 8 digits) of card number or the name of a bank,
	// in which case all the cards issued by that bank will be allowed.
	// The supported bank names are bni bca mandiri cimb bri and maybank. Default: allow all cards
	WhitelistBins []string `json:"whitelist_bins,omitempty"`

	DynamicDescriptor *DynamicDescriptor `json:"dynamic_descriptor,omitempty"`

	CardToken string `json:"card_token,omitempty"`
}

// InstallmentDetail : Represent installment detail
type InstallmentDetail struct {
	// Force installment when using credit card. Default: false
	Required bool `json:"required,omitempty"`

	// Available installment terms
	Terms *InstallmentTermsDetail `json:"terms,omitempty"`
}

// InstallmentTermsDetail : Represent installment available banks
type InstallmentTermsDetail struct {
	Bni     []int8 `json:"bni,omitempty"`
	Mandiri []int8 `json:"mandiri,omitempty"`
	Cimb    []int8 `json:"cimb,omitempty"`
	Mega    []int8 `json:"mega,omitempty"`
	Bca     []int8 `json:"bca,omitempty"`
	Bri     []int8 `json:"bri,omitempty"`
	Maybank []int8 `json:"maybank,omitempty"`
	Offline []int8 `json:"offline,omitempty"`
}

type DynamicDescriptor struct {
	// First 25 digit on customer’s billing statement. Mostly used to show the merchant or product name.
	// Only works for BNI.
	MerchantName string `json:"merchant_name,omitempty"`

	// Next 13 digit on customer’s billing statement. It works as secondary metadata on the statement.
	// Mostly used to show city name or region. Only works for BNI.
	CityName string `json:"city_name,omitempty"`

	// Last 2 digit on customer’s billing statement. Mostly used to show country code.
	// The format is ISO 3166-1 alpha-2. Only works for BNI.
	CountryCode string `json:"country_code,omitempty"`
}

// BcaVa : BCA Virtual Account is a virtual payment method offered by Bank BCA.
// https://snap-docs.midtrans.com/#bca-virtual-account
type BcaVa struct {
	// VaNumber : Custom VA Number, Length should be within 1 to 11.
	// https://snap-docs.midtrans.com/#custom-virtual-account-number
	VaNumber       string `json:"va_number,omitempty"`
	SubCompanyCode string `json:"sub_company_code,omitempty"`
	FreeText       struct {
		Inquiry []struct {
			En string `json:"en"`
			Id string `json:"id"`
		} `json:"inquiry,omitempty"`
		Payment []struct {
			En string `json:"en"`
			Id string `json:"id"`
		} `json:"payment,omitempty"`
	} `json:"free_text,omitempty"`
}

// BniVa : BNI Virtual Account is a virtual payment method offered by Bank BNI.
// https://snap-docs.midtrans.com/#bni-virtual-account
type BniVa struct {
	// VaNumber : Custom VA Number, Length should be within 1 to 8.
	// https://snap-docs.midtrans.com/#custom-virtual-account-number
	VaNumber string `json:"va_number,omitempty"`
}

// BriVa : BRI Virtual Account is a virtual payment method offered by Bank BRI.
// https://snap-docs.midtrans.com/#bri-virtual-account
type BriVa struct {
	// VaNumber : Custom VA Number, Length should be within 1 to 13.
	// https://snap-docs.midtrans.com/#custom-virtual-account-number
	VaNumber string `json:"va_number,omitempty"`
}

// PermataVa : Permata Virtual Account is a virtual payment method offered by Bank Permata.
// https://snap-docs.midtrans.com/#permata-virtual-account
type PermataVa struct {
	// VaNumber : Custom VA Number, Length should be 10. Only supported for b2b transactions.
	// https://snap-docs.midtrans.com/#custom-virtual-account-number
	VaNumber      string `json:"va_number,omitempty"`
	RecipientName string `json:"recipient_name,omitempty"`
}

// GopayDetails : Represent gopay detail
type GopayDetails struct {
	// EnableCallback : Enable redirect back to merchant from GoJek apps. Default: false
	EnableCallback bool `json:"enable_callback,omitempty"`

	// CallbackUrl : Determine where should customer be redirected from GoJek apps.
	// It supports both HTTP and deeplink. Default: same value as finish url
	CallbackUrl string `json:"callback_url,omitempty"`
}

// ShopeePayDetails : Represent shopeepay detail
type ShopeePayDetails struct {
	// CallbackUrl : Determine where should customer be redirected from Shopee apps.
	// It supports both HTTP and deeplink. Default: same value as finish url, if it’s also empty,
	// it will be redirected to default payment processed page
	CallbackUrl string `json:"callback_url,omitempty"`
}

// Cstore : Cstore object is for PaymentTypeAlfamart free text
type Cstore struct {
	AlfamartFreeText1 string `json:"alfamart_free_text_1,omitempty"`
	AlfamartFreeText2 string `json:"alfamart_free_text_2,omitempty"`
	AlfamartFreeText3 string `json:"alfamart_free_text_3,omitempty"`
}

// ExpiryDetails : Represent SNAP expiry details
type ExpiryDetails struct {
	// StartTime : Timestamp in yyyy-MM-dd HH:mm:ss Z format. If not specified,
	// transaction time will be used as start time (when customer charge)
	StartTime string `json:"start_time,omitempty"`

	// Unit Expiry unit. Options: day, hour, minute (plural term also accepted)
	Unit string `json:"unit"`

	// Duration Expiry duration
	Duration int64 `json:"duration"`
}

// Callbacks : Redirect URL after transaction is successfully paid (Overridden by JS callback).
// Can also be set via Snap Settings menu in your dashboard.
type Callbacks struct {
	Finish string `json:"finish,omitempty"`
}
