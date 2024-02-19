package main

import (
	"fmt"
	"time"

	"github.com/PaulSonOfLars/gotgbot/v2"
    "github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func cmd_fastpurge(bot *gotgbot.Bot, ctx *ext.Context) error {
	if ctx.EffectiveMessage.ReplyToMessage == nil {
		bot.SendMessage(ctx.EffectiveChat.Id, "Reply to a message to purge it.", nil)
		return nil
	}
	admins, err := ctx.EffectiveChat.GetAdministrators(bot, nil)
	check(err, "Failed to retrieve chat admins!")

	// make sure user is an admin
	isAdmin := false
	for _, admin := range admins {
		if admin.GetUser().Id == ctx.EffectiveUser.Id {
			isAdmin = true
			break
		}
	}

	if !isAdmin {
		ctx.Message.Reply(bot, "You must be an admin to purge messages!", nil)
		return nil
	}

	msg, _ := ctx.Message.Reply(bot, "Purging messages...", nil)

	startTime := time.Now()
	message_list := pyRange(ctx.EffectiveMessage.ReplyToMessage.MessageId,
							ctx.EffectiveMessage.MessageId + 1)
	for len(message_list) > 0 {
		if len(message_list) > 100 {
			bot.DeleteMessages(ctx.EffectiveChat.Id, message_list[:100], nil)
			message_list = message_list[100:]
		} else {
			bot.DeleteMessages(ctx.EffectiveChat.Id, message_list, nil)
			message_list = []int64{}
		}
	}
	endTime := time.Now()
	timeTaken := endTime.Sub(startTime).Seconds()
	_, _, err = msg.EditText(bot, fmt.Sprintf("Purged %d messages in %.3f seconds.",
											  ctx.EffectiveMessage.MessageId - ctx.EffectiveMessage.ReplyToMessage.MessageId,
											  timeTaken),
							 nil)
	return err
}
