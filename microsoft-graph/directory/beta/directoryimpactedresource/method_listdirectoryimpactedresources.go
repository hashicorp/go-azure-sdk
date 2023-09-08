package directoryimpactedresource

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

type ListDirectoryImpactedResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ImpactedResourceCollectionResponse
}

type ListDirectoryImpactedResourcesCompleteResult struct {
	Items []models.ImpactedResourceCollectionResponse
}

// ListDirectoryImpactedResources ...
func (c DirectoryImpactedResourceClient) ListDirectoryImpactedResources(ctx context.Context) (result ListDirectoryImpactedResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/directory/impactedResources",
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

// ListDirectoryImpactedResourcesComplete retrieves all the results into a single object
func (c DirectoryImpactedResourceClient) ListDirectoryImpactedResourcesComplete(ctx context.Context) (ListDirectoryImpactedResourcesCompleteResult, error) {
	return c.ListDirectoryImpactedResourcesCompleteMatchingPredicate(ctx, models.ImpactedResourceCollectionResponseOperationPredicate{})
}

// ListDirectoryImpactedResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryImpactedResourceClient) ListDirectoryImpactedResourcesCompleteMatchingPredicate(ctx context.Context, predicate models.ImpactedResourceCollectionResponseOperationPredicate) (result ListDirectoryImpactedResourcesCompleteResult, err error) {
	items := make([]models.ImpactedResourceCollectionResponse, 0)

	resp, err := c.ListDirectoryImpactedResources(ctx)
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

	result = ListDirectoryImpactedResourcesCompleteResult{
		Items: items,
	}
	return
}
