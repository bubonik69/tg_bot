package main

import (
	"bot/fileParsing"
	"bot/mathFunc"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"math/rand"
	"strings"
	"unicode/utf8"
	//"github.com/texttheater/golang-levenshtein/levenshtein"
	"log"
)

func main(){
		bot, err := tgbotapi.NewBotAPI("1881313779:AAGCCDiMrcv48Ood8NJcMYhS7WZ0vsfED3Y")
		if err != nil {
			log.Panic(err)
		}

		bot.Debug = true

		log.Printf("Authorized on account %s", bot.Self.UserName)

		u := tgbotapi.NewUpdate(0)
		u.Timeout = 60

		updates, err := bot.GetUpdatesChan(u)

		for update := range updates {
			if update.Message == nil { // ignore any non-Message Updates
				continue
			}
			// path to dictionary json  "fileParsing/BOT_CONFIG.json"
			go handlerMessage(bot,update,"fileParsing/BOT_CONFIG.json")
		}


	}



func handlerMessage(bot *tgbotapi.BotAPI, update tgbotapi.Update, pathDictJson string){

	myDict := fileParsing.Intents{
		Intents: make(map[string]fileParsing.ExampleRequest),
	}
	// получили словарь вопросов ответов
	myDict =fileParsing.JsonToData(pathDictJson)
	// переводим текст в нижний регистр
	str:=strings.ToLower(update.Message.Text)
	// убираем из текста все знаки припинания
	newStr:=""
	for _,char:=range str {
		if strings.ContainsAny(string(char),"абвгдеёжзийклмнопрстуфхцчшщъыьэюя "){
			newStr+=string(char)
		}
	}
	//сравнить полученный текст с нашей структурой
	var msg string
	for key,val:= range myDict.Intents{
		for _,answer:=range val.Examples{
			//if answer==newStr{
			distance := levenshtein.DistanceForStrings([]rune(newStr),[]rune(answer), levenshtein.DefaultOptions)
			//msg= string(distance)
			newDist:=float64(distance)/float64(mathFunc.MaxLenInt(utf8.RuneCountInString(newStr),utf8.RuneCountInString(answer)))
			if newDist<0.4{
				//	needResponse=key
				msg="нашел совападение" + key + ", дистанция:" + fmt.Sprint(newDist) + ", макс" + fmt.Sprint(mathFunc.MaxLenInt(utf8.RuneCountInString(newStr),utf8.RuneCountInString(answer)))
				msg=myDict.Intents[key].Responses[rand.Intn(len(myDict.Intents[key].Responses)-1)]
			}
		}

	}
	bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, msg))

}





