package main

import (
	"net/http"
	"google.golang.org/appengine"
	"time"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
	"encoding/json"
)

func PurchaseMain(w http.ResponseWriter, req *http.Request) {
	basicCtx := appengine.NewContext(req)
	ctx, _ := context.WithTimeout(basicCtx, 60*time.Second)

	// init
	log.Infof(ctx, "Now initializing...")
	initConfig()

	// decode
	log.Infof(ctx, "Decoding...")
	pd, err := toPurchaseFromJson(ctx, req)
	if err != nil {
		log.Errorf(ctx, err.Error())
		result, _ := json.Marshal(makeError(err))
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
		return
	}

	// dbにインサート
	log.Infof(ctx, "Inserting DB...")
	dbErr := insertPurchaseDetails(ctx, pd)
	if dbErr != nil {
		log.Errorf(ctx, dbErr.Error())
		result, _ := json.Marshal(makeError(dbErr))
		w.Header().Set("Content-Type", "application/json")
		w.Write(result)
		return
	}

	// 正しいResponse返す
	result, _ := json.Marshal(makeSuccess())
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)

	// 必要に応じてslack通知
	log.Infof(ctx, "Sending slack")
	reportPurchaseOnSlack(ctx, pd)

	log.Infof(ctx, "DONE!!!")
}