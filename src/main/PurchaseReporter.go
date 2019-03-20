package main
import (
	"golang.org/x/net/context"
	"strconv"
)

func reportPurchaseOnSlack(ctx context.Context, pd PurchaseDetail){
	reporter := Reporter{config.Slack.Name, config.Slack.Icon, "お買い上げありがとうございます"}
	sendNlope(ctx, makePurchaseMessage(pd), reporter)
}

func makePurchaseMessage(pd PurchaseDetail)(string){
	result := "おなまえ: " + pd.User + " 様\n" +
	"購入金額 : " + strconv.Itoa(int(pd.Price)) + "円 \n"

	return result
}
