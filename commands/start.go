/*  Copyright (C) 2020 by Anandpskerala@Github, < https://github.com/Anandpskerala >.
 *
 * This file is part of < https://github.com/Anandpskerala/ForwardTagRemoverBot > project,
 * and is released under the "GNU v3.0 License Agreement".
 * Please see < https://github.com/Anandpskerala/ForwardTagRemoverBot/blob/master/LICENSE >
 *
 * All rights reserved.
 */

package commands
import (
	"fmt"

	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/parsemode"
	"go.uber.org/zap"
)

func Start(b ext.Bot, u *gotgbot.Update) error {

	startButton := [][]ext.InlineKeyboardButton{make([]ext.InlineKeyboardButton, 2), make([]ext.InlineKeyboardButton, 1)}

	startButton[0][0] = ext.InlineKeyboardButton{
		Text: "📩 𝗦𝗢𝗨𝗥𝗖𝗘 𝗖𝗢𝗗𝗘",
		Url:  "https://t.me/NoSourceCode",
	}

	startButton[0][1] = ext.InlineKeyboardButton{
		Text: "🧕 𝗠𝗬 𝗖𝗥𝗘𝗔𝗧𝗢𝗥",
		Url:  "https://telegram.dog/Iggie",
	}

	startButton[1][0] = ext.InlineKeyboardButton{
		Text: "🥳 𝗦𝗨𝗣𝗣𝗢𝗥𝗧 𝗕𝗬 𝗝𝗢𝗜𝗡𝗜𝗡𝗚 𝗢𝗨𝗥 𝗖𝗛𝗔𝗡𝗡𝗘𝗟 🥳",
		Url:  "https://t.me/joinchat/AAAAAE-44AkxSyqIMj1tdQ",
	}

	markup := ext.InlineKeyboardMarkup{InlineKeyboard: &startButton}

	msg := b.NewSendableMessage(u.EffectiveChat.Id, fmt.Sprintf("*Hello* [%s](tg://user?id=%v) 😎, *I'm A 𝗙𝗢𝗥𝗪𝗔𝗥𝗗 𝗧𝗔𝗚 𝗥𝗘𝗠𝗢𝗩𝗘𝗥 𝗕𝗢𝗧.*\n\n*Send /help To Know What I Can Do*", u.EffectiveUser.FirstName, u.EffectiveUser.Id))
	msg.ReplyToMessageId = u.EffectiveMessage.MessageId
	msg.ReplyMarkup = &markup
	msg.ParseMode = parsemode.Markdown
	_, err := msg.Send()
	if err != nil {
		b.Logger.Warnw("Error in sending", zap.Error(err))
	}
	return err
}
