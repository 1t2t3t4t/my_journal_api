package database

type Error string

const (
	ErrorDuplicateUsername Error = "DupUsername"
)

func (e Error) Error() string {
	return string(e)
}
