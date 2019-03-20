package main

import (
"fmt"
"net/http"
"google.golang.org/appengine/log"
"bytes"
"golang.org/x/text/encoding/japanese"
"golang.org/x/text/transform"
"io"
"io/ioutil"
"time"
"github.com/nlopes/slack"
"google.golang.org/appengine/urlfetch"
"golang.org/x/net/context"
)

func SlackSender(w http.ResponseWriter, ctx *http.Request) {
	fmt.Fprint(w, "Hello, world222!")
}



func sendNlope(ctx context.Context, message string, reporter Reporter){
	sendNlopeWithChannel(ctx, message, reporter, config.Slack.Channel)
}

func sendNlopeWithChannel(ctx context.Context, message string, reporter Reporter, channel string){
	log.Infof(ctx, "Start sending with nlope")
	slack.SetHTTPClient(urlfetch.Client(ctx))
	api := slack.New(config.Slack.Token)
	params := slack.PostMessageParameters{}
	params.IconEmoji = reporter.Icon
	params.Username = reporter.Name
	message = reporter.Talk + "\n\n" + "======================================== \n" + message + "========================================"
	log.Infof(ctx, message)
	channelID, timestamp, err := api.PostMessage(channel, message, params)
	if err != nil {
		log.Errorf(ctx, "%s\n", err)
		return
	}
	log.Infof(ctx, "Message successfully sent to channel %s at %s", channelID, timestamp)
}

func BytesToShiftJIS(b []byte) (string, error) {
	return transformEncoding(bytes.NewReader(b), japanese.ShiftJIS.NewEncoder())
}

func transformEncoding( rawReader io.Reader, trans transform.Transformer) (string, error) {
	ret, err := ioutil.ReadAll(transform.NewReader(rawReader, trans))
	if err == nil {
		return string(ret), nil
	} else {
		return "", err
	}
}

func getNowString() (string){
	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)
	const layout2 = "2006-01-02 15:04"
	return nowJST.Format(layout2)
}