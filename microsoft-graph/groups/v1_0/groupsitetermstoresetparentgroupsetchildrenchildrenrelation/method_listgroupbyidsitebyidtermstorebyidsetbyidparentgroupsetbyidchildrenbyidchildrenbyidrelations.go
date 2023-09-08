package groupsitetermstoresetparentgroupsetchildrenchildrenrelation

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

type ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreRelationCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelationsCompleteResult struct {
	Items []models.TermStoreRelationCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelations ...
func (c GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelations(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetChildrenChildrenId) (result ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelationsOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelationsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelationsComplete(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetChildrenChildrenId) (ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelationsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelationsCompleteMatchingPredicate(ctx, id, models.TermStoreRelationCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreSetParentGroupSetChildrenChildrenRelationClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelationsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetChildrenChildrenId, predicate models.TermStoreRelationCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelationsCompleteResult, err error) {
	items := make([]models.TermStoreRelationCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelations(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreByIdSetByIdParentGroupSetByIdChildrenByIdChildrenByIdRelationsCompleteResult{
		Items: items,
	}
	return
}
