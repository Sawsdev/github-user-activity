package event

type Payload struct {
	Action string `json:"action"`
	Issue Issue `json:"issue"`
}