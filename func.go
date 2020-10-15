package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/yanzay/tbot/v2"
)

var picks = []string{"rock", "paper", "scissors"} // choices from where the bot picks

func init() {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
}

func makeButtons() *tbot.InlineKeyboardMarkup {
	// Create butttons with visible Text and CallbackData as a value.
	btnRock := tbot.InlineKeyboardButton{
		Text:         "tosh",
		CallbackData: "rock",
	}
	btnPaper := tbot.InlineKeyboardButton{
		Text:         "qog'oz",
		CallbackData: "paper",
	}
	btnScissors := tbot.InlineKeyboardButton{
		Text:         "qaychi",
		CallbackData: "scissors",
	}
	return &tbot.InlineKeyboardMarkup{
		InlineKeyboard: [][]tbot.InlineKeyboardButton{
			[]tbot.InlineKeyboardButton{btnRock, btnPaper, btnScissors},
		},
	}
}

func (a *application) draw(humanMove string) (msg string) {
	var result string
	botMove := picks[rand.Intn(len(picks))] // Generaate a random option for the bot

	// Determine the outcome
	switch humanMove {
	case botMove:
		result = "ham yutmadingiz, men ham yutmadim"
		a.draws++
	case options[botMove]:
		result = "yutkazdingiz"
		a.losses++
	default:
		result = "yutdingiz"
		a.wins++
	}
	var humanMoved string
	var botMoved string
	switch humanMove {
	case "rock":
		humanMoved = "tosh"
	case "paper":
		humanMoved = "qo'goz"
	default:
		humanMoved = "qaychi"
	}

	switch botMove {
	case "rock":
		botMoved = "tosh"
	case "paper":
		botMoved = "qo'goz"
	default:
		botMoved = "qaychi"
	}
	msg = fmt.Sprintf("Siz %s ! Siz %s ni tanladingiz va men %s ni tanladim. \n Buyruqlar: \n 1.o'ynash uchun /play ni bosing. \n 2.umumiy natijani bilish uchun /score ni bosing.\n 3. Natijani bekor qilish uchun /reset ni bosing.", result, humanMoved, botMoved)
	bot.HandleMessage("/start", app.startHandler)
	return
}
