package service

type Error string

const (
	ErrorInvalidUser Error = "InvalidUser"
)

func (e Error) Error() string {
	return string(e)
}
