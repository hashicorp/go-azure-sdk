package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceUsageListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ResourceUsage
}

type ResourceUsageListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ResourceUsage
}

type ResourceUsageListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResourceUsageListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResourceUsageList ...
func (c OpenapisClient) ResourceUsageList(ctx context.Context, id commonids.SubscriptionId) (result ResourceUsageListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Pager:      &ResourceUsageListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Cdn/checkResourceUsage", id.ID()),
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
		Values *[]ResourceUsage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ResourceUsageListComplete retrieves all the results into a single object
func (c OpenapisClient) ResourceUsageListComplete(ctx context.Context, id commonids.SubscriptionId) (ResourceUsageListCompleteResult, error) {
	return c.ResourceUsageListCompleteMatchingPredicate(ctx, id, ResourceUsageOperationPredicate{})
}

// ResourceUsageListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ResourceUsageListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ResourceUsageOperationPredicate) (result ResourceUsageListCompleteResult, err error) {
	items := make([]ResourceUsage, 0)

	resp, err := c.ResourceUsageList(ctx, id)
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

	result = ResourceUsageListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
