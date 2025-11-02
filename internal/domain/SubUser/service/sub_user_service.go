package subuserservice

import (
	subusermodel "e-ticket/internal/domain/SubUser/model"
	subuserrepository "e-ticket/internal/domain/SubUser/repository"
	authmodel "e-ticket/internal/domain/auth/model"
)

type Service struct {
	repository *subuserrepository.Repository
}

func NewSubUserService(repository subuserrepository.Repository) *Service {
	return &Service{repository: &repository}
}

// func (s *Service) GetSubUser(id int) (*subusermodel.SubUserEntity, error) {

// 	subUser, err := s.repository.FindById(id)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if subUser.Id != id {
// 		return nil, errors.New("not match")
// 	}

// 	return subUser, nil
// }

func (s *Service) GetAllSubUser(companyId int, role int) ([]authmodel.UserEntity, error) {

	userIdList, err := s.repository.FindUsersIdByCompanyAndRole(companyId, role)
	if err != nil {
		return nil, err
	}

	users, err := s.repository.FindUsersByIds(userIdList)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) CreateSubUser(subUserRequest subusermodel.SubUserCreateRequest) (*bool, error) {

	user, err := s.repository.FindIsUserByPhone(subUserRequest.Phone)
	if err != nil {
		return nil, err
	}

	if user.Id == 0 {
		newUserCreate, err := s.repository.InsertUser(
			subUserRequest.FirstName,
			subUserRequest.LastName,
			subUserRequest.Phone,
			subUserRequest.Email,
			"00000000",
		)
		if err != nil {
			return nil, err
		}
		if newUserCreate {
			user, err = s.repository.FindIsUserByPhone(subUserRequest.Phone)
			if err != nil {
				return nil, err
			}
		}
	}

	isAssign, err := s.repository.InsertCompanySubUser(subUserRequest.CompanyId, user.Id, subUserRequest.Role)
	if err != nil {
		return nil, err
	}

	return &isAssign, nil
}
