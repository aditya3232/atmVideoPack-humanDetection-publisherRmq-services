package publisher_human_detection

type HumanDetection struct {
	Tid           string `json:"tid"`
	DateTime      string `json:"date_time"`
	Person        string `json:"person"`
	ConvertedFile string `json:"converted_file"`
}
