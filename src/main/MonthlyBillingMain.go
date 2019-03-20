package main

import (
	"net/http"
	"google.golang.org/appengine"
	"time"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
	"strconv"
)

func MonthlyBillingMain(w http.ResponseWriter, req *http.Request) {
	basicCtx := appengine.NewContext(req)
	ctx, _ := context.WithTimeout(basicCtx, 60*time.Second)

	// init
	log.Infof(ctx, "Now initializing...")
	initConfig()

	// データを取得
	log.Infof(ctx, "Data Loading...")
	billingList, loadErr := loadUserBilling(ctx)
	if loadErr != nil {
		log.Errorf(ctx, loadErr.Error())
		return
	}

	log.Infof(ctx, "件数:" + strconv.Itoa(len(billingList)))

	// 一行づつ請求を行う
	log.Infof(ctx, "Billing...")
	for _, bill := range billingList {
		log.Infof(ctx, " value:" + bill.Name)
		billUser(ctx, bill)
	}

	log.Infof(ctx, "Make Earnings...")
	monthlyReport(ctx, billingList)
	// 請求結果を売ってる人に送る

	log.Infof(ctx, "DONE !!!!")
}
