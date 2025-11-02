package homeservice

import (
	homemodel "e-ticket/internal/domain/home/model"
	homerepository "e-ticket/internal/domain/home/repository"
	"e-ticket/pkg/utils"
)

type Service struct {
	repository *homerepository.Repository
}

func NewHomeService(repository homerepository.Repository) *Service {
	return &Service{repository: &repository}
}

func (s *Service) GetCompanyHomeData(userId int) (*homemodel.HomeEntity, error) {

	homeEntity := homemodel.HomeEntity{}

	// get company
	companies, err := s.repository.FindCompaniesByOwner(userId)
	if err != nil {
		return nil, err
	}
	homeEntity.Companies = &companies

	// get vehicles
	if homeEntity.Companies != nil && len(*homeEntity.Companies) > 0 {
		companyIds := utils.ModelsToElements(
			*homeEntity.Companies,
			func(company homemodel.CompaniesEntity) int {
				return company.Id
			},
		)
		vehicles, err := s.repository.FindVehiclesByCompanies(companyIds)
		if err != nil {
			return nil, err
		}
		homeEntity.Vehicles = &vehicles
	}

	// get owners
	if homeEntity.Vehicles != nil && len(*homeEntity.Vehicles) > 0 {
		vehicleIds := utils.ModelsToElements(
			*homeEntity.Vehicles,
			func(vehicle homemodel.VehiclesEntity) int {
				return vehicle.Id
			},
		)

		owners, err := s.repository.FindOwnersByVehicles(vehicleIds)
		if err != nil {
			return nil, err
		}
		homeEntity.Owners = &owners
	}

	return &homeEntity, nil
}

func (s *Service) GetOwnerHomeData(userId int) ([]homemodel.VehiclesEntity, error) {

	companies, err := s.repository.FindVehiclesByOwner(userId)
	if err != nil {
		return nil, err
	}

	return companies, nil
}
