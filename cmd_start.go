package main

import (
    "fmt"

    "github.com/PaulSonOfLars/gotgbot/v2"
    "github.com/PaulSonOfLars/gotgbot/v2/ext"
)

func start_cmd(bot *gotgbot.Bot, context *ext.Context) error {
    _, err := context.EffectiveMessage.Reply(bot, "Hi!", &gotgbot.SendMessageOpts{})
    if err != nil {
        return fmt.Errorf("[startcmd] Failed to reply message")
    }
    return nil
}
