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

type ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrensOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreTermCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrensCompleteResult struct {
	Items []models.TermStoreTermCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrens ...
func (c GroupSiteTermStoreGroupSetTermChildrenClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrens(ctx context.Context, id GroupSiteTermStoreGroupSetTermId) (result ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrensOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrensComplete retrieves all the results into a single object
func (c GroupSiteTermStoreGroupSetTermChildrenClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrensComplete(ctx context.Context, id GroupSiteTermStoreGroupSetTermId) (ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrensCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrensCompleteMatchingPredicate(ctx, id, models.TermStoreTermCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrensCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreGroupSetTermChildrenClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrensCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreGroupSetTermId, predicate models.TermStoreTermCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrensCompleteResult, err error) {
	items := make([]models.TermStoreTermCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrens(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdTermByIdChildrensCompleteResult{
		Items: items,
	}
	return
}
