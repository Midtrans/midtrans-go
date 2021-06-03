package midtrans

type PaymentType string

const (
	// PaymentTypeBankTransfer : bank_transfer
	PaymentTypeBankTransfer PaymentType = "bank_transfer"

	// PaymentTypeBNIVA : bni_va
	PaymentTypeBNIVA PaymentType = "bni_va"

	// PaymentTypePermataVA : permata_va
	PaymentTypePermataVA PaymentType = "permata_va"

	// PaymentTypeBCAVA : bca_va
	PaymentTypeBCAVA PaymentType = "bca_va"

	// PaymentTypeBRIVA : bca_va
	PaymentTypeBRIVA PaymentType = "bri_va"

	// PaymentTypeOtherVA : other_va
	PaymentTypeOtherVA PaymentType = "other_va"

	// PaymentTypeBCAKlikpay : bca_klikpay
	PaymentTypeBCAKlikpay PaymentType = "bca_klikpay"

	// PaymentTypeBRIEpay : bri_epay
	PaymentTypeBRIEpay PaymentType = "bri_epay"

	// PaymentTypeCreditCard : credit_card
	PaymentTypeCreditCard PaymentType = "credit_card"

	// PaymentTypeCimbClicks : cimb_clicks
	PaymentTypeCimbClicks PaymentType = "cimb_clicks"

	// PaymentTypeDanamonOnline : danamon_online
	PaymentTypeDanamonOnline PaymentType = "danamon_online"

	// PaymentTypeConvenienceStore : cstore
	PaymentTypeConvenienceStore PaymentType = "cstore"

	// PaymentTypeKlikBca : bca_klikbca
	PaymentTypeKlikBca PaymentType = "bca_klikbca"

	// PaymentTypeEChannel : echannel
	PaymentTypeEChannel PaymentType = "echannel"

	// PaymentTypeMandiriClickpay : mandiri_clickpay
	PaymentTypeMandiriClickpay PaymentType = "mandiri_clickpay"

	// PaymentTypeTelkomselCash : telkomsel_cash
	PaymentTypeTelkomselCash PaymentType = "telkomsel_cash"

	// PaymentTypeIndosatDompetku : indosat_dompetku
	PaymentTypeIndosatDompetku PaymentType = "indosat_dompetku"

	// PaymentTypeMandiriEcash : mandiri_ecash
	PaymentTypeMandiriEcash PaymentType = "mandiri_ecash"

	// PaymentTypeKioson : kioson
	PaymentTypeKioson PaymentType = "kioson"

	// PaymentTypeIndomaret : indomaret
	PaymentTypeIndomaret PaymentType = "indomaret"

	// PaymentTypeAlfamart : alfamart
	PaymentTypeAlfamart PaymentType = "alfamart"

	// PaymentTypeGiftCardIndo : gci
	PaymentTypeGiftCardIndo PaymentType = "gci"

	// PaymentTypeGopay : gopay
	PaymentTypeGopay PaymentType = "gopay"

	// PaymentTypeShopeepay : shopeepay
	PaymentTypeShopeepay PaymentType = "shopeepay"

	// PaymentTypeQris : qris
	PaymentTypeQris PaymentType = "qris"

	// PaymentTypeAkulaku : akulaku
	PaymentTypeAkulaku PaymentType = "akulaku"
)

// AllPaymentType : Get All available PaymentType
var AllPaymentType = []PaymentType{
	PaymentTypeGopay,
	PaymentTypeShopeepay,
	PaymentTypeQris,
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
	PaymentTypeIndosatDompetku,
	PaymentTypeEChannel,
	PaymentTypeIndomaret,
	PaymentTypeKioson,
	PaymentTypeGiftCardIndo,
	PaymentTypeAkulaku,
	PaymentTypeAlfamart,
	PaymentTypeConvenienceStore,
}
