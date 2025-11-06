package subuserservice

import (
	subusermodel "e-ticket/internal/domain/sub_user/model"
	subuserrepository "e-ticket/internal/domain/sub_user/repository"
	"e-ticket/internal/enum"
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

func (s *Service) GetAllSubUser(companyId int, role int) ([]subusermodel.UserEntity, error) {

	userIdList, err := s.repository.FindUsersIdByCompanyAndRole(companyId, role)
	if err != nil {
		return nil, err
	}

	if len(userIdList) == 0 {
		return []subusermodel.UserEntity{}, nil
	}
	users, err := s.repository.FindUsersByIds(userIdList)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *Service) CreateSubUser(subUserRequest subusermodel.SubUserCreateRequest) (bool, error) {

	user, err := s.repository.FindIsUserByPhone(subUserRequest.Phone)
	if err != nil {
		return false, err
	}

	if user.Id == 0 {
		newUserCreate, err := s.repository.InsertUser(
			subUserRequest.FirstName,
			subUserRequest.LastName,
			subUserRequest.Phone,
			subUserRequest.Email,
			enum.DefaultPassword,
		)
		if err != nil {
			return false, err
		}
		if newUserCreate {
			user, err = s.repository.FindIsUserByPhone(subUserRequest.Phone)
			if err != nil {
				return false, err
			}
		}
	}

	isCreated, err := s.repository.InsertCompanySubUser(subUserRequest.CompanyId, user.Id, subUserRequest.Role)
	if err != nil {
		return false, err
	}

	return isCreated, nil
}
