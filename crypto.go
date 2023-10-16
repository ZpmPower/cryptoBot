package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"reflect"
	"sort"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/lib/pq"
)

type CoinInfo struct {
	ExchangeName string  `json:"exchangeName"`
	Symbol       string  `json:"symbol"`
	FundingRate  float64 `json:"avgFundingRate"`
}

type ResponseData struct {
	Data []CoinInfo `json:"data"`
}

var botToken = os.Getenv("TOKEN")
var cryptoAPIToken = os.Getenv("CRYPTO_API_TOKEN")

func telegramBot() {

	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		panic(err)
	}

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}

		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {

			switch update.Message.Text {
			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm a crypto bot, i can provide you different info about crypto.")
				bot.Send(msg)

			case "/fear_and_greed":
				imageURL := "https://alternative.me/crypto/fear-and-greed-index.png"

				response, err := http.Get(imageURL)
				if err != nil {
					log.Panic(err)
				}
				defer response.Body.Close()
				imageData, err := io.ReadAll(response.Body)
				if err != nil {
					log.Panic(err)
				}
				imageReader := bytes.NewReader(imageData)
				photo := tgbotapi.NewPhoto(update.Message.Chat.ID, tgbotapi.FileReader{Reader: imageReader})
				photo.Caption = "Get more about \"fear and greed index\" - <b><a href=\"https://www.investopedia.com/terms/f/fear-and-greed-index.asp\">HERE</a></b>"
				photo.ParseMode = "HTML"
				bot.Send(photo)
			case "/funding_rate":
				url := "https://open-api.coinglass.com/public/v2/futures_coins_markets"

				req, err := http.NewRequest("GET", url, nil)
				if err != nil {
					fmt.Println("Error creating request:", err)
					return
				}

				req.Header.Add("coinglassSecret", cryptoAPIToken)
				req.Header.Add("accept", "application/json")

				client := http.Client{}
				res, err := client.Do(req)
				if err != nil {
					fmt.Println("Error making request:", err)
					return
				}
				defer res.Body.Close()

				if res.StatusCode != http.StatusOK {
					fmt.Println("Request failed with status code:", res.Status)
					return
				}

				var responseData ResponseData
				decoder := json.NewDecoder(res.Body)
				if err := decoder.Decode(&responseData); err != nil {
					fmt.Println("Error decoding response:", err)
					return
				}

				sort.SliceStable(responseData.Data, func(i, j int) bool {
					return responseData.Data[i].FundingRate > responseData.Data[j].FundingRate
				})

				responseString := "Biggest Funding rate:\n"
				for _, coin := range responseData.Data[:10] {
					responseString += fmt.Sprintf("%.6f: #%s\n", coin.FundingRate, coin.Symbol)
				}

				sort.SliceStable(responseData.Data, func(i, j int) bool {
					return responseData.Data[i].FundingRate < responseData.Data[j].FundingRate
				})

				responseString += fmt.Sprintf("\nLowest Funding rate:\n")
				for _, coin := range responseData.Data[:10] {
					responseString += fmt.Sprintf("%.6f: #%s\n", coin.FundingRate, coin.Symbol)
				}

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseString)
				bot.Send(msg)
			case "/crypto_news":
				numberOfNews := 10
				news, err := GetNews()
				fmt.Println(err)
				if err != nil {

					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "News error.")
					bot.Send(msg)
				}
				news = news[:numberOfNews]

				var responseString string
				for index, new := range news {
					responseString += fmt.Sprintf("%d. <a href=\"https://ru.investing.com/%s\">%s</a>\n", index+1, new.link, new.title)
					responseString += fmt.Sprintf("%s %s\n", new.source, new.timeAgo)
					responseString += fmt.Sprintf("%s\n\n", new.content)
				}
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, responseString)
				msg.ParseMode = "HTML"
				bot.Send(msg)
			}
		} else {
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Use the words for search.")
			bot.Send(msg)
		}
	}
}

func main() {

	time.Sleep(1 * time.Minute)
	telegramBot()
}
