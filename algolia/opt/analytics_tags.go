package opt

import "encoding/json"

type AnalyticsTagsOption struct {
	value []string
}

func AnalyticsTags(v []string) AnalyticsTagsOption {
	return AnalyticsTagsOption{v}
}

func (o AnalyticsTagsOption) Get() []string {
	return o.value
}

func (o AnalyticsTagsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *AnalyticsTagsOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = nil
		return nil
	}
	return json.Unmarshal(data, &o.value)
}