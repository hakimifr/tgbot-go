package main

import (
    "log"
    "os"
    "time"

    "github.com/PaulSonOfLars/gotgbot/v2"
    "github.com/PaulSonOfLars/gotgbot/v2/ext"
    "github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func main() {
    token, err := os.ReadFile(".token")
    check(err, "Cannot read token file")

    if string(token) == "" {
        panic("Token is empty")
    }
    var TOKEN string = string(token[0:len(token)-1]) // remove trailing \n

    botopts := gotgbot.BotOpts{}
    bot, err := gotgbot.NewBot(string(TOKEN), &botopts)
    check(err, "Failed to create bot")

    // create updater and dispatcher
    dispatcheropts := ext.DispatcherOpts{
        Error: func(b *gotgbot.Bot, ctx *ext.Context, err error) ext.DispatcherAction {
            // log the error and do nothing
            log.Println("Error while handling update: ", err.Error())
            return ext.DispatcherActionNoop
        },
        MaxRoutines: ext.DefaultMaxRoutines,
    }
    getupdatesopts := gotgbot.GetUpdatesOpts{
        Timeout: 60,
        RequestOpts: &gotgbot.RequestOpts{
            // such high value to make this stfu:
            // 2023/02/18 02:30:10 Failed to get updates; sleeping 1s: failed to execute POST request to getUpdates: Post "https://api.telegram.org/bot[REDACTED]/getUpdates": context deadline exceeded
            Timeout: time.Second * 120,
        },
    }
    pollingopts := ext.PollingOpts{
        DropPendingUpdates: true,
        GetUpdatesOpts: &getupdatesopts,
    }

    dispatcher := ext.NewDispatcher(&dispatcheropts)
    updater := ext.NewUpdater(dispatcher, nil)

    dispatcher.AddHandler(handlers.NewCommand("start", cmd_start))

    err = updater.StartPolling(bot, &pollingopts)
    check(err, "Failed to start polling")

    log.Printf("%s has been started", bot.User.Username)
    updater.Idle()
}
