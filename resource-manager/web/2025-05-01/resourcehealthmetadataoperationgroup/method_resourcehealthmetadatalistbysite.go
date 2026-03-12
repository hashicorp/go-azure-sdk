package resourcehealthmetadataoperationgroup

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

type ResourceHealthMetadataListBySiteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ResourceHealthMetadata
}

type ResourceHealthMetadataListBySiteCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ResourceHealthMetadata
}

type ResourceHealthMetadataListBySiteCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ResourceHealthMetadataListBySiteCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ResourceHealthMetadataListBySite ...
func (c ResourceHealthMetadataOperationGroupClient) ResourceHealthMetadataListBySite(ctx context.Context, id commonids.AppServiceId) (result ResourceHealthMetadataListBySiteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ResourceHealthMetadataListBySiteCustomPager{},
		Path:       fmt.Sprintf("%s/resourceHealthMetadata", id.ID()),
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

// ResourceHealthMetadataListBySiteComplete retrieves all the results into a single object
func (c ResourceHealthMetadataOperationGroupClient) ResourceHealthMetadataListBySiteComplete(ctx context.Context, id commonids.AppServiceId) (ResourceHealthMetadataListBySiteCompleteResult, error) {
	return c.ResourceHealthMetadataListBySiteCompleteMatchingPredicate(ctx, id, ResourceHealthMetadataOperationPredicate{})
}

// ResourceHealthMetadataListBySiteCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ResourceHealthMetadataOperationGroupClient) ResourceHealthMetadataListBySiteCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate ResourceHealthMetadataOperationPredicate) (result ResourceHealthMetadataListBySiteCompleteResult, err error) {
	items := make([]ResourceHealthMetadata, 0)

	resp, err := c.ResourceHealthMetadataListBySite(ctx, id)
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

	result = ResourceHealthMetadataListBySiteCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
