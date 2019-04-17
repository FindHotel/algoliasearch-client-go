package opt

import "encoding/json"

type SnippetEllipsisTextOption struct {
	value string
}

func SnippetEllipsisText(v string) *SnippetEllipsisTextOption {
	return &SnippetEllipsisTextOption{v}
}

func (o *SnippetEllipsisTextOption) Get() string {
	if o == nil {
		return "…"
	}
	return o.value
}

func (o SnippetEllipsisTextOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *SnippetEllipsisTextOption) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		o.value = "…"
		return nil
	}
	return json.Unmarshal(data, &o.value)
}

func (o *SnippetEllipsisTextOption) Equal(o2 *SnippetEllipsisTextOption) bool {
	if o2 == nil {
		return o.value == "…" || o.value == ""
	}
	return o.value == o2.value
}

func SnippetEllipsisTextEqual(o1, o2 *SnippetEllipsisTextOption) bool {
	if o1 != nil {
		return o1.Equal(o2)
	}
	if o2 != nil {
		return o2.Equal(o1)
	}
	return true
}
