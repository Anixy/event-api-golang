package helpers

func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}

type MyError struct {
	Code int
	Msg  string
	Data interface{}
}

var (
	LOGIN_UNKNOWN = NewError(202, "user does not exist")
	LOGIN_ERROR   = NewError(203, "wrong account or password")
	VALID_ERROR   = NewError(300, "parameter error")
	ERROR         = NewError(400, "operation failed")
	UNAUTHORIZED  = NewError(401, "you are not logged in.")
	NOT_FOUND     = NewError(404, "resources do not exist")
	INNER_ERROR   = NewError(500, "system exception")
)

func (e *MyError) Error() string {
	return e.Msg
}

func NewError(code int, msg string) *MyError {
	return &MyError{
		Msg:  msg,
		Code: code,
	}
}

func GetError(e *MyError, data interface{}) *MyError {
	return &MyError{
		Msg:  e.Msg,
		Code: e.Code,
		Data: data,
	}
}