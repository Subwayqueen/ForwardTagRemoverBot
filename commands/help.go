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
	"github.com/PaulSonOfLars/gotgbot"
	"github.com/PaulSonOfLars/gotgbot/ext"
	"github.com/PaulSonOfLars/gotgbot/parsemode"
	"go.uber.org/zap"
)

func Help(b ext.Bot, u *gotgbot.Update) error {

	helpButton := [][]ext.InlineKeyboardButton{make([]ext.InlineKeyboardButton, 2), make([]ext.InlineKeyboardButton, 1)}

	helpButton[0][0] = ext.InlineKeyboardButton{
		Text: "📩 𝗦𝗢𝗨𝗥𝗖𝗘 𝗖𝗢𝗗𝗘",
		Url:  "https://t.me/NoSourceCode",
	}

	helpButton[0][1] = ext.InlineKeyboardButton{
		Text: "🧕 𝗠𝗬 𝗖𝗥𝗘𝗔𝗧𝗢𝗥",
		Url:  "https://telegram.dog/IGGIE",
	}

	helpButton[1][0] = ext.InlineKeyboardButton{
		Text: "🥳 𝗦𝗨𝗣𝗣𝗢𝗥𝗧 𝗕𝗬 𝗝𝗢𝗜𝗡𝗜𝗡𝗚 𝗢𝗨𝗥 𝗖𝗛𝗔𝗡𝗡𝗘𝗟 🥳",
		Url:  "https://t.me/joinchat/AAAAAE-44AkxSyqIMj1tdQ",
	}

	markup := ext.InlineKeyboardMarkup{InlineKeyboard: &helpButton}

	msg := b.NewSendableMessage(u.EffectiveChat.Id, "*Forward Me A File, Video, Audio, Photo or Anything & \nI will Send You the File Back*\n\n*𝗛𝗼𝘄 𝗧𝗼 𝗦𝗲𝘁 𝗖𝗮𝗽𝘁𝗶𝗼𝗻?*\n*Reply 𝗖𝗮𝗽𝘁𝗶𝗼𝗻 to a File, Photo, Audio Or Media*")
	msg.ReplyToMessageId = u.EffectiveMessage.MessageId
	msg.ReplyMarkup = &markup
	msg.ParseMode = parsemode.Markdown
	_, err := msg.Send()
	if err != nil {
		b.Logger.Warnw("Error in sending", zap.Error(err))
	}
	return err
}
