package main
import (
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"log"
)

type UserBilling struct {
	Name string `json:"name"`
	SlackMemberId string `json:"slackMemberId"`
	PurchaseCount int64 `json:"purchaseCount"`
	TotalPrice int64 `json:"totalPrice"`
}

const MONTHLY_BILLING_SQL = "SELECT u.name, u.slackMemberId, COUNT(1) purchaseCount, SUM(pd.price) totalPrice FROM `<project>.purchase.detail` pd INNER JOIN`<project>.purchase.user` u ON pd.user = u.name WHERE TIMESTAMP_TRUNC(pd.createdAt, MONTH, 'Asia/Tokyo') =   TIMESTAMP_TRUNC(CURRENT_TIMESTAMP(), MONTH, 'Asia/Tokyo') GROUP BY u.name, u.slackMemberId"
func loadUserBilling (ctx context.Context)([]UserBilling, error){
	it, err := readQuery(ctx, MONTHLY_BILLING_SQL)
	if err != nil {
		return nil, err
	}
	var result []UserBilling
	for {
		var line UserBilling
		err := it.Next(&line)
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Println("Failed to Iterate Query:%v", err)
		}
		result = append(result, line)
	}
	return result, nil
}
