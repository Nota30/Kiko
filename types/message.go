package types

type MessageEmbedAuthor struct {
	Text string
	Icon string
}

type MessageEmbedFooter struct {
	Text string
	Icon string
}

type MessageEmbed struct {
	Author MessageEmbedAuthor
	Title string
	Color int
	Description string
	Timestamp bool
	Footer MessageEmbedFooter
}
