package groupsitetermstoresetparentgroupsettermchildrenrelation

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

type ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreRelationCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelationsCompleteResult struct {
	Items []models.TermStoreRelationCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelations ...
func (c GroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient) ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelations(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetTermChildrenId) (result ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelationsOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelationsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient) ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelationsComplete(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetTermChildrenId) (ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelationsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelationsCompleteMatchingPredicate(ctx, id, models.TermStoreRelationCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreSetParentGroupSetTermChildrenRelationClient) ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelationsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetTermChildrenId, predicate models.TermStoreRelationCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelationsCompleteResult, err error) {
	items := make([]models.TermStoreRelationCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelations(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrenByIdRelationsCompleteResult{
		Items: items,
	}
	return
}
