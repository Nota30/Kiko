package events

// -> Maybe not handle all events in one handler?

// import (
// 	"github.com/bwmarrin/discordgo"
// )

// func EventHandler(s *discordgo.Session, evt interface{}) {
// 	switch evt.(type) {
// 	case *discordgo.MessageCreate:
// 		event, ok := evt.(*discordgo.MessageCreate)
// 		if ok {
// 			MSGCreateHandler(s, event)
// 		}

// 	case *discordgo.InteractionCreate:
// 		event, ok := evt.(*discordgo.InteractionCreate)
// 		if ok {
// 			InteractionHandler(s, event)
// 		}
// 	}
// }
