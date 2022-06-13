package service

import (
	"fmt"
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Service struct {
	ChannelID string
}

func NewBotService(channelID string) Service {
	return Service{ChannelID: channelID}
}

func (s *Service) SendMessageToChannel(session *discordgo.Session, content string) {
	data := discordgo.MessageSend{
		Content: content,
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Style:    discordgo.PrimaryButton,
						Label:    "Jawab",
						CustomID: "-",
					},
				},
			},
		},
	}
	_, err := session.ChannelMessageSendComplex(s.ChannelID, &data)
	if err != nil {
		log.Printf("error send message to channel: %v", err)
	}
}

func (s *Service) RespondInteraction(session *discordgo.Session, in *discordgo.InteractionCreate) {
	data := discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Terimakasih telah mengambil pertanyaan ini.",
		},
	}
	err := session.InteractionRespond(in.Interaction, &data)
	if err != nil {
		log.Printf("Failed to send interaction response")
	}

	ch, err := session.MessageThreadStart(s.ChannelID, in.Message.ID, "new-thread", 0)
	if err != nil {
		log.Printf("failed to start thread: %v", err)
	}

	log.Printf("thread created with id: %v", ch.ID)

	_, err = session.ChannelMessageSend(ch.ID, fmt.Sprintf("Hi <@%s> lanjutkan jawaban disini", in.Member.User.ID))
	if err != nil {
		log.Printf("Failed to metion user in thread: %v", err)
	}

	time.Sleep(time.Second)
	err = session.InteractionResponseDelete(in.Interaction)
	if err != nil {
		log.Printf("Failed to delete interaction message: %v", err)
	}
}

func (s *Service) RespondMessage(session *discordgo.Session, msg *discordgo.MessageCreate) {
	if msg.Content == "selesai" && msg.ChannelID != s.ChannelID {
		_, err := session.ChannelDelete(msg.ChannelID)
		if err != nil {
			log.Printf("Failed to delete channel: %v", err)
		}
	}
}
