// Code generated by go generate. DO NOT EDIT.

package opt

import (
    "context"
    "strings"

    "github.com/algolia/algoliasearch-client-go/algolia/opt"
)

func ExtractAnalyticsTags(opts ...interface{}) *opt.AnalyticsTagsOption {
    for _, o := range opts {
        if v, ok := o.(opt.AnalyticsTagsOption); ok {
            return &v
        }
    }
    return nil
}