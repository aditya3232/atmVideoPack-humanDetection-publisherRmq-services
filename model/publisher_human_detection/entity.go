package publisher_human_detection

// ini entity data yg akan dikirim ke rmq, jadi gk pake tabel, karena gk dikirim ke db
type RmqPublisherHumanDetection struct {
	TidID                              *int   `json:"tid_id"`
	Tid                                string `json:"tid"`
	DateTime                           string `json:"date_time"`
	Person                             string `json:"person"`
	ConvertedFileCaptureHumanDetection string `json:"converted_file_capture_human_detection"`
	FileNameCaptureHumanDetection      string `json:"file_name_capture_human_detection"` // ini untuk balikan file name nya aja di api
}
