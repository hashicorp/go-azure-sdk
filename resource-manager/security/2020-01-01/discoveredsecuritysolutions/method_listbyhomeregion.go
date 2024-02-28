package discoveredsecuritysolutions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByHomeRegionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DiscoveredSecuritySolution
}

type ListByHomeRegionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DiscoveredSecuritySolution
}

// ListByHomeRegion ...
func (c DiscoveredSecuritySolutionsClient) ListByHomeRegion(ctx context.Context, id LocationId) (result ListByHomeRegionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/discoveredSecuritySolutions", id.ID()),
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
		Values *[]DiscoveredSecuritySolution `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByHomeRegionComplete retrieves all the results into a single object
func (c DiscoveredSecuritySolutionsClient) ListByHomeRegionComplete(ctx context.Context, id LocationId) (ListByHomeRegionCompleteResult, error) {
	return c.ListByHomeRegionCompleteMatchingPredicate(ctx, id, DiscoveredSecuritySolutionOperationPredicate{})
}

// ListByHomeRegionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiscoveredSecuritySolutionsClient) ListByHomeRegionCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate DiscoveredSecuritySolutionOperationPredicate) (result ListByHomeRegionCompleteResult, err error) {
	items := make([]DiscoveredSecuritySolution, 0)

	resp, err := c.ListByHomeRegion(ctx, id)
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

	result = ListByHomeRegionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
