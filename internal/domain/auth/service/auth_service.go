package authservice

import (
	"e-ticket/internal/config"
	authmodel "e-ticket/internal/domain/auth/model"
	authrepository "e-ticket/internal/domain/auth/repository"
	subusermodel "e-ticket/internal/domain/sub_user/model"
	"e-ticket/internal/middleware"
	apptoken "e-ticket/pkg/token"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type Service struct {
	repository *authrepository.Repository
}

func NewAuthService(repository authrepository.Repository) *Service {
	return &Service{repository: &repository}
}

func _generateToken(user subusermodel.UserEntity) (string, error) {
	appConfigModel, _ := config.Load()
	var tokenEntity middleware.TokenEnitty = middleware.TokenEnitty{
		Id:        user.Id,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		Email:     user.Email,
		Time:      time.Now(),
	}

	jsonData, _ := json.Marshal(tokenEntity)
	var jsonMap map[string]any
	err := json.Unmarshal(jsonData, &jsonMap)
	if err != nil {
		return "", err
	}

	token, err := apptoken.Encript(jsonMap, appConfigModel.Keys.JwtSecretKey)

	return token, err
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

	// Load config model from env
	token, err := _generateToken(*userEntity)
	if err != nil {
		return nil, err
	}
	auth.Token = "Bearer " + token

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
	fmt.Println(auth)

	// check company auth
	subUserRole, err := s.repository.FindSubUserRole(userEntity.Id, companyId)
	if err != nil {
		return nil, err
	}
	if subUserRole == nil {
		return nil, errors.New("role not found")
	}
	auth.Role = *subUserRole

	// Load config model from env
	token, err := _generateToken(*userEntity)
	if err != nil {
		return nil, err
	}
	auth.Token = "Bearer " + token

	return &auth, nil
}
