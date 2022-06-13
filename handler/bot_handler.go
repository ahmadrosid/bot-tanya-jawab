package handler

import (
	"log"

	"github.com/ahmadrosid/bot-tanya-jawab/service"
	"github.com/bwmarrin/discordgo"
)

type handler struct {
	service service.Service
}

func NewBotHandler(svc service.Service) handler {
	return handler{service: svc}
}

func (h *handler) OnReady(session *discordgo.Session, message *discordgo.Ready) {
	log.Println("ready to listen for discord events")
	h.service.SendMessageToChannel(session, "Hello there")
}

func (h *handler) OnInteraction(session *discordgo.Session, in *discordgo.InteractionCreate) {
	h.service.RespondInteraction(session, in)
}

func (h *handler) OnMessage(session *discordgo.Session, msg *discordgo.MessageCreate) {
	h.service.RespondMessage(session, msg)
}
