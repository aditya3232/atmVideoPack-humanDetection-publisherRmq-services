package publisher_human_detection

type PublisherHumanDetectionFormatter struct {
	Tid           string `json:"tid"`
	DateTime      string `json:"date_time"`
	Person        string `json:"person"`
	ConvertedFile string `json:"converted_file"`
}

// data ditampilkan dari input
func FormatPublisherHumanDetection(entityHumanDetection HumanDetection) PublisherHumanDetectionFormatter {
	formatter := PublisherHumanDetectionFormatter(entityHumanDetection)
	return formatter
}
