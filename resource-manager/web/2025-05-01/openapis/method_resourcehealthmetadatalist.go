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

type ResourceHealthMetadataListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ResourceHealthMetadata
}

type ResourceHealthMetadataListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ResourceHealthMetadata
}

type ResourceHealthMetadataListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResourceHealthMetadataListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResourceHealthMetadataList ...
func (c OpenapisClient) ResourceHealthMetadataList(ctx context.Context, id commonids.SubscriptionId) (result ResourceHealthMetadataListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ResourceHealthMetadataListCustomPager{},
		Path:       fmt.Sprintf("%s/providers/Microsoft.Web/resourceHealthMetadata", id.ID()),
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
		Values *[]ResourceHealthMetadata `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ResourceHealthMetadataListComplete retrieves all the results into a single object
func (c OpenapisClient) ResourceHealthMetadataListComplete(ctx context.Context, id commonids.SubscriptionId) (ResourceHealthMetadataListCompleteResult, error) {
	return c.ResourceHealthMetadataListCompleteMatchingPredicate(ctx, id, ResourceHealthMetadataOperationPredicate{})
}

// ResourceHealthMetadataListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ResourceHealthMetadataListCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate ResourceHealthMetadataOperationPredicate) (result ResourceHealthMetadataListCompleteResult, err error) {
	items := make([]ResourceHealthMetadata, 0)

	resp, err := c.ResourceHealthMetadataList(ctx, id)
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

	result = ResourceHealthMetadataListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
