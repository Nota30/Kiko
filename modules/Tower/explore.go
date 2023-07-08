package tower

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/Nota30/Kiko/models"
	"github.com/Nota30/Kiko/tools"
	"github.com/bwmarrin/discordgo"
)

func ExploreCMD(s *discordgo.Session, i *discordgo.InteractionCreate) {
	tools.AwaitInteraction(s, i)
	embed := tools.NewEmbed(i.Member, i.Member.User.Username + "'s Adventure")
	floor, err := models.Floor.FindActiveFloor(i.Member.User.ID)
	if err != nil {
		embed.Description = "You haven't registered yet! You can register using the `/register` command."
	} else {
		actions := make([]int, 0)
		actions = append(actions, 0, 0, 1, 1, 1, 1, 1, 2, 2, 2)
		rand.New(rand.NewSource(time.Now().UnixNano()))
		explore := toFixed(randFloat(5, 15), 1)
		nexplore := floor.Exploration + float32(explore)
		action := actions[rand.Intn(len(actions))]

		if floor.Exploration < 100 {
			if nexplore > 100 {
				nexplore = 100
			}
			floor.Exploration = nexplore
			_, err := models.Floor.SaveFloor(floor)
			if err != nil {
				tools.SendError(s, i, "Was unable to save the data.")
				return
			}
		} else {
			explore = 0
		}

		if action == 0 {
			embed.Description = fmt.Sprintf("# Exploration Complete\nDun dun dun!! You found nothing while exploring. Try exploring more! <:kiko_nervous:1007538837783859230>\n\n> **Floor**: `%d`\n> **Total Exploration**: `%.1f%%` (+%.1f%%)", floor.Floor, floor.Exploration, explore)
		} else if action == 1 {
			embed.Description = fmt.Sprintf("# Exploration Complete\n## Battle Drop\n```diff\n+ Battle card```\n> **Floor**: `%d`\n> **Total Exploration**: `%.1f%%` (+%.1f%%)", floor.Floor, floor.Exploration, explore)
		} else if action == 2 {
			embed.Description = fmt.Sprintf("# Exploration Complete\n## Mineral Deposit\n```diff\n+ Iron ore```\n> **Floor**: `%d`\n> **Total Exploration**: `%.1f%%` (+%.1f%%)", floor.Floor, floor.Exploration, explore)
		}

		if floor.Exploration > 55.5 {
			embed.Footer = &discordgo.MessageEmbedFooter{
				Text: "Note: You have explored more than 55.5% of the current floor! You can ascend floors or keep exploring the current floor.",
			}
		}
	}

	tools.ResponseEdit(s, i, &tools.MessageData{
		Embed: &embed,
	})
}

func randFloat(min, max float64)  float64 {
	return min + rand.Float64()*(max-min)
}

func round(num float64) int {
    return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
    output := math.Pow(10, float64(precision))
    return float64(round(num * output)) / output
}