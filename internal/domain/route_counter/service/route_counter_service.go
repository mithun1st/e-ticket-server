package routecounterservice

import (
	routecountermodel "e-ticket/internal/domain/route_counter/model"
	routecounterrepository "e-ticket/internal/domain/route_counter/repository"
)

type Service struct {
	repository *routecounterrepository.Repository
}

func NewRouteCounterService(repository routecounterrepository.Repository) *Service {
	return &Service{repository: &repository}
}

func (s *Service) GetCounters(companyId int) ([]routecountermodel.CounterEntity, error) {

	counters, err := s.repository.FindCountersByCompanyId(companyId)
	if err != nil {
		return nil, err
	}

	return counters, nil
}

func (s *Service) CreateCounters(companyId int, counter routecountermodel.CounterCreateEntity) (bool, error) {

	counters, err := s.repository.InsertCounter(companyId, counter)
	if err != nil {
		return false, err
	}

	return counters, nil
}

// func (s *Service) CreateCounter(counter routecountermodel.CounterEntity)(bool,error){

// }
// func (s *Service) GetAllRouteCounter() ([]routecountermodel.RouteCounterEntity, error) {

// 	list, err := s.repository.FindAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(list) == 0 {
// 		return nil, errors.New("empty list")
// 	}

// 	return list, nil
// }
