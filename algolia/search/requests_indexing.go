package search

import "github.com/algolia/algoliasearch-client-go/algolia/opt"

type IndexOperation struct {
	Destination string   `json:"destination"`
	Operation   string   `json:"operation"`
	Scopes      []string `json:"scope,omitempty"`
}

type deleteByReq struct {
	AroundLatLng      *opt.AroundLatLngOption      `json:"aroundLatLng,omitempty"`
	AroundRadius      *opt.AroundRadiusOption      `json:"aroundRadius,omitempty"`
	FacetFilters      *opt.FacetFiltersOption      `json:"facetFilters,omitempty"`
	Filters           *opt.FiltersOption           `json:"filters,omitempty"`
	InsideBoundingBox *opt.InsideBoundingBoxOption `json:"insideBoundingBox,omitempty"`
	InsidePolygon     *opt.InsidePolygonOption     `json:"insidePolygon,omitempty"`
	NumericFilters    *opt.NumericFiltersOption    `json:"numericFilters,omitempty"`
}

type getObjectsReq struct {
	IndexName            string                          `json:"indexName"`
	ObjectID             string                          `json:"objectID"`
	AttributesToRetrieve *opt.AttributesToRetrieveOption `json:"attributesToRetrieve,omitempty"`
}

type batchReq struct {
	Requests []BatchOperation `json:"requests,omitempty"`
}
