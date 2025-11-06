package vehicleservice

import (
	vehiclemodel "e-ticket/internal/domain/vehicle/model"
	vehiclerepository "e-ticket/internal/domain/vehicle/repository"
)

type Service struct {
	repository *vehiclerepository.Repository
}

func NewVehicleService(repository vehiclerepository.Repository) *Service {
	return &Service{repository: &repository}
}

func (s *Service) GetAllVehicle(companyId int, userId *int) ([]vehiclemodel.VehicleEntity, error) {

	var list []vehiclemodel.VehicleEntity
	var err error

	if userId == nil {
		list, err = s.repository.FindVehiclesByCompany(companyId)

	} else {
		list, err = s.repository.FindVehiclesByCompanyAndOwner(companyId, *userId)
	}

	if err != nil {
		return nil, err
	}

	return list, nil
}

func (s *Service) CreateVehicle(vehicle vehiclemodel.VehicleCreateRequest) (bool, error) {

	isCreated, err := s.repository.InsertVehicle(vehicle)
	if err != nil {
		return false, err
	}

	return isCreated, nil
}
