package resolver

type helloWorldResolver struct{}

func newHelloWorldResolver() *helloWorldResolver {
	return &helloWorldResolver{}
}

func (r *helloWorldResolver) Message() string {
	return "Hello World!"
}
