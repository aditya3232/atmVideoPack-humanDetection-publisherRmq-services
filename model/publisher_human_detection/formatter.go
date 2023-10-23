package publisher_human_detection

// formatter akan menampilkan response di api
type RmqPublisherHumanDetectionFormatter struct {
	TidID                         *int   `json:"tid_id"`
	Tid                           string `json:"tid"`
	DateTime                      string `json:"date_time"`
	Person                        string `json:"person"`
	FileNameCaptureHumanDetection string `json:"file_name_capture_human_detection"` // ini untuk balikan file name nya aja di api
}

func PublisherHumanDetectionFormat(rmqPublisherHumanDetection RmqPublisherHumanDetection) RmqPublisherHumanDetectionFormatter {
	var formatter RmqPublisherHumanDetectionFormatter

	formatter.TidID = rmqPublisherHumanDetection.TidID
	formatter.Tid = rmqPublisherHumanDetection.Tid
	formatter.DateTime = rmqPublisherHumanDetection.DateTime
	formatter.Person = rmqPublisherHumanDetection.Person
	formatter.FileNameCaptureHumanDetection = rmqPublisherHumanDetection.FileNameCaptureHumanDetection

	return formatter
}
