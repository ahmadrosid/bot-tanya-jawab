package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/ahmadrosid/bot-tanya-jawab/config"
	"github.com/ahmadrosid/bot-tanya-jawab/handler"
	"github.com/ahmadrosid/bot-tanya-jawab/service"
	"github.com/bwmarrin/discordgo"
)

func initBot(cfg config.Config, callback func(ds *discordgo.Session)) {
	dc, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatalf("Failed to init bot: %v", err)
	}

	dc.Identify.Intents = discordgo.IntentGuilds | discordgo.IntentGuildMessages //| discordgo.IntentGuildMembers
	callback(dc)

	err = dc.Open()
	if err != nil {
		log.Fatalf("Failed to connect websocket to discord server: %v", err)
	}

	defer dc.Close()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("closing connection to discord server")
}

func main() {
	cfg := config.Get()
	println("token" + cfg.Token)
	initBot(cfg, func(ds *discordgo.Session) {
		log.Println("Success to connect with discord server")

		service := service.NewBotService(cfg.Channel)
		hdl := handler.NewBotHandler(service)
		ds.AddHandler(hdl.OnReady)
		ds.AddHandler(hdl.OnInteraction)
		ds.AddHandler(hdl.OnMessage)
	})
}
