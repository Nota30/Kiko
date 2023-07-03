package tools

import (
	"github.com/Nota30/Kiko/config"
	"github.com/bwmarrin/discordgo"
)

func NewEmbed(member *discordgo.Member, author string) discordgo.MessageEmbed {
	if len(author) == 0 {
		author = member.User.Username
	}
	
	return discordgo.MessageEmbed{
		Color: config.Color.Default,
		Author: &discordgo.MessageEmbedAuthor{
			Name: author,
			IconURL: member.AvatarURL(""),
		},
	}
}

func NewErrorEmbed(member *discordgo.Member, description string) discordgo.MessageEmbed {	
	return discordgo.MessageEmbed{
		Color: config.Color.Default,
		Author: &discordgo.MessageEmbedAuthor{
			Name: "An error occured",
			IconURL: member.AvatarURL(""),
		},
		Description: description,
	}
}