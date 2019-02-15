package opt

import "encoding/json"

type MinWordSizeFor2TyposOption struct {
	value int
}

func MinWordSizeFor2Typos(v int) MinWordSizeFor2TyposOption {
	return MinWordSizeFor2TyposOption{v}
}

func (o MinWordSizeFor2TyposOption) Get() int {
	return o.value
}

func (o MinWordSizeFor2TyposOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *MinWordSizeFor2TyposOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = 8
		return nil
	}
	return json.Unmarshal(data, &o.value)
}