package types

import "encoding/json"

type graphqlScalarType interface {
	json.Marshaler
	json.Unmarshaler
	ImplementsGraphQLType(name string) bool
	UnmarshalGraphQL(input any) error
}
