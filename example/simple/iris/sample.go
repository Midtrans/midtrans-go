package main

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/example"
	"github.com/midtrans/midtrans-go/iris"
	"math/rand"
	"strconv"
	"time"
)

var irisCreator iris.Client
var irisApprover iris.Client

func setupIrisGateway() {
	irisCreator.New(example.IrisCreatorKeySandbox, midtrans.Sandbox)
	irisApprover.New(example.IrisApproverKeySandbox, midtrans.Sandbox)
}

func GetBalance() {
	res, err := irisCreator.GetBalance()
	if err != nil {
		fmt.Println("Error: ", err.GetMessage())
	}
	fmt.Println("Response: ", res)
}

func CreateAndUpdateBeneficiaries() {
	newBeneficiaries := mockBeneficiaries()
	resp, _ := irisCreator.CreateBeneficiaries(newBeneficiaries)
	fmt.Println("Iris Create beneficiary response: ", resp)

	getListAndUpdateBeneficiaries(newBeneficiaries)
}

func getListAndUpdateBeneficiaries(beneficiaries iris.Beneficiaries) {
	beneficiariesList, _ := irisCreator.GetBeneficiaries()

	b := iris.Beneficiaries{}
	for _, account := range beneficiariesList {
		if account.AliasName == beneficiaries.AliasName {
			b = account
			break
		}
	}

	updateBeneficiaries := iris.Beneficiaries{
		Name:      b.Name,
		Account:   b.Account,
		Bank:      b.Bank,
		AliasName: b.AliasName + "edt",
		Email:     b.Email,
	}

	resp, _ := irisCreator.UpdateBeneficiaries(b.AliasName, updateBeneficiaries)
	fmt.Println("Iris Update Beneficiary: ", resp)
}

func createPayout() []iris.CreatePayoutDetailResponse {
	p := iris.CreatePayoutDetailReq{
		BeneficiaryName:    "Tony Stark",
		BeneficiaryAccount: "1380011819286",
		BeneficiaryBank:    "mandiri",
		BeneficiaryEmail:   "tony.stark@mail.com",
		Amount:             random(),
		Notes:              "MidGoUnitTestApproved",
	}
	var payouts []iris.CreatePayoutDetailReq
	payouts = append(payouts, p)

	cp := iris.CreatePayoutReq{Payouts: payouts}

	payoutReps, err := irisCreator.CreatePayout(cp)
	fmt.Println(payoutReps, err)

	return payoutReps.Payouts
}

func GetPayoutDetails(refNo string) {
	payoutReps, _ := irisCreator.GetPayoutDetails(refNo)
	fmt.Println("Iris Payout details", payoutReps)
}

func CreateAndApprovePayout() {
	var payouts = createPayout()

	var refNos []string
	refNos = append(refNos, payouts[0].ReferenceNo)

	ap := iris.ApprovePayoutReq{
		ReferenceNo: refNos,
		OTP:         "335163",
	}

	approveResp, _ := irisApprover.ApprovePayout(ap)
	fmt.Println("Iris Approve payout resp: ", approveResp)
}

func CreateAndRejectPayout() {
	var payouts = createPayout()

	var refNos []string
	refNos = append(refNos, payouts[0].ReferenceNo)

	ap := iris.RejectPayoutReq{
		ReferenceNo:  refNos,
		RejectReason: "MidGoUnitTest",
	}

	approveResp, _ := irisApprover.RejectPayout(ap)
	fmt.Println("Iris reject payout resp: ", approveResp)
}

func PayoutHistory() {
	fromDate, toDate := generateDate()

	resp, _ := irisApprover.GetTransactionHistory(fromDate, toDate)
	fmt.Println("Iris Payout history: ", resp)
}

func GetTopUpChannels() {
	resp, _ := irisApprover.GetTopUpChannels()
	fmt.Println("Iris TopUp Channels resp: ", resp)
}

func GetListBeneficiaryBank() {
	resp, _ := irisApprover.GetBeneficiaryBanks()
	fmt.Println("Iris Beneficiary Banks Resp: ", resp)
}

func ValidateBankAccount() {
	resp, _ := irisApprover.ValidateBankAccount("danamon", "000001137298")
	fmt.Println("Validate Bank Account Resp: ", resp)
}

func main() {
	setupIrisGateway()

	GetBalance()

	CreateAndUpdateBeneficiaries()

	GetPayoutDetails("ssss")

	CreateAndApprovePayout()

	CreateAndRejectPayout()

	PayoutHistory()

	GetTopUpChannels()

	GetListBeneficiaryBank()

	ValidateBankAccount()
}

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

func mockBeneficiaries() iris.Beneficiaries {
	var random = random()
	return iris.Beneficiaries{
		Name:      "MidGoUnitTest" + random,
		Account:   random,
		Bank:      "bca",
		AliasName: "midgotest" + random,
		Email:     "midgo" + random + "@mail.com",
	}
}
