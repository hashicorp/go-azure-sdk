package groupsitetermstoresettermchildrenrelation

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

type ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreRelationCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelationsCompleteResult struct {
	Items []models.TermStoreRelationCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelations ...
func (c GroupSiteTermStoreSetTermChildrenRelationClient) ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelations(ctx context.Context, id GroupSiteTermStoreSetTermChildrenId) (result ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelationsOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelationsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreSetTermChildrenRelationClient) ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelationsComplete(ctx context.Context, id GroupSiteTermStoreSetTermChildrenId) (ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelationsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelationsCompleteMatchingPredicate(ctx, id, models.TermStoreRelationCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreSetTermChildrenRelationClient) ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelationsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreSetTermChildrenId, predicate models.TermStoreRelationCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelationsCompleteResult, err error) {
	items := make([]models.TermStoreRelationCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelations(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreSetByIdTermByIdChildrenByIdRelationsCompleteResult{
		Items: items,
	}
	return
}
