package modules

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/Nota30/Kiko/tools"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/bwmarrin/discordgo"
)

func GPT(s *discordgo.Session, m *discordgo.MessageCreate) {
	client := gpt3.NewClient(tools.GetEnv("OpenAIApiKey"))

	goCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	response, err := getResponse(goCtx, client, m.Content)
	if err != nil {
		fmt.Println(err)
	}

	embed := &discordgo.MessageEmbed{}
	embed.Title = "Chat GPT"
	embed.Color = 42622
	embed.Description = response

	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{}
	embed.Thumbnail.URL = "https://preview.redd.it/what-does-the-chatgpt-symbol-mean-does-it-have-meaning-like-v0-ogw5okm1bz5a1.jpg?auto=webp&s=51f427fda04a94889e4abfcbc6e6848b77be81ba"

	_, err = s.ChannelMessageSendEmbed(m.ChannelID, embed)

	if err != nil {
		fmt.Println(err)
	}

}

func getResponse(ctx context.Context, client gpt3.Client, question string) (response string, err error) {
	sb := strings.Builder{}

	err = client.CompletionStreamWithEngine(
		ctx,
		gpt3.TextDavinci003Engine,
		gpt3.CompletionRequest{
			Prompt: []string{
				question,
			},
			MaxTokens:   gpt3.IntPtr(3000),
			Temperature: gpt3.Float32Ptr(0),
		},
		func(resp *gpt3.CompletionResponse) {
			text := resp.Choices[0].Text

			sb.WriteString(text)
		},
	)
	if err != nil {
		return "", err
	}

	response = sb.String()
	response = strings.TrimLeft(response, "\n")

	return response, nil
}
