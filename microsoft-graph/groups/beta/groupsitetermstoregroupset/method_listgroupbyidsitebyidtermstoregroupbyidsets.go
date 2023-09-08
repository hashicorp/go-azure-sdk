package groupsitetermstoregroupset

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

type ListGroupByIdSiteByIdTermStoreGroupByIdSetsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreSetCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreGroupByIdSetsCompleteResult struct {
	Items []models.TermStoreSetCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreGroupByIdSets ...
func (c GroupSiteTermStoreGroupSetClient) ListGroupByIdSiteByIdTermStoreGroupByIdSets(ctx context.Context, id GroupSiteTermStoreGroupId) (result ListGroupByIdSiteByIdTermStoreGroupByIdSetsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/sets", id.ID()),
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

// ListGroupByIdSiteByIdTermStoreGroupByIdSetsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreGroupSetClient) ListGroupByIdSiteByIdTermStoreGroupByIdSetsComplete(ctx context.Context, id GroupSiteTermStoreGroupId) (ListGroupByIdSiteByIdTermStoreGroupByIdSetsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreGroupByIdSetsCompleteMatchingPredicate(ctx, id, models.TermStoreSetCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreGroupByIdSetsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreGroupSetClient) ListGroupByIdSiteByIdTermStoreGroupByIdSetsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreGroupId, predicate models.TermStoreSetCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreGroupByIdSetsCompleteResult, err error) {
	items := make([]models.TermStoreSetCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreGroupByIdSets(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreGroupByIdSetsCompleteResult{
		Items: items,
	}
	return
}
