package main
import (
	"golang.org/x/net/context"
	"strconv"
)

func billUser(ctx context.Context, bill UserBilling){
	reporter := Reporter{config.Slack.Name, config.Slack.Icon, "まいど！今月の支払い金額のお知らせです！"}
	//sendNlopeWithChannel(ctx, makeBillMessage(bill), reporter, "U4B4P5Q56")
	sendNlopeWithChannel(ctx, makeBillMessage(bill), reporter, bill.SlackMemberId)
}

func monthlyReport(ctx context.Context, bills []UserBilling){
	reporter := Reporter{config.Slack.Name, config.Slack.Icon, "今月の売り上げ合計金額は〜〜〜〜！？"}
	sendNlope(ctx, makeMonthlyEarningMessage(bills), reporter)
	//sendNlopeWithChannel(ctx, makeMonthlyEarningMessage(bills), reporter, "U4B4P5Q56")
}

func makeBillMessage(bill UserBilling)(string){
	result := "" + bill.Name + "様の今月のお支払い詳細\n" +
		"購入回数 : " + strconv.Itoa(int(bill.PurchaseCount)) + "回 \n" +
		"▪▪合計金額▪▪ : " + strconv.Itoa(int(bill.TotalPrice)) + "円 \n"
	return result
}

func makeMonthlyEarningMessage(bills []UserBilling)(string){
	var totalEarnings = 0
	var message = ""
	// 一行づつ請求を行う
	for _, bill := range bills {
		totalEarnings = totalEarnings + int(bill.TotalPrice)
		message = message + makeOneLineEarning(bill)
	}
	return "合計金額：" + strconv.Itoa(totalEarnings) + "円 \n\n" +
		"お客様ごと売り上げ\n" +
		"---------------------------------------- \n" +
		"" + message
}

func makeOneLineEarning(bill UserBilling)(string){
	return bill.Name + " 様" + " " + strconv.Itoa(int(bill.TotalPrice)) + "円 \n"
}