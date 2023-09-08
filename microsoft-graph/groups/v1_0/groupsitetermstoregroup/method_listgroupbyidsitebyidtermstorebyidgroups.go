package groupsitetermstoregroup

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

type ListGroupByIdSiteByIdTermStoreByIdGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.TermStoreGroupCollectionResponse
}

type ListGroupByIdSiteByIdTermStoreByIdGroupsCompleteResult struct {
	Items []models.TermStoreGroupCollectionResponse
}

// ListGroupByIdSiteByIdTermStoreByIdGroups ...
func (c GroupSiteTermStoreGroupClient) ListGroupByIdSiteByIdTermStoreByIdGroups(ctx context.Context, id GroupSiteTermStoreId) (result ListGroupByIdSiteByIdTermStoreByIdGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/groups", id.ID()),
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

// ListGroupByIdSiteByIdTermStoreByIdGroupsComplete retrieves all the results into a single object
func (c GroupSiteTermStoreGroupClient) ListGroupByIdSiteByIdTermStoreByIdGroupsComplete(ctx context.Context, id GroupSiteTermStoreId) (ListGroupByIdSiteByIdTermStoreByIdGroupsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdTermStoreByIdGroupsCompleteMatchingPredicate(ctx, id, models.TermStoreGroupCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdTermStoreByIdGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteTermStoreGroupClient) ListGroupByIdSiteByIdTermStoreByIdGroupsCompleteMatchingPredicate(ctx context.Context, id GroupSiteTermStoreId, predicate models.TermStoreGroupCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdTermStoreByIdGroupsCompleteResult, err error) {
	items := make([]models.TermStoreGroupCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdTermStoreByIdGroups(ctx, id)
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

	result = ListGroupByIdSiteByIdTermStoreByIdGroupsCompleteResult{
		Items: items,
	}
	return
}
