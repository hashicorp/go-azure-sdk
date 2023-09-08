package groupsitelist

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

type ListGroupByIdSiteByIdListsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ListCollectionResponse
}

type ListGroupByIdSiteByIdListsCompleteResult struct {
	Items []models.ListCollectionResponse
}

// ListGroupByIdSiteByIdLists ...
func (c GroupSiteListClient) ListGroupByIdSiteByIdLists(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdListsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/lists", id.ID()),
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
		Values *[]models.ListCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdListsComplete retrieves all the results into a single object
func (c GroupSiteListClient) ListGroupByIdSiteByIdListsComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdListsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdListsCompleteMatchingPredicate(ctx, id, models.ListCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdListsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteListClient) ListGroupByIdSiteByIdListsCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.ListCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdListsCompleteResult, err error) {
	items := make([]models.ListCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdLists(ctx, id)
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

	result = ListGroupByIdSiteByIdListsCompleteResult{
		Items: items,
	}
	return
}
