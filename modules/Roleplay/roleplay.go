package roleplay

import (
	"github.com/bwmarrin/discordgo"
)

func Actions (s *discordgo.Session,  m *discordgo.MessageCreate, action string) {
	switch action {
		// case "bully": 
		// case "cuddle": 
		case "hug":
			title := "**{0}** is hugging **{1}**"
			footer := "{2} got hugged {1} times and {0} hugged others {3} times"
			var mention *discordgo.User

			if len(m.Mentions) <= 0 {
				mention = m.Author
				title = "**{0}** is hugging themselves"
				footer = "{0} got hugged {1} times and hugged others {3} times"
			} else {
				mention = m.Mentions[0]
			}

			if mention == m.Author {
				title = "**{0}** is hugging themselves"
				footer = "{0} got hugged {1} times and hugged others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
		// case "kiss": 
		// case "lick": 
		// case "pat": 
		// case "bonk":
		// case "yeet": 
		// case "highfive": 
		// case "handhold": 
		// case "bite": 
		// case "slap": 
		// case "kill": 
		// case "kick": 
		// case "poke":
	}
}
