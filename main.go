package main

import (
    "net/http"
    "log"
    "os"
    "time"

    "github.com/PaulSonOfLars/gotgbot/v2"
    "github.com/PaulSonOfLars/gotgbot/v2/ext"
    "github.com/PaulSonOfLars/gotgbot/v2/ext/handlers"
)

func check(err error, message string) {
    if err != nil {
        panic(message + ": " + err.Error())
    }
}

func main() {
    token, err := os.ReadFile(".token")
    check(err, "Cannot read token file")

    if string(token) == "" {
        panic("Token is empty")
    }
    var TOKEN string = string(token[0:len(token)-1]) // remove trailing \n

    // create new bot
    bot, err := gotgbot.NewBot(string(TOKEN), &gotgbot.BotOpts{
        Client: http.Client{},
        DefaultRequestOpts: &gotgbot.RequestOpts{
            Timeout: gotgbot.DefaultTimeout,
            APIURL: gotgbot.DefaultAPIURL,
        },
    })
    check(err, "Failed to create bot")

    // create updater and dispatcher
    updater := ext.NewUpdater(&ext.UpdaterOpts{
        Dispatcher: ext.NewDispatcher(&ext.DispatcherOpts{
            Error: func(bot *gotgbot.Bot, context *ext.Context, err error) ext.DispatcherAction {
                log.Println("An error occured while handling update: ", err.Error())
                return ext.DispatcherActionNoop
            },
            MaxRoutines: ext.DefaultMaxRoutines,
        }),
    })

    dispatcher := updater.Dispatcher
    dispatcher.AddHandler(handlers.NewCommand("start", start_cmd))

    err = updater.StartPolling(bot, &ext.PollingOpts{
        DropPendingUpdates: true,
        GetUpdatesOpts: gotgbot.GetUpdatesOpts{
            Timeout: 60,
            RequestOpts: &gotgbot.RequestOpts{
                // such high value to make this stfu:
                // 2023/02/18 02:30:10 Failed to get updates; sleeping 1s: failed to execute POST request to getUpdates: Post "https://api.telegram.org/bot[REDACTED]/getUpdates": context deadline exceeded
                Timeout: time.Second * 120,
            },
        },
    })
    check(err, "Failed to start polling")

    log.Printf("%s has been started", bot.User.Username)
    updater.Idle()
}
