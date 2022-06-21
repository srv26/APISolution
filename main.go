package main

import (
	"net/http"

	"main.go/src/apis/product_api"
)

func main() {
	http.HandleFunc("/cars/describe", product_api.DecribeTable)
	http.HandleFunc("/cars/name/filer?color=green", product_api.QuerySelectNameFromGivenColor)
	http.HandleFunc("/cars/filter?type=minivan&color=green", product_api.QuerySelectFromGivenColorAndType)
	http.HandleFunc("/innerjoin", product_api.FindInnerJoin)
	http.HandleFunc("/leftjoin", product_api.FindLeftJoin)
	http.HandleFunc("/rightjoin", product_api.FindRightJoin)
	http.HandleFunc("/fulljoin", product_api.FindFullJoin)
	http.ListenAndServe(":8080", nil)
}
