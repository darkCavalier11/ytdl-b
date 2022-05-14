package models

import "github.com/darkCavalier11/downloader_backend/grpc_module/gen"

// FileMeta struct is necessary for parsing the dumped json in the stdout and convert
// it to generated grpc struct to send to client
type FileMeta struct {
	Id             string        `json:"id"`
	Title          string        `json:"title"`
	Formats        []*FileFormat `json:"formats"`
	Thumbnail      string        `json:"thumbnail"`
	DurationString string        `json:"duration_string"`
	Channel        string        `json:"channel"`
	LikeCount      int           `json:"like_count"`
	ViewCount      int           `json:"view_count"`
}

func (fileMeta *FileMeta) ConvertToGRPCFileMeta() *gen.FileMeta {
	grpcFileFormats := make([]*gen.FileFormat, 0)
	for _, format := range fileMeta.Formats {
		grpcFileFormats = append(grpcFileFormats, format.ConvertToGRPCFileFormat())
	}
	return &gen.FileMeta{
		Id:             fileMeta.Id,
		Title:          fileMeta.Title,
		Formats:        grpcFileFormats,
		Thumbnail:      fileMeta.Thumbnail,
		DurationString: fileMeta.DurationString,
		Channel:        fileMeta.Channel,
		LikeCount:      int64(fileMeta.LikeCount),
		ViewCount:      int64(fileMeta.ViewCount),
	}
}
