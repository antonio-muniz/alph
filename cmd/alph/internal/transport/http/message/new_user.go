package message

type NewUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r NewUserRequest) NewPointer() interface{} {
	return &NewUserRequest{}
}

func (r NewUserRequest) Dereference(pointer interface{}) interface{} {
	return *(pointer.(*NewUserRequest))
}

type NewUserResponse struct{}
