package groupsitetermstoresetchildrenchildren

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

type ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrensOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreTermCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrensCompleteResult struct {
	Items []models.TermStoreTermCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrens ...
func (c GroupSiteTermStoreSetChildrenChildrenClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrens(ctx context.Context, id GroupSiteTermStoreSetChildrenId) (result ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrensOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/children", id.ID()),
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
		Values *[]models.TermStoreTermCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrensComplete retrieves all the results into a single object
func (c GroupSiteTermStoreSetChildrenChildrenClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrensComplete(ctx context.Context, id GroupSiteTermStoreSetChildrenId) (ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrensCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrensCompleteMatchingPredicate(ctx, id, models.TermStoreTermCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrensCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreSetChildrenChildrenClient) ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrensCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreSetChildrenId, predicate models.TermStoreTermCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrensCompleteResult, err error) {
	items := make([]models.TermStoreTermCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrens(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreByIdSetByIdChildrenByIdChildrensCompleteResult{
		Items: items,
	}
	return
}
