package coreapi

import (
	"github.com/midtrans/midtrans-go"
)

// ChargeReqWithMap : Represent Charge request with map payload
type ChargeReqWithMap map[string]interface{}

// ChargeReq : Represent Charge request payload
type ChargeReq struct {
	PaymentType        CoreapiPaymentType          `json:"payment_type"`
	TransactionDetails midtrans.TransactionDetails `json:"transaction_details"`

	Items           *[]midtrans.ItemDetails   `json:"item_details,omitempty"`
	CustomerDetails *midtrans.CustomerDetails `json:"customer_details,omitempty"`

	CreditCard      *CreditCardDetails      `json:"credit_card,omitempty"`
	BankTransfer    *BankTransferDetails    `json:"bank_transfer,omitempty"`
	EChannel        *EChannelDetail         `json:"echannel,omitempty"`
	Gopay           *GopayDetails           `json:"gopay,omitempty"`
	ShopeePay       *ShopeePayDetails       `json:"shopeepay,omitempty"`
	Qris            *QrisDetails            `json:"qris,omitempty"`
	BCAKlikPay      *BCAKlikPayDetails      `json:"bca_klikpay,omitempty"`
	BCAKlikBCA      *BcaKlikBCADetails      `json:"bca_klikbca,omitempty"`
	MandiriClickPay *MandiriClickPayDetails `json:"mandiri_clickpay,omitempty"`
	CIMBClicks      *CIMBClicksDetails      `json:"cimb_clicks,omitempty"`

	ConvStore *ConvStoreDetails `json:"cstore,omitempty"`

	CustomExpiry *CustomExpiry `json:"custom_expiry,omitempty"`
	CustomField1 *string       `json:"custom_field1,omitempty"`
	CustomField2 *string       `json:"custom_field2,omitempty"`
	CustomField3 *string       `json:"custom_field3,omitempty"`
	Metadata     interface{}   `json:"metadata,omitempty"`
}

// CreditCardDetails : Represent credit card detail for PaymentTypeCreditCard payment type
type CreditCardDetails struct {
	// TokenID represents customer credit card information
	TokenID string `json:"token_id"`

	// Authentication Flag to enable the 3D secure authentication. Default value is false.
	Authentication bool `json:"authentication,omitempty"`

	// Bank Acquiring bank. Valid values: `midtrans.BankBca` `midtrans.BankMandiri`, `midtrans.BankBni`,
	//`midtrans.BankCimb`, `midtrans.BankMaybank`, and `midtrans.BankBri`
	Bank string `json:"bank,omitempty"`

	// InstallmentTerm for installment tenor
	InstallmentTerm int8 `json:"installment_term,omitempty"`

	// Type Used on preauthorization feature. Valid value: authorize
	Type string `json:"type,omitempty"`

	// Bins List of credit card's BIN (Bank Identification Number) that is allowed for transaction
	Bins []string `json:"bins,omitempty"`

	// SaveTokenID Used on 'one click' or 'two clicks' feature. Enabling it will return a `Response.SavedCardTokenID` on the response
	// and notification body that can be used for the next transaction
	SaveTokenID bool `json:"save_token_id,omitempty"`

	// PointRedeemAmount For Mandiri Point, you can only do Full Redemption.(use -1 for Full Redemption)
	PointRedeemAmount int64 `json:"point_redeem_amount,omitempty"`

	// Determines how the transaction status is updated to the merchant frontend. Possible values are js_event (default) and form
	CallbackType string `json:"callback_type,omitempty"`
}

// BankTransferDetails : Represent bank_transfer detail
type BankTransferDetails struct {
	Bank     midtrans.Bank                  `json:"bank"`
	VaNumber string                         `json:"va_number,omitempty"`
	Permata  *PermataBankTransferDetail     `json:"permata,omitempty"`
	FreeText *BCABankTransferDetailFreeText `json:"free_text,omitempty"`
	Bca      *BcaBankTransferDetail         `json:"bca,omitempty"`
}

