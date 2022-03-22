package api

type Category struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	TotalChannels  int    `json:"total_channels"`
	OnlineChannels int    `json:"online_channels"`
	Viewers        int    `json:"viewers"`
	Adult          bool   `json:"adult"`
}

type Channel struct {
	UserId       int64  `json:"user_id"`
	Name         string `json:"name"`
	Avatar       string `json:"avatar"`
	Online       bool   `json:"online"`
	Viewers      int64  `json:"viewers"`
	ViewersTotal int64  `json:"viewers_total"`
	Thumbnails   struct {
		Web      string `json:"web"`
		WebLarge string `json:"web_large"`
		Mobile   string `json:"mobile"`
		Tablet   string `json:"tablet"`
	} `json:"thumbnails"`
	Followers         int64    `json:"followers"`
	Subscribers       int64    `json:"subscribers"`
	Adult             bool     `json:"adult"`
	Category          []string `json:"category"`
	AccountType       string   `json:"account_type"`
	Commissions       bool     `json:"commissions"`
	Recordings        bool     `json:"recordings"`
	Title             string   `json:"title"`
	DescriptionPanels []struct {
		Title      string `json:"title"`
		Body       string `json:"body"`
		Image      string `json:"image"`
		ImageLink  string `json:"image_link"`
		ButtonText string `json:"button_text"`
		ButtonLink string `json:"button_link"`
		Position   int    `json:"position"`
	} `json:"description_panels"`
	Private        bool   `json:"private"`
	PrivateMessage string `json:"private_message"`
	Gaming         bool   `json:"gaming"`
	ChatSettings   struct {
		GuestChat bool        `json:"guest_chat"`
		Links     bool        `json:"links"`
		Level     interface{} `json:"level"`
	} `json:"chat_settings"`
	LastLive    *string  `json:"last_live"`
	Tags        []string `json:"tags"`
	Multistream []struct {
		UserID int    `json:"user_id"`
		Name   string `json:"name"`
		Online bool   `json:"online"`
		Adult  bool   `json:"adult"`
	} `json:"multistream"`
	Languages []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"languages"`
	Following    bool   `json:"following"`
	CreationDate string `json:"creation_date"`
}

type Video struct {
	Title     string `json:"title"`
	File      string `json:"file"`
	Filesize  int64  `json:"filesize"`
	Duration  int64  `json:"duration"`
	Views     int64  `json:"views"`
	Timestamp string `json:"timestamp"`
	Adult     bool   `json:"adult"`
}

type Online struct {
	UserId     int    `json:"user_id"`
	Name       string `json:"name"`
	Avatar     string `json:"avatar"`
	Title      string `json:"title"`
	Viewers    int    `json:"viewers"`
	Thumbnails struct {
		Web      string `json:"web"`
		WebLarge string `json:"web_large"`
		Mobile   string `json:"mobile"`
		Tablet   string `json:"tablet"`
	} `json:"thumbnails"`
	Category    string `json:"category"`
	AccountType string `json:"account_type"`
	Adult       bool   `json:"adult"`
	Gaming      bool   `json:"gaming"`
	Commissions bool   `json:"commissions"`
	Multistream bool   `json:"multistream"`
	Languages   []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"languages"`
	Following bool `json:"following"`
}

type Stream struct {
	Channel struct {
		Avatar       string  `json:"avatar"`
		StreamName   string  `json:"stream_name"`
		Name         string  `json:"name"`
		Online       bool    `json:"online"`
		UserId       int     `json:"user_id"`
		Adult        bool    `json:"adult"`
		OfflineImage *string `json:"offline_image"`
	} `json:"channel"`
	ShowAds int     `json:"show_ads"`
	Url     *string `json:"url"`
}

type Notification struct {
	Body      string `json:"body"`
	Uri       string `json:"uri"`
	Timestamp bool   `json:"timestamp"`
}
