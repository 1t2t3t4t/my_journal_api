package resolver

type ResolvingError string

const (
	ResolvingErrorNotLoggedInUser ResolvingError = "NotLoggedIn"
)

func (r ResolvingError) Error() string {
	return string(r)
}