// PermataBankTransferDetail : Represent Recipient for bank transfer Permata
type PermataBankTransferDetail struct {
	RecipientName string `json:"recipient_name,omitempty"`
}

// BCABankTransferDetailFreeText : Represent BCA bank_transfer detail free_text
type BCABankTransferDetailFreeText struct {
	Inquiry []BCABankTransferLangDetail `json:"inquiry,omitempty"`
	Payment []BCABankTransferLangDetail `json:"payment,omitempty"`
}

// BCABankTransferLangDetail : Represent BCA bank_transfer lang detail
type BCABankTransferLangDetail struct {
	LangID string `json:"id,omitempty"`
	LangEN string `json:"en,omitempty"`
}

// BcaBankTransferDetail : BCA sub company code directed for this transactions
// NOTE: Please contact Midtrans Sales Team.
type BcaBankTransferDetail struct {
	SubCompanyCode string `json:"sub_company_code,omitempty"`
}

// EChannelDetail : Represent Mandiri Bill bank transfer detail
type EChannelDetail struct {
	BillInfo1 string `json:"bill_info1"`
	BillInfo2 string `json:"bill_info2"`
	BillInfo3 string `json:"bill_info3,omitempty"`
	BillInfo4 string `json:"bill_info4,omitempty"`
	BillInfo5 string `json:"bill_info5,omitempty"`
	BillInfo6 string `json:"bill_info6,omitempty"`
	BillInfo7 string `json:"bill_info7,omitempty"`
	BillInfo8 string `json:"bill_info8,omitempty"`
	BillKey   string `json:"bill_key,omitempty"`
}

// BCAKlikPayDetails : Represent Internet Banking for BCA KlikPay
type BCAKlikPayDetails struct {
	Desc    string `json:"description"`
	MiscFee int64  `json:"misc_fee,omitempty"`
}

// BcaKlikBCADetails : Represent Internet Banking BCA KlikBCA detail
type BcaKlikBCADetails struct {
	Desc   string `json:"description"`
	UserID string `json:"user_id"`
}

// MandiriClickPayDetails : Represent Mandiri ClickPay detail
type MandiriClickPayDetails struct {
	// TokenID token id from Get card token Step
	TokenID string `json:"token_id"`
	Input1  string `json:"input1"`
	Input2  string `json:"input2"`

	// Input3 5-digits random number you gave to the customer
	Input3 string `json:"input3"`

	// Token Number generated by customer's physical token
	Token string `json:"token"`
}

// CIMBClicksDetails : Represent CIMB Clicks detail
type CIMBClicksDetails struct {
	Desc string `json:"description"`
}

// QrisDetails QRIS is a QR payment standard in Indonesia that is developed by Bank Indonesia (BI).
// Users could scan and pay the QR from any payment providers registered as the issuer
type QrisDetails struct {
	Acquirer string `json:"acquirer,omitempty"`
}

// ConvStoreDetails : Represent cstore detail
type ConvStoreDetails struct {
	Store   string `json:"store"`
	Message string `json:"message,omitempty"`

	AlfamartFreeText1 string `json:"alfamart_free_text_1,omitempty"`
	AlfamartFreeText2 string `json:"alfamart_free_text_2,omitempty"`
	AlfamartFreeText3 string `json:"alfamart_free_text_3,omitempty"`
}

// GopayDetails : Represent gopay detail
type GopayDetails struct {
	EnableCallback     bool   `json:"enable_callback,omitempty"`      // To determine appending callback url in the deeplink. Default value: false
	CallbackUrl        string `json:"callback_url,omitempty"`         // To determine where GO-JEK apps will redirect after successful payment. Can be HTTP or deeplink url. Default value: callback_url in dashboard settings
	AccountID          string `json:"account_id,omitempty"`           // Required for GoPay tokenization. Linked customer account ID from create pay account API.
	PaymentOptionToken string `json:"payment_option_token,omitempty"` // Required for GoPay tokenization. Token to specify the payment option made by the customer from get pay account API metadata.
	PreAuth            bool   `json:"pre_auth,omitempty"`             // To make payment mode into reservation of customer balance only. Once, customer balance is reserved, a subsequent capture call is expected to be initiated by merchants.
	Recurring          bool   `json:"recurring,omitempty"`
}

