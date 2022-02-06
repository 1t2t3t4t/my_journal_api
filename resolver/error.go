package resolver

type ResolvingError string

const (
	ResolvingErrorNotLoggedInUser ResolvingError = "NotLoggedIn"
	ResolvingErrorUserNotFound    ResolvingError = "UserNotFound"
)

func (r ResolvingError) Error() string {
	return string(r)
}
