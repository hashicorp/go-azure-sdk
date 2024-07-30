package recommendationimpactedresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListRecommendationImpactedResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ImpactedResource
}

type ListRecommendationImpactedResourcesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ImpactedResource
}

type ListRecommendationImpactedResourcesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRecommendationImpactedResourcesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRecommendationImpactedResources ...
func (c RecommendationImpactedResourceClient) ListRecommendationImpactedResources(ctx context.Context, id DirectoryRecommendationId) (result ListRecommendationImpactedResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRecommendationImpactedResourcesCustomPager{},
		Path:       fmt.Sprintf("%s/impactedResources", id.ID()),
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
		Values *[]beta.ImpactedResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRecommendationImpactedResourcesComplete retrieves all the results into a single object
func (c RecommendationImpactedResourceClient) ListRecommendationImpactedResourcesComplete(ctx context.Context, id DirectoryRecommendationId) (ListRecommendationImpactedResourcesCompleteResult, error) {
	return c.ListRecommendationImpactedResourcesCompleteMatchingPredicate(ctx, id, ImpactedResourceOperationPredicate{})
}

// ListRecommendationImpactedResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RecommendationImpactedResourceClient) ListRecommendationImpactedResourcesCompleteMatchingPredicate(ctx context.Context, id DirectoryRecommendationId, predicate ImpactedResourceOperationPredicate) (result ListRecommendationImpactedResourcesCompleteResult, err error) {
	items := make([]beta.ImpactedResource, 0)

	resp, err := c.ListRecommendationImpactedResources(ctx, id)
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

	result = ListRecommendationImpactedResourcesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