// ShopeePayDetails : Represent shopeepay detail
type ShopeePayDetails struct {
	CallbackUrl string `json:"callback_url,omitempty"`
}

// CustomExpiry : Represent Core API custom_expiry
type CustomExpiry struct {
	// OrderTime Time when the order is created in merchant website. Format: yyyy-MM-dd hh:mm:ss Z.
	// If attribute undefined, expiry time starts from transaction time
	OrderTime string `json:"order_time,omitempty"`

	// ExpiryDuration Time duration the payment will remain valid
	ExpiryDuration int `json:"expiry_duration,omitempty"`

	// Unit for expiry_duration. Valid values are: second, minute, hour, or day.
	// NOTE: If attribute undefined, default unit is minute
	Unit string `json:"unit,omitempty"`
}

// CaptureReq : Represent Capture request payload
type CaptureReq struct {
	TransactionID string  `json:"transaction_id"`
	GrossAmt      float64 `json:"gross_amount"`
}

// RefundReq : Represent Refund request payload
type RefundReq struct {
	RefundKey string `json:"refund_key"`
	Amount    int64  `json:"amount"`
	Reason    string `json:"reason"`
}

type SubscriptionReq struct {
	// Name Subscription's name that will be used to generate transaction's order id.
	// Note: Allowed symbols are dash(-), underscore(_), tilde (~), and dot (.)
	Name string `json:"name"`

	// Amount that will be used to make recurring charge. Note: Do not use decimal
	Amount int64 `json:"amount"`

	// Currency ISO-4217 representation for 3 digit alphabetic currency code. Note: Currently only support IDR
	Currency string `json:"currency"`

	// PaymentType Transaction payment method. Note: currently only support credit_card and gopay
	PaymentType SubscriptionPaymentType `json:"payment_type"`

	// Token Saved payment token. Note: For `credit_card` should use `saved_token_id` received in charge response.
	// For gopay should use payment_options. token received in get pay account response
	Token string `json:"token"`

	// Schedule Subscription schedule details
	Schedule ScheduleDetails `json:"schedule"`

	// Metadata of subscription from merchant, the size must be less than 1KB
	Metadata interface{} `json:"metadata,omitempty"`

	// CustomerDetails Customer details information
	CustomerDetails *midtrans.CustomerDetails `json:"customer_details,omitempty"`

	// Gopay subscription information, required if payment type is gopay
	Gopay *GopaySubscriptionDetails `json:"gopay,omitempty"`
}

type GopaySubscriptionDetails struct {
	AccountId string `json:"account_id"` // Gopay Account ID from Core API
}

//ScheduleDetails Create Subscription schedule object
type ScheduleDetails struct {
	// Subscription's interval given by merchant
	Interval int `json:"interval"`

	// Interval temporal unit Note: currently only support day, week, and month
	IntervalUnit string `json:"interval_unit"`

	// MaxInterval Maximum interval of subscription. Subscription will end after maximum interval is reached
	MaxInterval int `json:"max_interval"`

	// StartTime Timestamp of subscription, format: yyyy-MM-dd HH:mm:ss Z. The value must be after the current time.
	// If specified, first payment will happen on start_time. If start_time is not specified, the default value for
	// start_time will be current time and first payment will happen on one interval after current time.
	StartTime string `json:"start_time,omitempty"`
}

type PaymentAccountReq struct {
	PaymentType  CoreapiPaymentType   `json:"payment_type"`  // Payment channel where the account register to
	GopayPartner *GopayPartnerDetails `json:"gopay_partner"` // GoPay linking specific parameters
}

type GopayPartnerDetails struct {
	PhoneNumber string `json:"phone_number"`           // Phone number linked to the customer account
	CountryCode string `json:"country_code"`           // Country code associated to the phone number
	RedirectURL string `json:"redirect_url,omitempty"` // URL where user will be redirected to after finishing the confirmation on Gojek app
}
