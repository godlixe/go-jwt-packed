package dto

type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id" binding:"required"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"Email" form:"Email" binding:"required, email"`
	Password string `json:"password,omitempty" form:"password, omitempty" binding:"min=6"`
}

// type UserCreateDTO struct {
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"Email" form:"Email" binding:"required" validate:"email"`
// 	Password string `json:"password,omitempty" form:"password, omitempty" validate:"min:6"`
// }
