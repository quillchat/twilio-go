package twilio

import (
	"context"
	"net/url"
	"time"
)

type VideoRoomService struct {
	client *Client
}

const videoRoomPathPart = "Rooms"

type VideoRoom struct {
	Sid                         string    `json:"sid"`
	Status                      string    `json:"status"`
	DateCreated                 time.Time `json:"date_created"`
	DateUpdated                 time.Time `json:"date_updated"`
	AccountSid                  string    `json:"account_sid"`
	EnableTurn                  bool      `json:"enable_turn"`
	Type                        string    `json:"type"`
	UniqueName                  string    `json:"unique_name"`
	StatusCallback              string    `json:"status_callback"`
	StatusCallbackMethod        string    `json:"status_callback_method"`
	EndTime                     time.Time `json:"end_time"`
	MaxParticipants             int       `json:"max_participants"`
	RecordParticipantsOnConnect bool      `json:"record_participants_on_connect"`
	VideoCodecs                 []string  `json:"video_codecs"`
	MediaRegion                 string    `json:"media_region"`
	Url                         string    `json:"url"`
	Links                       []string  `json:"links"`
}

func (vr *VideoRoomService) Get(ctx context.Context, sid string) (*VideoRoom, error) {
	room := new(VideoRoom)
	err := vr.client.GetResource(ctx, videoRoomPathPart, sid, room)
	return room, err
}

func (vr *VideoRoomService) Create(ctx context.Context, data url.Values) (*VideoRoom, error) {
	room := new(VideoRoom)
	err := vr.client.CreateResource(ctx, videoRoomPathPart, data, room)
	return room, err

}
