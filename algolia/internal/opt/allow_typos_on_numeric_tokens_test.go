package opt

import (
	"encoding/json"
	"testing"

	"github.com/algolia/algoliasearch-client-go/algolia/opt"
	"github.com/stretchr/testify/require"
)

func TestAllowTyposOnNumericTokens(t *testing.T) {
	for _, c := range []struct {
		opts     []interface{}
		expected opt.AllowTyposOnNumericTokensOption
	}{
		{
			opts:     []interface{}{nil},
			expected: opt.AllowTyposOnNumericTokens(false),
		},
		{
			opts:     []interface{}{opt.AllowTyposOnNumericTokens(true)},
			expected: opt.AllowTyposOnNumericTokens(true),
		},
		{
			opts:     []interface{}{opt.AllowTyposOnNumericTokens(false)},
			expected: opt.AllowTyposOnNumericTokens(false),
		},
	} {
		var (
			in  = ExtractAllowTyposOnNumericTokens(c.opts...)
			out opt.AllowTyposOnNumericTokensOption
		)
		data, err := json.Marshal(&in)
		require.NoError(t, err)
		err = json.Unmarshal(data, &out)
		require.NoError(t, err)
		require.Equal(t, c.expected, out)
	}
}