package midtrans

type PaymentType string

const (
	// SourceBankTransfer : bank_transfer
	SourceBankTransfer PaymentType = "bank_transfer"

	// SourceBNIVA : bni_va
	SourceBNIVA PaymentType = "bni_va"

	// PermataVA : permata_va
	PermataVA PaymentType = "permata_va"

	// SourceBCAVA : bca_va
	SourceBCAVA PaymentType = "bca_va"

	// SourceBRIVA : bca_va
	SourceBRIVA PaymentType = "bri_va"

	// SourceOtherVA : other_va
	SourceOtherVA PaymentType = "other_va"

	// SourceBCAKlikpay : bca_klikpay
	SourceBCAKlikpay PaymentType = "bca_klikpay"

	// SourceBRIEpay : bri_epay
	SourceBRIEpay PaymentType = "bri_epay"

	// SourceCreditCard : credit_card
	SourceCreditCard PaymentType = "credit_card"

	// SourceCimbClicks : cimb_clicks
	SourceCimbClicks PaymentType = "cimb_clicks"

	// SourceDanamonOnline : danamon_online
	SourceDanamonOnline PaymentType = "danamon_online"

	// SourceConvenienceStore : cstore
	SourceConvenienceStore PaymentType = "cstore"

	// SourceKlikBca : bca_klikbca
	SourceKlikBca PaymentType = "bca_klikbca"

	// SourceEChannel : echannel
	SourceEChannel PaymentType = "echannel"

	// SourceMandiriClickpay : mandiri_clickpay
	SourceMandiriClickpay PaymentType = "mandiri_clickpay"

	// SourceTelkomselCash : telkomsel_cash
	SourceTelkomselCash PaymentType = "telkomsel_cash"

	// SourceIndosatDompetku : indosat_dompetku
	SourceIndosatDompetku PaymentType = "indosat_dompetku"

	// SourceMandiriEcash : mandiri_ecash
	SourceMandiriEcash PaymentType = "mandiri_ecash"

	// SourceKioson : kioson
	SourceKioson PaymentType = "kioson"

	// SourceIndomaret : indomaret
	SourceIndomaret PaymentType = "indomaret"

	// SourceAlfamart : alfamart
	SourceAlfamart PaymentType = "alfamart"

	// SourceGiftCardIndo : gci
	SourceGiftCardIndo PaymentType = "gci"

	// SourceGopay : gopay
	SourceGopay PaymentType = "gopay"

	// SourceShopeepay : shopeepay
	SourceShopeepay PaymentType = "shopeepay"

	// SourceQris : qris
	SourceQris PaymentType = "qris"

	// SourceAkulaku : akulaku
	SourceAkulaku PaymentType = "akulaku"
)

// AllPaymentSource : Get All available PaymentType
var AllPaymentSource = []PaymentType{
	SourceGopay,
	SourceShopeepay,
	SourceCreditCard,
	SourceBankTransfer,
	SourceBNIVA,
	PermataVA,
	SourceBCAVA,
	SourceBRIVA,
	SourceOtherVA,
	SourceMandiriClickpay,
	SourceCimbClicks,
	SourceDanamonOnline,
	SourceKlikBca,
	SourceBCAKlikpay,
	SourceBRIEpay,
	SourceMandiriEcash,
	SourceTelkomselCash,
	SourceIndosatDompetku,
	SourceEChannel,
	SourceIndomaret,
	SourceKioson,
	SourceGiftCardIndo,
	SourceAkulaku,
	SourceAlfamart,
	SourceConvenienceStore,
}
