package models

type FileMeta struct {
	Id             string       `json:"id"`
	Title          string       `json:"title"`
	Formats        []FileFormat `json:"formats"`
	Thumbnail      string       `json:"thumbnail"`
	DurationString string       `json:"duration_string"`
	Channel        string       `json:"channel"`
	LikeCount      int          `json:"like_count"`
	ViewCount      int          `json:"view_count"`
}
