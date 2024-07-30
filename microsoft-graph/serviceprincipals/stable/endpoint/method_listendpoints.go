package endpoint

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

type ListEndpointsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Endpoint
}

type ListEndpointsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Endpoint
}

type ListEndpointsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEndpointsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEndpoints ...
func (c EndpointClient) ListEndpoints(ctx context.Context, id ServicePrincipalId) (result ListEndpointsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListEndpointsCustomPager{},
		Path:       fmt.Sprintf("%s/endpoints", id.ID()),
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
		Values *[]stable.Endpoint `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEndpointsComplete retrieves all the results into a single object
func (c EndpointClient) ListEndpointsComplete(ctx context.Context, id ServicePrincipalId) (ListEndpointsCompleteResult, error) {
	return c.ListEndpointsCompleteMatchingPredicate(ctx, id, EndpointOperationPredicate{})
}

// ListEndpointsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EndpointClient) ListEndpointsCompleteMatchingPredicate(ctx context.Context, id ServicePrincipalId, predicate EndpointOperationPredicate) (result ListEndpointsCompleteResult, err error) {
	items := make([]stable.Endpoint, 0)

	resp, err := c.ListEndpoints(ctx, id)
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

	result = ListEndpointsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
