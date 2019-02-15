package opt

import "encoding/json"

type DistinctOption struct {
	value int
}

func Distinct(v int) DistinctOption {
	return DistinctOption{v}
}

func (o DistinctOption) Get() int {
	return o.value
}

func (o DistinctOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *DistinctOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 0
		return nil
	}
	return json.Unmarshal(data, &o.value)
}