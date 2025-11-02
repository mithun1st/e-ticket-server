package authservice

import (
	authmodel "e-ticket/internal/domain/auth/model"
	authrepository "e-ticket/internal/domain/auth/repository"
	"errors"
)

type Service struct {
	repository *authrepository.Repository
}

func NewAuthService(repository authrepository.Repository) *Service {
	return &Service{repository: &repository}
}

func (s *Service) GetAuthCompany(email string, phone string, password string) (*authmodel.AuthEntity, error) {

	var auth authmodel.AuthEntity

	// check user credential
	userEntity, err := s.repository.FindUser(email, phone, password)
	if err != nil {
		return nil, err
	}

	if !((userEntity.Email != nil && *userEntity.Email == email) || (userEntity.Phone != "" && userEntity.Phone == phone)) {
		return nil, errors.New("userid or password incorrect")
	}
	auth.UserEntity = *userEntity

	// find company and role
	companyId, role, err := s.repository.FindCompanyUser(userEntity.Id)
	if err != nil {
		return nil, err
	}
	auth.CompanyId = companyId
	auth.Role = *role

	return &auth, nil
}

func (s *Service) GetAuthSubUser(companyId int, email string, phone string, password string) (*authmodel.AuthEntity, error) {

	var auth authmodel.AuthEntity

	// check user auth
	userEntity, err := s.repository.FindUser(email, phone, password)
	if err != nil {
		return nil, err
	}

	if !((userEntity.Email != nil && *userEntity.Email == email) || (userEntity.Phone != "" && userEntity.Phone == phone)) {
		return nil, errors.New("userid or password incorrect")
	}
	auth.UserEntity = *userEntity

	// check company auth
	subUserRole, err := s.repository.FindSubUserRole(userEntity.Id, companyId)
	if err != nil {
		return nil, err
	}
	if subUserRole == nil {
		return nil, errors.New("role not found")
	}
	// auth.CompanyId = companyId
	auth.Role = *subUserRole

	return &auth, nil
}

// func (s *Service) GetAllAuth() ([]authmodel.AuthEntity, error) {

// 	list, err := s.repository.FindAll()
// 	if err != nil {
// 		return nil, err
// 	}

// 	if len(list) == 0 {
// 		return nil, errors.New("empty list")
// 	}

// 	return list, nil
// }
