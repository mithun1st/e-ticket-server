package scheduleservice

import (
	schedulemodel "e-ticket/internal/domain/schedule/model"
	schedulerepository "e-ticket/internal/domain/schedule/repository"
)

type Service struct {
	repository *schedulerepository.Repository
}

func NewScheduleService(repository schedulerepository.Repository) *Service {
	return &Service{repository: &repository}
}

func (s *Service) GetSchedulesById(companyId int) (*[]schedulemodel.ScheduleEntity, error) {

	schedules, err := s.repository.FindSchedulesById(companyId)
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (s *Service) CreateSchedules(companyId int, schedule schedulemodel.ScheduleCreateEntity) (bool, error) {

	isCerated, err := s.repository.InsertSchedule(companyId, schedule)
	if err != nil {
		return false, err
	}

	return isCerated, nil
}
