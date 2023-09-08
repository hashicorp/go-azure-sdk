package directoryrecommendationimpactedresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListDirectoryRecommendationByIdImpactedResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ImpactedResourceCollectionResponse
}

type ListDirectoryRecommendationByIdImpactedResourcesCompleteResult struct {
	Items []models.ImpactedResourceCollectionResponse
}

// ListDirectoryRecommendationByIdImpactedResources ...
func (c DirectoryRecommendationImpactedResourceClient) ListDirectoryRecommendationByIdImpactedResources(ctx context.Context, id DirectoryRecommendationId) (result ListDirectoryRecommendationByIdImpactedResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.ImpactedResourceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDirectoryRecommendationByIdImpactedResourcesComplete retrieves all the results into a single object
func (c DirectoryRecommendationImpactedResourceClient) ListDirectoryRecommendationByIdImpactedResourcesComplete(ctx context.Context, id DirectoryRecommendationId) (ListDirectoryRecommendationByIdImpactedResourcesCompleteResult, error) {
	return c.ListDirectoryRecommendationByIdImpactedResourcesCompleteMatchingPredicate(ctx, id, models.ImpactedResourceCollectionResponseOperationPredicate{})
}

// ListDirectoryRecommendationByIdImpactedResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryRecommendationImpactedResourceClient) ListDirectoryRecommendationByIdImpactedResourcesCompleteMatchingPredicate(ctx context.Context, id DirectoryRecommendationId, predicate models.ImpactedResourceCollectionResponseOperationPredicate) (result ListDirectoryRecommendationByIdImpactedResourcesCompleteResult, err error) {
	items := make([]models.ImpactedResourceCollectionResponse, 0)

	resp, err := c.ListDirectoryRecommendationByIdImpactedResources(ctx, id)
	if err != nil {
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

	result = ListDirectoryRecommendationByIdImpactedResourcesCompleteResult{
		Items: items,
	}
	return
}
