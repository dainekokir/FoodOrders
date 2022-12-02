package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//	type dataStore interface {
//		AddRequest(string)
//	}
type ApiInterface struct {
	//Store            dataStore
	Port             string
	Guid             string
	URL_telegram     string
	URL_dataProvider string
}

func (ai *ApiInterface) Test(w http.ResponseWriter, r *http.Request) {
	//ai.Store.AddRequest("This_is_shit_was_start")

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("This_is_shit_was_start"))
	return
}

func (ai *ApiInterface) Order(res http.ResponseWriter, req *http.Request) {

	body := &webhookReqBody{}
	if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}

	//ai.Store.AddRequest(body.Message.Text)
	fmt.Println("Handler get", body.Message.Text)

	if body.Message.Text == "/start" {
		sendMessageToTelegram(
			body.Message.Chat.ID,
			"Для получения заказа на текущий день, укажите ID пользователя \n\nСохранить\\перезаписать свой номер: My-ID пользователя \n(пример: My-123)",
			ai.URL_telegram,
		)
	}

	matched, err := regexp.MatchString(`^my-\d+$`, strings.ToLower(body.Message.Text))
	if matched {
		UserPortalID := 0
		re, err := regexp.Compile(`^my-(?P<number>\d+)$`)
		if err != nil {
			return
		}
		res := re.FindAllStringSubmatch(strings.ToLower(body.Message.Text), -1)
		for _, v := range res {
			for kk, vv := range re.SubexpNames() {
				if vv == "number" {
					UserPortalID, _ = strconv.Atoi(v[kk])
					break
				}
			}
		}

		if UserPortalID > 0 {
			sendMsgBtn(
				body.Message.Chat.ID,
				UserPortalID,
				ai.URL_telegram,
			)
		}
	}

	UserPersonPortalID, err := strconv.Atoi(body.Message.Text)
	if err != nil {
		return
	}

	if err := getFoodDish(body.Message.Chat.ID, UserPersonPortalID, ai.URL_telegram, ai.Guid, ai.URL_dataProvider); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}

	fmt.Println("reply sent")
}

func sendMsgBtn(chatID int64, userID int, URL_telegram string) error {
	reqBody := &sendMessageBtnReqBody{
		ChatID: chatID,
		Text:   "Сохранен номер: " + fmt.Sprint(userID),
		ReplyMarkup: ReplyKeyboardMarkup{
			Keyboard:        [][]KeyboardButton{{KeyboardButton{Text: fmt.Sprint(userID)}}},
			ResizeKeyboard:  true,
			OneTimeKeyboard: false,
			Selective:       false,
		},
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	// Send a post request with your token
	_, err = http.Post(URL_telegram, "application/json", bytes.NewBuffer(reqBytes))

	if err != nil {
		return err
	}

	return nil
}

// sendMessageToTelegram Оправка сообщения на сервер телеграмма
func sendMessageToTelegram(chatID int64, message string, URL_telegram string) error {
	// Create the request body struct
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   message,
	}
	// Create the JSON body from the struct
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
	// Send a post request with your token
	res, err := http.Post(URL_telegram, "application/json", bytes.NewBuffer(reqBytes))

	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}

// getFoodDish Получение заказа от сервиса еды
func getFoodDish(chatID int64, UserID int, URL_telegram string, Guid string, URL_dataProvider string) error {
	tm := time.Now()
	tmStr := tm.Format("01-02-2006")

	query := URL_dataProvider + "?UserId=" + strconv.Itoa(UserID) + "&Date=" + tmStr + "&Guid=" + Guid
	resp, err := http.Get(query)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	response := "Заказ на " + time.Now().Format("02-01-2006") + "\r\n" + string(body)
	err = sendMessageToTelegram(chatID, response, URL_telegram)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	return nil
}

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}
