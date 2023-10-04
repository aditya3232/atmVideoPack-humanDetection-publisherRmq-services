package publisher_human_detection

import (
	"mime/multipart"
)

type HumanDetectionInput struct {
	Tid           string                `form:"tid" binding:"required"`
	DateTime      string                `form:"date_time" binding:"required"`
	Person        string                `form:"person" binding:"required"`
	File          *multipart.FileHeader `form:"file" binding:"required"`
	ConvertedFile string                `form:"converted_file"`
}
