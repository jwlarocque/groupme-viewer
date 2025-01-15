package main

type Chat struct {
	Metadata Conversation
	Messages []Message
}

type Conversation struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	ImageUrl string `json:"image_url"`
	CreatorUserId string `json:"creator_user_id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	ShareUrl string `json:"share_url"`
	Members []struct {
		UserId string `json:"user_id"`
		Nickname string `json:"nickname"`
		ImageUrl string `json:"image_url"`
		Id string `json:"id"`
		Roles []string `json:"roles"`
		Name string `json:"name"`
	} `json:"members"`
}

type Message struct {
	GroupId string `json:"group_id"`
	Id string `json:"id"`
	Name string `json:"name"`
	SenderId string `json:"sender_id"`
	CreatedAt int64 `json:"created_at"`
	Text string `json:"text"`
	Attachments []struct {
		Type string `json:"type"`
		Url string `json:"url"`
		UserId string `json:"user_id"`
		ReplyId string `json:"reply_id"`
	} `json:"attachments"`
	Reactions []struct {
		Type string `json:"type"`
		UserIds []string `json:"user_ids"`
		Code string `json:"code"`
	} `json:"reactions"`
}