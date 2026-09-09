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

type ResourceGuardsGetDeleteProtectedItemRequestsObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Resource
}

type ResourceGuardsGetDeleteProtectedItemRequestsObjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Resource
}

type ResourceGuardsGetDeleteProtectedItemRequestsObjectsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResourceGuardsGetDeleteProtectedItemRequestsObjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResourceGuardsGetDeleteProtectedItemRequestsObjects ...
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetDeleteProtectedItemRequestsObjects(ctx context.Context, id ResourceGuardId) (result ResourceGuardsGetDeleteProtectedItemRequestsObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ResourceGuardsGetDeleteProtectedItemRequestsObjectsCustomPager{},
		Path:       fmt.Sprintf("%s/deleteProtectedItemRequests", id.ID()),
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

// ResourceGuardsGetDeleteProtectedItemRequestsObjectsComplete retrieves all the results into a single object
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetDeleteProtectedItemRequestsObjectsComplete(ctx context.Context, id ResourceGuardId) (ResourceGuardsGetDeleteProtectedItemRequestsObjectsCompleteResult, error) {
	return c.ResourceGuardsGetDeleteProtectedItemRequestsObjectsCompleteMatchingPredicate(ctx, id, ResourceOperationPredicate{})
}

// ResourceGuardsGetDeleteProtectedItemRequestsObjectsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DppBaseResourceOperationGroupClient) ResourceGuardsGetDeleteProtectedItemRequestsObjectsCompleteMatchingPredicate(ctx context.Context, id ResourceGuardId, predicate ResourceOperationPredicate) (result ResourceGuardsGetDeleteProtectedItemRequestsObjectsCompleteResult, err error) {
	items := make([]Resource, 0)

	resp, err := c.ResourceGuardsGetDeleteProtectedItemRequestsObjects(ctx, id)
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

	result = ResourceGuardsGetDeleteProtectedItemRequestsObjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
