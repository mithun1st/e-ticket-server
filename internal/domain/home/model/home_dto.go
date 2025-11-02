package homemodel

// type HomeHeader struct {
// 	Auth string `header:"Authorization"`
// }

// type HomePath struct {
// 	P1 string `uri:"homeId" binding:"omitempty,len=4"`
// }

type HomeQuery struct {
	UserId int `form:"userId" binding:"required"`
	// UserRole int `form:"userRole" binding:"required"`
}

// type HomeRequest struct {
// 	Id int    `json:"id" binding:"omitempty,gte=18,lte=21"`
// 	B1 string `json:"b1" binding:"required,email"`
// 	B2 string `json:"b2" binding:"required,min=8,max=12"`
// 	B4 string `json:"b4" binding:"omitempty,len=4"`
// 	B5 bool   `json:"b5" binding:"omitempty"`
// }
