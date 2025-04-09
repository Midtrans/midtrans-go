package main

import (
	"fmt"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/example"
	"github.com/midtrans/midtrans-go/snap"
)

var client snap.Client

func initializeClient() {
	client.New(example.SandboxServerKey1, midtrans.Sandbox)
}

func createGopayTokenizationTransaction() {
	// Create a transaction request with GoPay tokenization enabled
	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  "GopayLinking-" + example.Random(),
			GrossAmt: 10000,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: "John",
			LName: "Doe",
			Email: "john@doe.com",
			Phone: "081234567890",
		},
		EnabledPayments: []snap.SnapPaymentType{snap.PaymentTypeGopay},
		Gopay: &snap.GopayDetails{
			EnableCallback: true,
			CallbackUrl:   "https://example.com/gopay/callback",
			Tokenization:  true,
			PhoneNumber:   "081234567890",
			CountryCode:   "62",
		},
		UserId: "user-" + example.Random(), // Important: Store this user_id to manage the linked account
	}

	// Create transaction to get Snap token
	resp, err := client.CreateTransaction(req)
	if err != nil {
		fmt.Println("Error creating transaction:", err.GetMessage())
		return
	}

	fmt.Println("Transaction created successfully")
	fmt.Println("Snap Token:", resp.Token)
	fmt.Println("Redirect URL:", resp.RedirectURL)
	fmt.Println("User ID (save this):", req.UserId)
}

func getGopayAccountInfo(userID string) {
	// Get GoPay account information for a linked account
	resp, err := client.GetGopayAccount(userID)
	if err != nil {
		fmt.Println("Error getting account info:", err.GetMessage())
		return
	}

	fmt.Println("GoPay Account Information:")
	fmt.Printf("User ID: %s\n", resp.UserID)
	fmt.Printf("Account Status: %s\n", resp.AccountStatus)
	fmt.Printf("Balance: %s\n", resp.Balance)
	fmt.Printf("Phone Number: %s\n", resp.PhoneNumber)
	fmt.Printf("Country Code: %s\n", resp.CountryCode)
	fmt.Printf("Account ID: %s\n", resp.AccountID)
}

func unlinkGopayAccount(userID string) {
	// Unlink a GoPay account
	err := client.UnlinkGopayAccount(userID)
	if err != nil {
		fmt.Println("Error unlinking account:", err.GetMessage())
		return
	}

	fmt.Println("GoPay account unlinked successfully")
}

func main() {
	// Initialize the Snap client
	initializeClient()

	// Example 1: Create a transaction with GoPay tokenization
	fmt.Println("=== Creating Transaction with GoPay Tokenization ===")
	createGopayTokenizationTransaction()

	// For demonstration, we'll use a test user ID
	// In real implementation, you should store and reuse the user_id from createGopayTokenizationTransaction
	testUserID := "test-user-id"

	// Example 2: Get GoPay account information
	fmt.Println("\n=== Getting GoPay Account Information ===")
	getGopayAccountInfo(testUserID)

	// Example 3: Unlink GoPay account
	fmt.Println("\n=== Unlinking GoPay Account ===")
	unlinkGopayAccount(testUserID)
} 