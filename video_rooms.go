package twilio

type VideoRoomService struct {
	client *Client
}

const videoRoomPathPart = "Rooms"

type VideoRoom struct {
	Sid        string `json:"sid"`
	EnableTurn bool   `json:"enable_turn"`
	Type       string `json:"type"`
}
