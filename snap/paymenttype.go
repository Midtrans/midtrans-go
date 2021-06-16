package snap

type SnapPaymentType string

const (
	// PaymentTypeCreditCard : credit_card
	PaymentTypeCreditCard SnapPaymentType = "credit_card"

	// PaymentTypeMandiriClickpay : mandiri_clickpay
	PaymentTypeMandiriClickpay SnapPaymentType = "mandiri_clickpay"

	// PaymentTypeCimbClicks : cimb_clicks
	PaymentTypeCimbClicks SnapPaymentType = "cimb_clicks"

	// PaymentTypeKlikBca : bca_klikbca
	PaymentTypeKlikBca SnapPaymentType = "bca_klikbca"

	// PaymentTypeBCAKlikpay : bca_klikpay
	PaymentTypeBCAKlikpay SnapPaymentType = "bca_klikpay"

	// PaymentTypeBRIEpay : bri_epay
	PaymentTypeBRIEpay SnapPaymentType = "bri_epay"

	// PaymentTypeTelkomselCash : telkomsel_cash
	PaymentTypeTelkomselCash SnapPaymentType = "telkomsel_cash"

	// PaymentTypeEChannel : echannel
	PaymentTypeEChannel SnapPaymentType = "echannel"

	// PaymentTypeMandiriEcash : mandiri_ecash
	PaymentTypeMandiriEcash SnapPaymentType = "mandiri_ecash"

	// PaymentTypePermataVA : permata_va
	PaymentTypePermataVA SnapPaymentType = "permata_va"

	// PaymentTypeOtherVA : other_va If you want to use other_va, either permata_va or bni_va
	// because Midtrans handles other bank transfer as either Permata or BNI VA.
	PaymentTypeOtherVA SnapPaymentType = "other_va"

	// PaymentTypeBCAVA : bca_va
	PaymentTypeBCAVA SnapPaymentType = "bca_va"

	// PaymentTypeBNIVA : bni_va
	PaymentTypeBNIVA SnapPaymentType = "bni_va"

	// PaymentTypeBRIVA : bca_va
	PaymentTypeBRIVA SnapPaymentType = "bri_va"

	// PaymentTypeBankTransfer : bank_transfer
	PaymentTypeBankTransfer SnapPaymentType = "bank_transfer"

	// PaymentTypeConvenienceStore : cstore
	PaymentTypeConvenienceStore SnapPaymentType = "cstore"

	// PaymentTypeIndomaret : indomaret
	PaymentTypeIndomaret SnapPaymentType = "indomaret"

	// PaymentTypeKioson : kioson
	PaymentTypeKioson SnapPaymentType = "kioson"

	// PaymentTypeDanamonOnline : danamon_online
	PaymentTypeDanamonOnline SnapPaymentType = "danamon_online"

	// PaymentTypeAkulaku : akulaku
	PaymentTypeAkulaku SnapPaymentType = "akulaku"

	// PaymentTypeGopay : gopay
	PaymentTypeGopay SnapPaymentType = "gopay"

	// PaymentTypeShopeepay : shopeepay
	PaymentTypeShopeepay SnapPaymentType = "shopeepay"

	// PaymentTypeAlfamart : alfamart
	PaymentTypeAlfamart SnapPaymentType = "alfamart"
)

// AllSnapPaymentType : Get All available SnapPaymentType
var AllSnapPaymentType = []SnapPaymentType{
	PaymentTypeGopay,
	PaymentTypeShopeepay,
	PaymentTypeCreditCard,
	PaymentTypeBankTransfer,
	PaymentTypeBNIVA,
	PaymentTypePermataVA,
	PaymentTypeBCAVA,
	PaymentTypeBRIVA,
	PaymentTypeOtherVA,
	PaymentTypeMandiriClickpay,
	PaymentTypeCimbClicks,
	PaymentTypeDanamonOnline,
	PaymentTypeKlikBca,
	PaymentTypeBCAKlikpay,
	PaymentTypeBRIEpay,
	PaymentTypeMandiriEcash,
	PaymentTypeTelkomselCash,
	PaymentTypeEChannel,
	PaymentTypeIndomaret,
	PaymentTypeKioson,
	PaymentTypeAkulaku,
	PaymentTypeAlfamart,
	PaymentTypeConvenienceStore,
}
