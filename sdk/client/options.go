package client

import (
	"net/http"
	"net/url"
	
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type Options interface {
	ToHeaders() *Headers
	ToOData() *odata.Query
	ToQuery() *QueryParams
}

type Headers map[string]string

func (h Headers) Append(other http.Header) http.Header {
	for k, v := range h {
		other.Set(k, v)
	}
	return other
}

func (h Headers) Headers() http.Header {
	return h.Append(make(http.Header))
}

type QueryParams map[string]string

func (q QueryParams) Values() url.Values {
	va := make(url.Values)
	for k, v := range q {
		va.Set(k, v)
	}
	return va
}
