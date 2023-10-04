package publisher_human_detection

import "github.com/aditya3232/gatewatchApp-services.git/helper"

type Service interface {
	CreateQueueHumanDetection(input HumanDetectionInput) (HumanDetection, error)
}

type service struct {
	humanDetectionRepository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

// public message to rmq
func (s *service) CreateQueueHumanDetection(input HumanDetectionInput) (HumanDetection, error) {
	// convert img to base64 with helper.ConvertFileToBase64()
	imgBase64String, err := helper.ConvertFileToBase64(input.File)
	if err != nil {
		return HumanDetection{}, err
	}

	// create human detection
	input.ConvertedFile = imgBase64String

	_, err = s.humanDetectionRepository.CreateQueueHumanDetection(input)
	if err != nil {
		return HumanDetection{}, err
	}

	return HumanDetection{
		Tid:           input.Tid,
		DateTime:      input.DateTime,
		Person:        input.Person,
		ConvertedFile: input.ConvertedFile,
	}, nil
}
