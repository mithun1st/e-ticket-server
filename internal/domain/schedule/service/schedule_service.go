package scheduleservice

import (
	schedulemodel "e-ticket/internal/domain/schedule/model"
	schedulerepository "e-ticket/internal/domain/schedule/repository"
	"time"
)

type Service struct {
	repository *schedulerepository.Repository
}

func NewScheduleService(repository schedulerepository.Repository) *Service {
	return &Service{repository: &repository}
}

func (s *Service) GetSchedulesById(companyId int, date time.Time) (*[]schedulemodel.ScheduleEntity, error) {

	schedules, err := s.repository.FindSchedulesById(companyId, date)
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
