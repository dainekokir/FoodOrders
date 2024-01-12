package main

import (
	f_api "TelegramOfficeFood/api"
	"net/http"
	"os"
)

func main() {
	api := f_api.ApiInterface{
		Port:             os.Getenv("PORT"),
		Guid:             os.Getenv("GUID"),
		Token_Telegram:   os.Getenv("TELEGA_TOKEN"),
		URL_dataProvider: os.Getenv("DATA_PROVIDER"),
	}

	http.HandleFunc("/test", api.Test)
	http.HandleFunc("/order", api.Order)

	http.ListenAndServe(":"+api.Port, nil)
}
