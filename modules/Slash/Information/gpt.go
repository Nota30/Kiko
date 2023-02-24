package modules

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/Nota30/Kiko/tools"

	"github.com/PullRequestInc/go-gpt3"
	"github.com/bwmarrin/discordgo"
)

type record struct {
	Word     string
	Language string
}

type profanityStruct struct {
	Records []record
}

func GPT(s *discordgo.Session, m *discordgo.MessageCreate) {
	resp, err := http.Get("https://raw.githubusercontent.com/turalus/encycloDB/master/Dirty%20Words/DirtyWords.json")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.

	defer resp.Body.Close()

	//Create a variable of the same type as our model
	pResp := &profanityStruct{}
	//Decode the data
	if err = json.NewDecoder(resp.Body).Decode(&pResp); err != nil {
		log.Fatal(err)
	}

	//Convert the body to type string
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

	embed.Thumbnail = &discordgo.MessageEmbedThumbnail{}
	embed.Thumbnail.URL = "https://preview.redd.it/what-does-the-chatgpt-symbol-mean-does-it-have-meaning-like-v0-ogw5okm1bz5a1.jpg?auto=webp&s=51f427fda04a94889e4abfcbc6e6848b77be81ba"

	if success, errMsg := moderationTest(response, pResp); success {
		embed.Description = response
	} else {
		embed.Description = errMsg
	}

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

func moderationTest(str string, pResp *profanityStruct) (bool, string) {
	for _, word := range strings.Split(str, " ") {
		if word[0] == '@' {
			return false, "Whoa there buddy! Let's not do a ping okay?"
		}
		for _, record := range pResp.Records {
			for _, part := range strings.Split(record.Word, " ") {
				if strings.ToLower(word) == part {
					return false, "Whoa there buddy! Let's not say bad words"
				}
			}
		}
	}

	return true, ""

}
