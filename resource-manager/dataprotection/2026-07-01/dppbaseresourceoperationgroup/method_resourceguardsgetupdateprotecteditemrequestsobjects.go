package dppbaseresourceoperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceGuardsGetUpdateProtectedItemRequestsObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Resource
}

type ResourceGuardsGetUpdateProtectedItemRequestsObjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Resource
}

type ResourceGuardsGetUpdateProtectedItemRequestsObjectsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResourceGuardsGetUpdateProtectedItemRequestsObjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResourceGuardsGetUpdateProtectedItemRequestsObjects ...
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetUpdateProtectedItemRequestsObjects(ctx context.Context, id ResourceGuardId) (result ResourceGuardsGetUpdateProtectedItemRequestsObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ResourceGuardsGetUpdateProtectedItemRequestsObjectsCustomPager{},
		Path:       fmt.Sprintf("%s/updateProtectedItemRequests", id.ID()),
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
		Values *[]Resource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ResourceGuardsGetUpdateProtectedItemRequestsObjectsComplete retrieves all the results into a single object
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetUpdateProtectedItemRequestsObjectsComplete(ctx context.Context, id ResourceGuardId) (ResourceGuardsGetUpdateProtectedItemRequestsObjectsCompleteResult, error) {
	return c.ResourceGuardsGetUpdateProtectedItemRequestsObjectsCompleteMatchingPredicate(ctx, id, ResourceOperationPredicate{})
}

// ResourceGuardsGetUpdateProtectedItemRequestsObjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetUpdateProtectedItemRequestsObjectsCompleteMatchingPredicate(ctx context.Context, id ResourceGuardId, predicate ResourceOperationPredicate) (result ResourceGuardsGetUpdateProtectedItemRequestsObjectsCompleteResult, err error) {
	items := make([]Resource, 0)

	resp, err := c.ResourceGuardsGetUpdateProtectedItemRequestsObjects(ctx, id)
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

	result = ResourceGuardsGetUpdateProtectedItemRequestsObjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
