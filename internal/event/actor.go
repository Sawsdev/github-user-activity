package event

type Actor struct {
	Id int `json:"id"`
	Login string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarId string `json:"gravatar_id"`
	Url string `json:"url"`
	AvatarUrl string `json:"avatar_url"`
}