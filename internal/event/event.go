package event

type Event struct {
	Id string `json:"id"`
	Type string `json:"type"`
	Actor Actor `json:"actor"`
	Repo Repo  `json:"repo"`
	Payload Payload `json:"payload"`
	Public bool `json:"public"`
	CreatedAt string `json:"created_at"`
}
