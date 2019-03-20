# 利用方法
## /purchase
購入処理を行う
POSTで、購入情報を付与すること

requestのサンプル
```
{
  "user":"zawa",
  "price":1250
}
```

responseのサンプル
```
{
  "IsOk": true,
  "message": ""
}
```

curlで試す場合のサンプル
```
curl -H 'Content-Type: application/json' -H 'User-Agent: Android' -d '{"user":"zawa","price":1250}' http://localhost:8080/purchase | jq .
```

## /monthly_billing
請求処理および、売上通知を行う
GETで良い。
