package groupsitelistitemversion

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

type ListGroupByIdSiteByIdListByIdItemByIdVersionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ListItemVersionCollectionResponse
}

type ListGroupByIdSiteByIdListByIdItemByIdVersionsCompleteResult struct {
	Items []models.ListItemVersionCollectionResponse
}

// ListGroupByIdSiteByIdListByIdItemByIdVersions ...
func (c GroupSiteListItemVersionClient) ListGroupByIdSiteByIdListByIdItemByIdVersions(ctx context.Context, id GroupSiteListItemId) (result ListGroupByIdSiteByIdListByIdItemByIdVersionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]models.ListItemVersionCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdListByIdItemByIdVersionsComplete retrieves all the results into a single object
func (c GroupSiteListItemVersionClient) ListGroupByIdSiteByIdListByIdItemByIdVersionsComplete(ctx context.Context, id GroupSiteListItemId) (ListGroupByIdSiteByIdListByIdItemByIdVersionsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdListByIdItemByIdVersionsCompleteMatchingPredicate(ctx, id, models.ListItemVersionCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdListByIdItemByIdVersionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteListItemVersionClient) ListGroupByIdSiteByIdListByIdItemByIdVersionsCompleteMatchingPredicate(ctx context.Context, id GroupSiteListItemId, predicate models.ListItemVersionCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdListByIdItemByIdVersionsCompleteResult, err error) {
	items := make([]models.ListItemVersionCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdListByIdItemByIdVersions(ctx, id)
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

	result = ListGroupByIdSiteByIdListByIdItemByIdVersionsCompleteResult{
		Items: items,
	}
	return
}
