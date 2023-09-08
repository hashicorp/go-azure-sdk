package groupsitetermstore

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

type ListGroupByIdSiteByIdTermStoresOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreStoreCollectionResponse
}

type ListGroupByIdSiteByIdTermStoresCompleteResult struct {
	Items []models.TermStoreStoreCollectionResponse
}

// ListGroupByIdSiteByIdTermStores ...
func (c GroupSiteTermStoreClient) ListGroupByIdSiteByIdTermStores(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdTermStoresOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/termStores", id.ID()),
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
		Values *[]models.TermStoreStoreCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdTermStoresComplete retrieves all the results into a single object
func (c GroupSiteTermStoreClient) ListGroupByIdSiteByIdTermStoresComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdTermStoresCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoresCompleteMatchingPredicate(ctx, id, models.TermStoreStoreCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoresCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreClient) ListGroupByIdSiteByIdTermStoresCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.TermStoreStoreCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoresCompleteResult, err error) {
	items := make([]models.TermStoreStoreCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStores(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoresCompleteResult{
		Items: items,
	}
	return
}
