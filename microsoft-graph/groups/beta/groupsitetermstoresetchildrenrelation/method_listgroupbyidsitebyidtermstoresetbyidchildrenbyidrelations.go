package groupsitetermstoresetchildrenrelation

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

type ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreRelationCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelationsCompleteResult struct {
	Items []models.TermStoreRelationCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelations ...
func (c GroupSiteTermStoreSetChildrenRelationClient) ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelations(ctx context.Context, id GroupSiteTermStoreSetChildrenId) (result ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelationsOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelationsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreSetChildrenRelationClient) ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelationsComplete(ctx context.Context, id GroupSiteTermStoreSetChildrenId) (ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelationsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelationsCompleteMatchingPredicate(ctx, id, models.TermStoreRelationCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreSetChildrenRelationClient) ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelationsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreSetChildrenId, predicate models.TermStoreRelationCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelationsCompleteResult, err error) {
	items := make([]models.TermStoreRelationCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelations(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreSetByIdChildrenByIdRelationsCompleteResult{
		Items: items,
	}
	return
}
