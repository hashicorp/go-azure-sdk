package groupsitetermstoregroupsetchildren

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

type ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrensOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreTermCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrensCompleteResult struct {
	Items []models.TermStoreTermCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrens ...
func (c GroupSiteTermStoreGroupSetChildrenClient) ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrens(ctx context.Context, id GroupSiteTermStoreGroupSetId) (result ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrensOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrensComplete retrieves all the results into a single object
func (c GroupSiteTermStoreGroupSetChildrenClient) ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrensComplete(ctx context.Context, id GroupSiteTermStoreGroupSetId) (ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrensCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrensCompleteMatchingPredicate(ctx, id, models.TermStoreTermCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrensCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreGroupSetChildrenClient) ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrensCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreGroupSetId, predicate models.TermStoreTermCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrensCompleteResult, err error) {
	items := make([]models.TermStoreTermCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrens(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdChildrensCompleteResult{
		Items: items,
	}
	return
}
