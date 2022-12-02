package main

import (
	f_api "TelegramOfficeFood/api"
	//store "TelegramOfficeFood/db"
	"net/http"
	"os"
)

func main() {
	//db := store.InitDb()
	//db.Open(os.Getenv("DATABASE_URL"))
	api := f_api.ApiInterface{
		//Store:            db,
		Port:             os.Getenv("PORT"),
		Guid:             os.Getenv("GUID"),
		Token_Telegram:   os.Getenv("TELEGA_TOKEN"),
		URL_dataProvider: os.Getenv("DATA_PROVIDER"),
	}

	http.HandleFunc("/test", api.Test)
	http.HandleFunc("/order", api.Order)

	http.ListenAndServe(":"+api.Port, nil)
}
