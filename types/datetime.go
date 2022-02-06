package types

import (
	"encoding/json"
	"fmt"
	"time"
)

type DateTime struct {
	Time time.Time
}

var _ graphqlScalarType = (*DateTime)(nil)

func (d *DateTime) ImplementsGraphQLType(name string) bool {
	return name == "DateTime"
}

func (d *DateTime) UnmarshalGraphQL(input any) error {
	switch input := input.(type) {
	case string:
		dateTime, err := time.Parse(time.RFC3339, input)
		if err != nil {
			return err
		}
		d.Time = dateTime.UTC()
	default:
		return fmt.Errorf("unsupported type")
	}
	return nil
}

func (d DateTime) MarshalJSON() ([]byte, error) {
	utcTime := d.Time.UTC()
	return json.Marshal(utcTime)
}

func (d *DateTime) UnmarshalJSON(i []byte) error {
	var t time.Time
	err := json.Unmarshal(i, &t)
	if err != nil {
		return err
	}
	d.Time = t.UTC()
	return nil
}
