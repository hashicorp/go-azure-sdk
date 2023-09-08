package groupsitetermstoreset

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

type ListGroupByIdSiteByIdTermStoreSetsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreSetCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreSetsCompleteResult struct {
	Items []models.TermStoreSetCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreSets ...
func (c GroupSiteTermStoreSetClient) ListGroupByIdSiteByIdTermStoreSets(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdTermStoreSetsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/termStore/sets", id.ID()),
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
		Values *[]models.TermStoreSetCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdTermStoreSetsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreSetClient) ListGroupByIdSiteByIdTermStoreSetsComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdTermStoreSetsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreSetsCompleteMatchingPredicate(ctx, id, models.TermStoreSetCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreSetsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreSetClient) ListGroupByIdSiteByIdTermStoreSetsCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.TermStoreSetCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreSetsCompleteResult, err error) {
	items := make([]models.TermStoreSetCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreSets(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreSetsCompleteResult{
		Items: items,
	}
	return
}
