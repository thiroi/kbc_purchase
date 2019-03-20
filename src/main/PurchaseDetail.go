package main

import (
	"google.golang.org/api/iterator"
	"google.golang.org/appengine/log"
	"golang.org/x/net/context"
	"net/http"
	"strconv"
	"io"
	"encoding/json"
	"fmt"
	"errors"
	"time"
	"cloud.google.com/go/bigquery"
)


type PurchaseDetail struct {
	User string `json:"user"`
	Price int64 `json:"price"`
	CreatedAt time.Time `json:"createdAt"`
}

var defaultDetail = PurchaseDetail {"error", 3500, time.Now()}

func toPurchaseFromJson(ctx context.Context, req *http.Request)(PurchaseDetail, error) {
	if req.Method != "POST" {
		return defaultDetail, errors.New("Request MethodがPOSTではありません")
	}

	if req.Header.Get("Content-Type") != "application/json" {
		return defaultDetail, errors.New("Content-Typeが異常です")
	}

	//To allocate slice for request body
	length, err := strconv.Atoi(req.Header.Get("Content-Length"))
	if err != nil {
		return defaultDetail, errors.New("Content-Lengthが異常です")
	}

	//Read body data to parse json
	body := make([]byte, length)
	length, err = req.Body.Read(body)
	if err != nil && err != io.EOF {
		return defaultDetail, errors.New("Bodyの取得に失敗しました")
	}

	var pd PurchaseDetail
	err = json.Unmarshal([]byte(body), &pd)
	if err != nil {
		log.Errorf(ctx, err.Error())
		return defaultDetail, errors.New("Bodyの取得に失敗しました")
	}
	fmt.Printf("%v\n", pd)

	result := PurchaseDetail{pd.User, pd.Price, time.Now()}
	log.Infof(ctx, "USER:" + result.User)

	return result , nil
}

func insertPurchaseDetails(ctx context.Context, pd PurchaseDetail)(error){

	//var pdList []PurchaseDetail
	//pdList = append(pdList, pd)
	client, err := bigquery.NewClient(ctx, config.Bq.Project)
	if err != nil {
		return err
	}
	defer client.Close()
	u := client.Dataset("purchase").Table("detail").Uploader()

	err = u.Put(ctx, pd)
	if err != nil {
		if multiError, ok := err.(bigquery.PutMultiError); ok {
			for _, err1 := range multiError {
				for _, err2 := range err1.Errors {
					log.Errorf(ctx, "ERR2:" + err2.Error())
				}
				log.Errorf(ctx, "ERR1:" + err1.Error())
			}
		} else {
			log.Errorf(ctx, "ERR:" + err.Error())
		}
		return err
	}

	return nil
}


const GET_MONTHLY_SUMMARY = ""
func loadPurchaseDetails(ctx context.Context)([]PurchaseDetail, error){
	it, err := readQuery(ctx, GET_MONTHLY_SUMMARY)
	if err != nil {
		return nil, err
	}
	var result []PurchaseDetail
	for {
		var line PurchaseDetail
		err := it.Next(&line)
		if err == iterator.Done {
			break
		}
		if err != nil {
			break
		}
		result = append(result, line)
	}
	return result, nil
}