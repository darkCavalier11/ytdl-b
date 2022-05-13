package models

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
