package groupsitetermstoregroupset

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

type ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreSetCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetsCompleteResult struct {
	Items []models.TermStoreSetCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSets ...
func (c GroupSiteTermStoreGroupSetClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSets(ctx context.Context, id GroupSiteTermStoreGroupId) (result ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetsOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreGroupSetClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetsComplete(ctx context.Context, id GroupSiteTermStoreGroupId) (ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetsCompleteMatchingPredicate(ctx, id, models.TermStoreSetCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreGroupSetClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreGroupId, predicate models.TermStoreSetCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetsCompleteResult, err error) {
	items := make([]models.TermStoreSetCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreByIdGroupByIdSets(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetsCompleteResult{
		Items: items,
	}
	return
}
