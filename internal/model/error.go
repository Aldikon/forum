package model

var FillingErr *fillingError

var (
	ErrEmail         = &fillingError{"Email is not free"}
	ErrNotFoundUser  = &fillingError{"Not found user"}
	ErrTokenExpired  = &fillingError{"Token expired"}
	ErrCategoryExist = &fillingError{"A category that does not exist"}
)

type fillingError struct {
	Message string
}

func (f *fillingError) Error() string {
	return f.Message
}
