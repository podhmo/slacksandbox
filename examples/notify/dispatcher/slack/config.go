package slack

// Config :
type Config struct {
	Token    string        `json:"token"`
	Channels ChannelsConfig `json:"channels"`
}

// ChannelsConfig :
type ChannelsConfig struct {
	Accessed string `json:"accessed"`
}
