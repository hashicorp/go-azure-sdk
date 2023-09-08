package groupsitetermstoresetparentgroupset

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

type ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreSetCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetsCompleteResult struct {
	Items []models.TermStoreSetCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSets ...
func (c GroupSiteTermStoreSetParentGroupSetClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSets(ctx context.Context, id GroupSiteTermStoreSetId) (result ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/parentGroup/sets", id.ID()),
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

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreSetParentGroupSetClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetsComplete(ctx context.Context, id GroupSiteTermStoreSetId) (ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetsCompleteMatchingPredicate(ctx, id, models.TermStoreSetCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreSetParentGroupSetClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreSetId, predicate models.TermStoreSetCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetsCompleteResult, err error) {
	items := make([]models.TermStoreSetCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSets(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetsCompleteResult{
		Items: items,
	}
	return
}
