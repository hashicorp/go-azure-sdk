package meonenoteresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListMeOnenoteResourcesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.OnenoteResourceCollectionResponse
}

type ListMeOnenoteResourcesCompleteResult struct {
	Items []models.OnenoteResourceCollectionResponse
}

// ListMeOnenoteResources ...
func (c MeOnenoteResourceClient) ListMeOnenoteResources(ctx context.Context) (result ListMeOnenoteResourcesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/me/onenote/resources",
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
		Values *[]models.OnenoteResourceCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListMeOnenoteResourcesComplete retrieves all the results into a single object
func (c MeOnenoteResourceClient) ListMeOnenoteResourcesComplete(ctx context.Context) (ListMeOnenoteResourcesCompleteResult, error) {
	return c.ListMeOnenoteResourcesCompleteMatchingPredicate(ctx, models.OnenoteResourceCollectionResponseOperationPredicate{})
}

// ListMeOnenoteResourcesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeOnenoteResourceClient) ListMeOnenoteResourcesCompleteMatchingPredicate(ctx context.Context, predicate models.OnenoteResourceCollectionResponseOperationPredicate) (result ListMeOnenoteResourcesCompleteResult, err error) {
	items := make([]models.OnenoteResourceCollectionResponse, 0)

	resp, err := c.ListMeOnenoteResources(ctx)
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

	result = ListMeOnenoteResourcesCompleteResult{
		Items: items,
	}
	return
}
