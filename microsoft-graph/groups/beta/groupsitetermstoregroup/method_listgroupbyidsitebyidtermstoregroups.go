package groupsitetermstoregroup

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

type ListGroupByIdSiteByIdTermStoreGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreGroupCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreGroupsCompleteResult struct {
	Items []models.TermStoreGroupCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreGroups ...
func (c GroupSiteTermStoreGroupClient) ListGroupByIdSiteByIdTermStoreGroups(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdTermStoreGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/termStore/groups", id.ID()),
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
		Values *[]models.TermStoreGroupCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdTermStoreGroupsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreGroupClient) ListGroupByIdSiteByIdTermStoreGroupsComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdTermStoreGroupsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreGroupsCompleteMatchingPredicate(ctx, id, models.TermStoreGroupCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreGroupClient) ListGroupByIdSiteByIdTermStoreGroupsCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.TermStoreGroupCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreGroupsCompleteResult, err error) {
	items := make([]models.TermStoreGroupCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreGroups(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreGroupsCompleteResult{
		Items: items,
	}
	return
}
