package groupsitetermstoregroupsetchildren

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

type ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrensOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreTermCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrensCompleteResult struct {
	Items []models.TermStoreTermCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrens ...
func (c GroupSiteTermStoreGroupSetChildrenClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrens(ctx context.Context, id GroupSiteTermStoreGroupSetId) (result ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrensOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrensComplete retrieves all the results into a single object
func (c GroupSiteTermStoreGroupSetChildrenClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrensComplete(ctx context.Context, id GroupSiteTermStoreGroupSetId) (ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrensCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrensCompleteMatchingPredicate(ctx, id, models.TermStoreTermCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrensCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreGroupSetChildrenClient) ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrensCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreGroupSetId, predicate models.TermStoreTermCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrensCompleteResult, err error) {
	items := make([]models.TermStoreTermCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrens(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreByIdGroupByIdSetByIdChildrensCompleteResult{
		Items: items,
	}
	return
}
