package groupsitetermstoresetparentgroupsetrelation

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

type ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreRelationCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelationsCompleteResult struct {
	Items []models.TermStoreRelationCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelations ...
func (c GroupSiteTermStoreSetParentGroupSetRelationClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelations(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetId) (result ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/relations", id.ID()),
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
		Values *[]models.TermStoreRelationCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelationsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreSetParentGroupSetRelationClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelationsComplete(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetId) (ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelationsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelationsCompleteMatchingPredicate(ctx, id, models.TermStoreRelationCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreSetParentGroupSetRelationClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelationsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetId, predicate models.TermStoreRelationCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelationsCompleteResult, err error) {
	items := make([]models.TermStoreRelationCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelations(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdRelationsCompleteResult{
		Items: items,
	}
	return
}
