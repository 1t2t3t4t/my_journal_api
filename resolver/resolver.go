package resolver

type Resolver struct{}

func NewResolver() *Resolver {
	return &Resolver{}
}

func (r *Resolver) HelloWorld() *helloWorldResolver {
	return newHelloWorldResolver()
}