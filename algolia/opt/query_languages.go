package opt

import "encoding/json"

type QueryLanguagesOption struct {
	value []string
}

func QueryLanguages(v []string) QueryLanguagesOption {
	return QueryLanguagesOption{v}
}

func (o QueryLanguagesOption) Get() []string {
	return o.value
}

func (o QueryLanguagesOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *QueryLanguagesOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}