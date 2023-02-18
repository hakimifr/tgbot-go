package main

import (
    "fmt"

    "github.com/PaulSonOfLars/gotgbot/v2"
    "github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func cmd_start(bot *gotgbot.Bot, context *ext.Context) error {
    _, err := context.EffectiveMessage.Reply(bot, "Hi!", &gotgbot.SendMessageOpts{})
    if err != nil {
        return fmt.Errorf("[%s] Failed to reply message: %w", getFunctionName(), err)
    }
    return nil
}
