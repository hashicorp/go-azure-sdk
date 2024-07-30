package conditionalaccesnamedlocation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListConditionalAccesNamedLocationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.NamedLocation
}

type ListConditionalAccesNamedLocationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.NamedLocation
}

type ListConditionalAccesNamedLocationsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListConditionalAccesNamedLocationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListConditionalAccesNamedLocations ...
func (c ConditionalAccesNamedLocationClient) ListConditionalAccesNamedLocations(ctx context.Context) (result ListConditionalAccesNamedLocationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListConditionalAccesNamedLocationsCustomPager{},
		Path:       "/identity/conditionalAccess/namedLocations",
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]stable.NamedLocation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListConditionalAccesNamedLocationsComplete retrieves all the results into a single object
func (c ConditionalAccesNamedLocationClient) ListConditionalAccesNamedLocationsComplete(ctx context.Context) (ListConditionalAccesNamedLocationsCompleteResult, error) {
	return c.ListConditionalAccesNamedLocationsCompleteMatchingPredicate(ctx, NamedLocationOperationPredicate{})
}

// ListConditionalAccesNamedLocationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ConditionalAccesNamedLocationClient) ListConditionalAccesNamedLocationsCompleteMatchingPredicate(ctx context.Context, predicate NamedLocationOperationPredicate) (result ListConditionalAccesNamedLocationsCompleteResult, err error) {
	items := make([]stable.NamedLocation, 0)

	resp, err := c.ListConditionalAccesNamedLocations(ctx)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListConditionalAccesNamedLocationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
