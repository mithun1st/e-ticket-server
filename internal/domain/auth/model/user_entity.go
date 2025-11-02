package authmodel

type UserEntity struct {
	Id        int     `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  *string `json:"lastName"`
	Phone     string  `json:"phone"`
	Email     *string `json:"email"`
	IsActive  bool    `json:"isActive"`
}
