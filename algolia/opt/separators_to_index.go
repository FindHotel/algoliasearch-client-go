package opt

import "encoding/json"

type SeparatorsToIndexOption struct {
	value string
}

func SeparatorsToIndex(v string) SeparatorsToIndexOption {
	return SeparatorsToIndexOption{v}
}

func (o SeparatorsToIndexOption) Get() string {
	return o.value
}

func (o SeparatorsToIndexOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *SeparatorsToIndexOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = &#34;&#34;
		return nil
	}
	return json.Unmarshal(data, &o.value)
}