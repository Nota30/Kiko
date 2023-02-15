package configs

type Config struct {
	Bot struct {
		Token string `envconfig:"DISCORD_TOKEN"`
	}
	Color struct {
		Default string `yaml:"default"`
		Error   string `yaml:"error"`
	} `yaml:"colors"`
}