package groupsitetermstoresetparentgroupsettermchildren

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

type ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrensOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreTermCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrensCompleteResult struct {
	Items []models.TermStoreTermCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrens ...
func (c GroupSiteTermStoreSetParentGroupSetTermChildrenClient) ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrens(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetTermId) (result ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrensOperationResponse, err error) {
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

// ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrensComplete retrieves all the results into a single object
func (c GroupSiteTermStoreSetParentGroupSetTermChildrenClient) ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrensComplete(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetTermId) (ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrensCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrensCompleteMatchingPredicate(ctx, id, models.TermStoreTermCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrensCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreSetParentGroupSetTermChildrenClient) ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrensCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreSetParentGroupSetTermId, predicate models.TermStoreTermCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrensCompleteResult, err error) {
	items := make([]models.TermStoreTermCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrens(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreSetByIdParentGroupSetByIdTermByIdChildrensCompleteResult{
		Items: items,
	}
	return
}
