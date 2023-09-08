package groupsitetermstoregroupsettermchildren

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

type ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrensOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreTermCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrensCompleteResult struct {
	Items []models.TermStoreTermCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrens ...
func (c GroupSiteTermStoreGroupSetTermChildrenClient) ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrens(ctx context.Context, id GroupSiteTermStoreGroupSetTermId) (result ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrensOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrensComplete retrieves all the results into a single object
func (c GroupSiteTermStoreGroupSetTermChildrenClient) ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrensComplete(ctx context.Context, id GroupSiteTermStoreGroupSetTermId) (ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrensCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrensCompleteMatchingPredicate(ctx, id, models.TermStoreTermCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrensCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreGroupSetTermChildrenClient) ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrensCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreGroupSetTermId, predicate models.TermStoreTermCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrensCompleteResult, err error) {
	items := make([]models.TermStoreTermCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrens(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreGroupByIdSetByIdTermByIdChildrensCompleteResult{
		Items: items,
	}
	return
}
