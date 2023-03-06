package roleplay

import (
	"strings"
	"time"

	"github.com/Nota30/Kiko/config"
	"github.com/bwmarrin/discordgo"
)

func Actions(s *discordgo.Session,  m *discordgo.MessageCreate, cmd []string) {
	var action string = cmd[0]
	var mention *discordgo.User
	mention = getMention(s, m, cmd)

	cd := inCD(m.Author.ID, m.GuildID, action)

	if cd != 0 {
		embed := &discordgo.MessageEmbed{
			Color: config.Color.Error,
			Description: "Roleplay Cooldown! Please try again in `" + duration(cd) + "`",
		}
		errmsg, _ := s.ChannelMessageSendEmbed(m.ChannelID, embed)
		time.AfterFunc(6*time.Second, func() {
              s.ChannelMessageDelete(errmsg.ChannelID, errmsg.ID)
      	})
		return
	}

	switch action {
		case "bully":
			title := "**{0}** is bullying **{1}**"
			footer := "{2} got bullied {1} times and {0} bullied others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is bullying themselves"
				footer = "{0} got bullied {1} times and bullied others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		case "cuddle":
			title := "**{0}** is cuddling **{1}**"
			footer := "{2} got cuddled {1} times and {0} cuddled others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is cuddling themselves"
				footer = "{0} got cuddled {1} times and cuddled others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		case "hug":
			title := "**{0}** is hugging **{1}**"
			footer := "{2} got hugged {1} times and {0} hugged others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is hugging themselves"
				footer = "{0} got hugged {1} times and hugged others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		case "kiss":
			title := "**{0}** is kissing **{1}**"
			footer := "{2} got kissed {1} times and {0} kissed others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is kissing themselves"
				footer = "{0} got kissed {1} times and kissed others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed) 
		case "lick":
			title := "**{0}** is licking **{1}**"
			footer := "{2} got licked {1} times and {0} licked others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is licking themselves"
				footer = "{0} got licked {1} times and licked others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)  
		case "pat":
			title := "**{0}** is patting **{1}**"
			footer := "{2} got patted {1} times and {0} patted others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is patting themselves"
				footer = "{0} got patted {1} times and patted others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)  
		case "bonk":
			title := "**{0}** is bonking **{1}**"
			footer := "{2} got bonked {1} times and {0} bonked others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is bonking themselves"
				footer = "{0} got bonked {1} times and bonked others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
		case "yeet":
			title := "**{0}** is yeeting **{1}**"
			footer := "{2} got yeeted {1} times and {0} yeeted others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is yeeting themselves"
				footer = "{0} got yeeted {1} times and yeeted others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		case "highfive":
			title := "**{0}** hivefives **{1}**"
			footer := "{2} got hivefived {1} times and {0} hivefived others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** hivefives themselves"
				footer = "{0} got hivefived {1} times and hivefived others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		case "handhold":
			title := "**{0}** holds hands with **{1}**"
			footer := "{2} got handholded {1} times and {0} handholded others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** holds hands with themselves"
				footer = "{0} got handholded {1} times and handholded others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		case "bite":
			title := "**{0}** is biting **{1}**"
			footer := "{2} got bit {1} times and {0} bit others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is biting themselves"
				footer = "{0} got bit {1} times and bit others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		case "slap":
			title := "**{0}** is slapping **{1}**"
			footer := "{2} got slapped {1} times and {0} slapped others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is slapping themselves"
				footer = "{0} got slapped {1} times and slapped others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		case "kill":
			title := "**{0}** is killing **{1}**"
			footer := "{2} got killed {1} times and {0} killed others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is killing themselves"
				footer = "{0} got killed {1} times and killed others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		case "kick":
			title := "**{0}** is kicking **{1}**"
			footer := "{2} got kicked {1} times and {0} kicked others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is kicking themselves"
				footer = "{0} got kicked {1} times and kicked others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)

		case "poke":
			title := "**{0}** is pocking **{1}**"
			footer := "{2} got pocked {1} times and {0} pocked others {3} times"

			if mention.ID == m.Author.ID {
				title = "**{0}** is pocking themselves"
				footer = "{0} got pocked {1} times and pocked others {3} times"
			}

			embed := generateEmbed(action, title, footer, *m.Author, *mention)
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}
}

func getMention(s *discordgo.Session, m *discordgo.MessageCreate, arg []string) *discordgo.User {
	if len(arg) <= 1 {
		return m.Author
	}

	if len(m.Mentions) > 0 {
		return m.Mentions[0]
	}
	
	user, err := s.User(arg[1])
	if err != nil {
		return m.Author
	}

	return user
}

func duration(d time.Duration) string {
    s := d.String()
    if strings.HasSuffix(s, "m0s") {
        s = s[:len(s)-2]
    }
    if strings.HasSuffix(s, "h0m") {
        s = s[:len(s)-2]
    }
    return s
}
