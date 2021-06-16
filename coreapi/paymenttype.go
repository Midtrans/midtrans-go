package coreapi

type CoreapiPaymentType string
type SubscriptionPaymentType = CoreapiPaymentType

const (
	// PaymentTypeBankTransfer : bank_transfer
	PaymentTypeBankTransfer CoreapiPaymentType = "bank_transfer"

	// PaymentTypeGopay : gopay
	PaymentTypeGopay CoreapiPaymentType = "gopay"

	// PaymentTypeShopeepay : shopeepay
	PaymentTypeShopeepay CoreapiPaymentType = "shopeepay"

	// PaymentTypeQris : qris
	PaymentTypeQris CoreapiPaymentType = "qris"

	// PaymentTypeCreditCard : credit_card
	PaymentTypeCreditCard CoreapiPaymentType = "credit_card"

	// PaymentTypeEChannel : echannel
	PaymentTypeEChannel CoreapiPaymentType = "echannel"

	// PaymentTypeBCAKlikpay : bca_klikpay
	PaymentTypeBCAKlikpay CoreapiPaymentType = "bca_klikpay"

	// PaymentTypeKlikBca : bca_klikbca
	PaymentTypeKlikBca CoreapiPaymentType = "bca_klikbca"

	// PaymentTypeBRIEpay : bri_epay
	PaymentTypeBRIEpay CoreapiPaymentType = "bri_epay"

	// PaymentTypeCimbClicks : cimb_clicks
	PaymentTypeCimbClicks CoreapiPaymentType = "cimb_clicks"

	// PaymentTypeDanamonOnline : danamon_online
	PaymentTypeDanamonOnline CoreapiPaymentType = "danamon_online"

	// PaymentTypeConvenienceStore : cstore
	PaymentTypeConvenienceStore CoreapiPaymentType = "cstore"

	// PaymentTypeAkulaku : akulaku
	PaymentTypeAkulaku CoreapiPaymentType = "akulaku"

	// PaymentTypeMandiriClickpay : mandiri_clickpay
	PaymentTypeMandiriClickpay CoreapiPaymentType = "mandiri_clickpay"
)
