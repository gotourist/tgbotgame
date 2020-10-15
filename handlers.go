package main

import (
	"fmt"

	"github.com/yanzay/tbot/v2"
)

// Handle the /start command here
func (a *application) startHandler(m *tbot.Message) {
	msg := "Bu bot siz bilan tosh, qaychi, qog'oz o'yinini o'ynaydi.\n Buyruqlar: \n 1.o'ynash uchun /play ni bosing. \n 2.umumiy natijani bilish uchun /score ni bosing.\n 3. Natijani bekor qilish uchun /reset ni bosing."
	a.client.SendMessage(m.Chat.ID, msg)
}

// Handle the /play command here
func (a *application) playHandler(m *tbot.Message) {
	buttons := makeButtons()
	a.client.SendMessage(m.Chat.ID, "Siz tanlang:", tbot.OptInlineKeyboardMarkup(buttons))
}

// Handle the /score command here
func (a *application) scoreHandler(m *tbot.Message) {
	msg := fmt.Sprintf("Natijangiz:\n Yutishlar: %v\n Tengliklar: %v\n Yutkazishlar: %v \n Buyruqlar: \n 1.o'ynash uchun /play ni bosing. \n 2.umumiy natijani bilish uchun /score ni bosing.\n 3. Natijani bekor qilish uchun /reset ni bosing.", a.wins, a.draws, a.losses)
	a.client.SendMessage(m.Chat.ID, msg)
}

// Handle the /reset command here
func (a *application) resetHandler(m *tbot.Message) {
	a.wins, a.draws, a.losses = 0, 0, 0
	msg := "Umumiy natija 0 ga tenglandi. \n Buyruqlar: \n 1.o'ynash uchun /play ni bosing. \n 2.umumiy natijani bilish uchun /score ni bosing.\n 3. Natijani bekor qilish uchun /reset ni bosing."
	a.client.SendMessage(m.Chat.ID, msg)
}

// Handle buttton presses here
func (a *application) callbackHandler(cq *tbot.CallbackQuery) {
	humanMove := cq.Data
	msg := a.draw(humanMove)
	a.client.DeleteMessage(cq.Message.Chat.ID, cq.Message.MessageID)
	a.client.SendMessage(cq.Message.Chat.ID, msg)
}
