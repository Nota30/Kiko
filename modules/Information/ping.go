package information

import (
	"github.com/Nota30/Kiko/tools"
	"github.com/bwmarrin/discordgo"
)

func PingCMD(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Pong!!",
		},
	})

	if err != nil {
		tools.SendError(s, i, "An error occured while responding.")
	}
}
