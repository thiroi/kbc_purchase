package main

type PurchaseResponse struct{
	IsOk bool `json:"IsOk"`
	Message string `json:"message"`
}


func makeError(err error)(PurchaseResponse){
	return PurchaseResponse{false, err.Error()}
}

func makeSuccess()(PurchaseResponse){
	return PurchaseResponse{true, ""}
}