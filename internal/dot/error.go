package dot

type errorPage struct {
	Code    int
	Message string
}

func NewErrorDot(code int, m string) *errorPage {
	return &errorPage{
		Code:    code,
		Message: m,
	}
}
