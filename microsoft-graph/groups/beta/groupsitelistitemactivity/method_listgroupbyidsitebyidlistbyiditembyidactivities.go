package groupsitelistitemactivity

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

type ListGroupByIdSiteByIdListByIdItemByIdActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemActivityOLDCollectionResponse
}

type ListGroupByIdSiteByIdListByIdItemByIdActivitiesCompleteResult struct {
	Items []models.ItemActivityOLDCollectionResponse
}

// ListGroupByIdSiteByIdListByIdItemByIdActivities ...
func (c GroupSiteListItemActivityClient) ListGroupByIdSiteByIdListByIdItemByIdActivities(ctx context.Context, id GroupSiteListItemId) (result ListGroupByIdSiteByIdListByIdItemByIdActivitiesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/activities", id.ID()),
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
		Values *[]models.ItemActivityOLDCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdListByIdItemByIdActivitiesComplete retrieves all the results into a single object
func (c GroupSiteListItemActivityClient) ListGroupByIdSiteByIdListByIdItemByIdActivitiesComplete(ctx context.Context, id GroupSiteListItemId) (ListGroupByIdSiteByIdListByIdItemByIdActivitiesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdListByIdItemByIdActivitiesCompleteMatchingPredicate(ctx, id, models.ItemActivityOLDCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdListByIdItemByIdActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteListItemActivityClient) ListGroupByIdSiteByIdListByIdItemByIdActivitiesCompleteMatchingPredicate(ctx context.Context, id GroupSiteListItemId, predicate models.ItemActivityOLDCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdListByIdItemByIdActivitiesCompleteResult, err error) {
	items := make([]models.ItemActivityOLDCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdListByIdItemByIdActivities(ctx, id)
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

	result = ListGroupByIdSiteByIdListByIdItemByIdActivitiesCompleteResult{
		Items: items,
	}
	return
}
