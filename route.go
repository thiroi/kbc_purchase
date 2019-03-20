package src

import (
	"net/http"
	"src/main"
)

func init() {
// net/http
http.HandleFunc("/purchase", main.PurchaseMain)
http.HandleFunc("/monthly_billing", main.MonthlyBillingMain)
}
