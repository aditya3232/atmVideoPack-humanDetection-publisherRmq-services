package publisher_human_detection

import (
	"github.com/aditya3232/atmVideoPack-humanDetection-publisherRmq-services.git/helper"
	"github.com/aditya3232/atmVideoPack-humanDetection-publisherRmq-services.git/model/tb_tid"
)

type Service interface {
	/*
		- input ada 2, pertama input dari API
		- kedua input yang akan dimasukkan ke RMQ
		- disini returnnya yg akan ditampilkan di API adalah inputan rmq,
		- disini parameter adalah input
	*/
	CreateQueueHumanDetection(input RmqPublisherHumanDetectionInput) (RmqPublisherHumanDetection, error)
}

type service struct {
	humanDetectionRepository Repository
	tbTidRepository          tb_tid.Repository
}

func NewService(repository Repository, tbTidRepository tb_tid.Repository) *service {
	return &service{repository, tbTidRepository}
}

// public message to rmq
func (s *service) CreateQueueHumanDetection(input RmqPublisherHumanDetectionInput) (RmqPublisherHumanDetection, error) {
	var rmqPublisherHumanDetection RmqPublisherHumanDetection

	// validasi is image
	err := helper.IsImage(input.FileCaptureHumanDetection)
	if err != nil {
		return rmqPublisherHumanDetection, err
	}

	// convert image to jpg
	err = helper.ConvertImageToJpg(input.FileCaptureHumanDetection)
	if err != nil {
		return rmqPublisherHumanDetection, err
	}

	// convert img
	imgBase64String, err := helper.ConvertFileToBase64(input.FileCaptureHumanDetection)
	if err != nil {
		return rmqPublisherHumanDetection, err
	}

	// get name file
	fileName := input.FileCaptureHumanDetection.Filename

	// get id from input tid
	tidID, err := s.tbTidRepository.GetOneByTid(input.Tid)
	if err != nil {
		return rmqPublisherHumanDetection, err
	}

	newRmqPublisherHumanDetection := RmqPublisherHumanDetection{
		TidID:                              &tidID.ID,
		Tid:                                input.Tid,
		DateTime:                           input.DateTime,
		Person:                             input.Person,
		ConvertedFileCaptureHumanDetection: imgBase64String,
		FileNameCaptureHumanDetection:      fileName,
	}

	newHumanDetection, err := s.humanDetectionRepository.CreateQueueHumanDetection(newRmqPublisherHumanDetection)
	if err != nil {
		return newHumanDetection, err
	}

	return newHumanDetection, nil
}
