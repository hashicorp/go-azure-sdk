package cloudendpoints

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBySyncGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]CloudEndpoint
}

type ListBySyncGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []CloudEndpoint
}

type ListBySyncGroupCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListBySyncGroupCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListBySyncGroup ...
func (c CloudEndpointsClient) ListBySyncGroup(ctx context.Context, id SyncGroupId) (result ListBySyncGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListBySyncGroupCustomPager{},
		Path:       fmt.Sprintf("%s/cloudEndpoints", id.ID()),
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
		Values *[]CloudEndpoint `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBySyncGroupComplete retrieves all the results into a single object
func (c CloudEndpointsClient) ListBySyncGroupComplete(ctx context.Context, id SyncGroupId) (ListBySyncGroupCompleteResult, error) {
	return c.ListBySyncGroupCompleteMatchingPredicate(ctx, id, CloudEndpointOperationPredicate{})
}

// ListBySyncGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudEndpointsClient) ListBySyncGroupCompleteMatchingPredicate(ctx context.Context, id SyncGroupId, predicate CloudEndpointOperationPredicate) (result ListBySyncGroupCompleteResult, err error) {
	items := make([]CloudEndpoint, 0)

	resp, err := c.ListBySyncGroup(ctx, id)
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

	result = ListBySyncGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
