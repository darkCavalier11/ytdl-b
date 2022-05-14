package models

import "github.com/darkCavalier11/downloader_backend/grpc_module/gen"

type FileFormat struct {
	FormatId   string  `json:"format_id"`
	FormatNote string  `json:"format_note"`
	Format     string  `json:"format"`
	FileSize   int     `json:"filesize"`
	Url        string  `json:"url"`
	Ext        string  `json:"ext"`
	Resolution string  `json:"resolution"`
	AudioExt   string  `json:"audio_ext"`
	VideoExt   string  `json:"video_ext"`
	Abr        float32 `json:"abr,omitempty"`
}

// ConvertToGRPCFileFormat returns a gen.FileFormat to send directly to the client
func (fileFormat *FileFormat) ConvertToGRPCFileFormat() *gen.FileFormat {
	return &gen.FileFormat{
		FormatId:   fileFormat.FormatId,
		FormatNote: fileFormat.FormatNote,
		Format:     fileFormat.Format,
		FileSize:   int64(fileFormat.FileSize),
		Ext:        fileFormat.Ext,
		Resolution: fileFormat.Resolution,
		AudioExt:   fileFormat.AudioExt,
		VideoExt:   fileFormat.VideoExt,
		Abr:        fileFormat.Abr,
	}
}
