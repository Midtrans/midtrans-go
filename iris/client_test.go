package iris

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	assert "github.com/stretchr/testify/require"
	"math/rand"
	"strconv"
	"testing"
	"time"
)

var irisCreatorKeySandbox = "IRIS-330198f0-e49d-493f-baae-585cfded355d"
var irisApproverKeySandbox = "IRIS-1595c12b-6814-4e5a-bbbb-9bc18193f47b"

func random() string {
	rand.Seed(time.Now().UnixNano())
	return strconv.Itoa(rand.Intn(2000-1000) + 100000)
}

func generateDate() (string, string) {
	t := time.Now()
	var fromDate = t.AddDate(0, -1, 0).Format("2006-01-02")
	var toDate = t.Format("2006-01-02")
	return fromDate, toDate
}

func mockBeneficiaries() Beneficiaries {
	var random = random()
	return Beneficiaries{
		Name:      "MidGoUnitTest" + random,
		Account:   random,
		Bank:      "bca",
		AliasName: "midgotest" + random,
		Email:     "midgo" + random + "@mail.com",
	}
}

func TestGetBalance(t *testing.T) {
	var g = Gateway{}
	g.New(irisCreatorKeySandbox, midtrans.Sandbox)
	resp2, err2 := g.Iris.GetBalance()
	assert.Nil(t, err2)
	assert.NotNil(t, resp2)
}

func TestCreateAndUpdateBeneficiaries(t *testing.T) {
	g := Gateway{}
	g.New(irisCreatorKeySandbox, midtrans.Sandbox)

	newBeneficiaries := mockBeneficiaries()
	resp1, err1 := g.Iris.CreateBeneficiaries(newBeneficiaries)
	assert.Nil(t, err1)
	assert.Equal(t, resp1.Status, "created")

	getListAndUpdateBeneficiaries(t, newBeneficiaries)
}

func getListAndUpdateBeneficiaries(t *testing.T, beneficiaries Beneficiaries) {
	g := Gateway{}
	g.New(irisCreatorKeySandbox, midtrans.Sandbox)
	beneficiariesList, _ := g.Iris.GetBeneficiaries()

	b := Beneficiaries{}
	for _, account := range beneficiariesList {
		if account.AliasName == beneficiaries.AliasName {
			b = account
			break
		}
	}

	updateBeneficiaries := Beneficiaries{
		Name:      b.Name,
		Account:   b.Account,
		Bank:      b.Bank,
		AliasName: b.AliasName + "edt",
		Email:     b.Email,
	}

	resp, _ := g.Iris.UpdateBeneficiaries(b.AliasName, updateBeneficiaries)
	assert.Equal(t, resp.Status, "updated")
}

func createPayout() []CreatePayoutDetailResponse {
	p := CreatePayoutDetailReq{
		BeneficiaryName:    "Tony Stark",
		BeneficiaryAccount: "1380011819286",
		BeneficiaryBank:    "mandiri",
		BeneficiaryEmail:   "tony.stark@mail.com",
		Amount:             random(),
		Notes:              "MidGoUnitTestApproved",
	}
	var payouts []CreatePayoutDetailReq
	payouts = append(payouts, p)

	cp := CreatePayoutReq{Payouts: payouts}

	g := Gateway{}
	g.New(irisCreatorKeySandbox, midtrans.Sandbox)
	payoutReps, err := g.Iris.CreatePayout(cp)
	fmt.Println(payoutReps, err)

	return payoutReps.Payouts
}

func getPayoutDetails(refNo string) string {
	g := Gateway{}
	g.New(irisCreatorKeySandbox, midtrans.Sandbox)
	payoutReps, err := g.Iris.GetPayoutDetails(refNo)
	fmt.Println(payoutReps, err)
	return payoutReps.ReferenceNo
}

func TestCreateAndApprovePayout(t *testing.T) {
	var payouts = createPayout()
	assert.Equal(t, payouts[0].Status, "queued")

	var refNos []string
	refNos = append(refNos, payouts[0].ReferenceNo)

	ap := ApprovePayoutReq{
		ReferenceNo: refNos,
		OTP:         "335163",
	}
	gw := Gateway{}
	gw.New(irisApproverKeySandbox, midtrans.Sandbox)
	approveResp, err2 := gw.Iris.ApprovePayout(ap)
	assert.Nil(t, err2)
	assert.Equal(t, approveResp.Status, "ok")

	assert.Equal(t, getPayoutDetails(payouts[0].ReferenceNo), payouts[0].ReferenceNo)
}

func TestCreateAndRejectPayout(t *testing.T) {
	var payouts = createPayout()
	assert.Equal(t, payouts[0].Status, "queued")

	var refNos []string
	refNos = append(refNos, payouts[0].ReferenceNo)

	ap := RejectPayoutReq{
		ReferenceNo:  refNos,
		RejectReason: "MidGoUnitTest",
	}
	gw := Gateway{}
	gw.New(irisApproverKeySandbox, midtrans.Sandbox)
	approveResp, err2 := gw.Iris.RejectPayout(ap)
	assert.Nil(t, err2)
	assert.Equal(t, approveResp.Status, "ok")
}

func TestPayoutHistory(t *testing.T)  {
	fromDate, toDate := generateDate()

	gw := Gateway{}
	gw.New(irisApproverKeySandbox, midtrans.Sandbox)
	resp, err := gw.Iris.GetTransactionHistory(fromDate, toDate)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}


func TestGetTopUpChannels(t *testing.T)  {
	gw := Gateway{}
	gw.New(irisApproverKeySandbox, midtrans.Sandbox)
	resp, err := gw.Iris.GetTopUpChannels()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}

func TestGetListBeneficiaryBank(t *testing.T)  {
	gw := Gateway{}
	gw.New(irisApproverKeySandbox, midtrans.Sandbox)
	resp, err := gw.Iris.GetBeneficiaryBanks()
	assert.Nil(t, err)
	assert.NotNil(t, resp)
}


func TestValidateBankAccount(t *testing.T)  {
	gw := Gateway{}
	gw.New(irisApproverKeySandbox, midtrans.Sandbox)
	resp, err := gw.Iris.ValidateBankAccount("danamon", "000001137298")
	assert.Nil(t, err)
	assert.Equal(t, resp.AccountNo, "000001137298")
}