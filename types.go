package main

import "time"

type Snowflake int64

type User struct {
	Id Snowflake            `json:"id,string"`
	Username string         `json:"username,omitempty"`
	Discriminator string 	`json:"discriminator,omitempty"`
	Avatar string 			`json:"avatar,omitempty"`
	// Flags can be used to Identify someone's Badges through Bit-Shifting
	PublicFlags int 		`json:"public_flags,omitempty"`
}

type Webhook struct {
	Id Snowflake 			`json:"id,string"`
	Type int 				`json:"type,omitempty"`
	GuildId Snowflake 		`json:"guild_id,string,omitempty"`
	ChannelId Snowflake	 	`json:"channel_id,string,omitempty"`
	User User 				`json:"user,omitempty"`
	Name string 			`json:"name,omitempty"`
	Avatar string 			`json:"avatar,omitempty"`
	Token string 			`json:"token,omitempty"`
	Url string 				`json:"url,omitempty"`
}

type WebhookMessage struct {
	Content   string `json:"content,omitempty"`
	Username  string `json:"username,omitempty"`
	AvatarURL string `json:"avatar_url,omitempty"`
	Embeds    *[]Embed `json:"embeds,omitempty"`
}


type Embed struct {
	Title       string    `json:"title,omitempty"`
	Description string    `json:"description,omitempty"`
	URL         string    `json:"url,omitempty"`
	Color       int       `json:"color,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Footer      *Footer    `json:"footer,omitempty"`
	Thumbnail   *Thumbnail `json:"thumbnail,omitempty"`
	Image       *Image     `json:"image,omitempty"`
	Author      *Author    `json:"author,omitempty"`
	Fields      *[]Field  `json:"fields,omitempty"`
}

type Footer struct {
	IconUrl string `json:"icon_url,omitempty"`
	Text    string `json:"text"`
}

type Thumbnail struct {
	URL string `json:"url"`
}

type Image struct {
	URL string `json:"url"`
}

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline"`
}

type LogType struct {
	LogTitle string
	LogDescription string
	LogFields  *[]Field
}
