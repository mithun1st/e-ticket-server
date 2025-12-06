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

	isCreated, err := s.repository.InsertCounter(companyId, counter)
	if err != nil {
		return false, err
	}

	return isCreated, nil
}

func (s *Service) GetRoutes(companyId int) ([]routecountermodel.RouteEntity, error) {

	routes, err := s.repository.FindRoutesByCompanyId(companyId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func (s *Service) CreateRoute(companyId int, route routecountermodel.RouteCreateEntity) (bool, error) {

	isCreated, err := s.repository.InsertRoute(companyId, route)
	if err != nil {
		return false, err
	}

	return isCreated, nil
}

func (s *Service) GetRouteCounters(companyId int, routeId int) ([]routecountermodel.RouteCounterEntity, error) {

	routes, err := s.repository.FindRouteCountersByCompanyAndRoute(companyId, routeId)
	if err != nil {
		return nil, err
	}

	return routes, nil
}

func (s *Service) CreateRouteCounter(companyId int, routeId int, routeCounters []routecountermodel.RouteCounterCreateEntity) (bool, error) {

	isCreated, err := s.repository.InsertRouteCounter(companyId, routeId, routeCounters)
	if err != nil {
		return false, err
	}

	return isCreated, nil
}
