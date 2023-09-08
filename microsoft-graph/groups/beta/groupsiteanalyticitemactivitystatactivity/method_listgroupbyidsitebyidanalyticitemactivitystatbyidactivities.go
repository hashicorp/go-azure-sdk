package groupsiteanalyticitemactivitystatactivity

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

type ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivitiesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemActivityCollectionResponse
}

type ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivitiesCompleteResult struct {
	Items []models.ItemActivityCollectionResponse
}

// ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivities ...
func (c GroupSiteAnalyticItemActivityStatActivityClient) ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivities(ctx context.Context, id GroupSiteAnalyticItemActivityStatId) (result ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivitiesOperationResponse, err error) {
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
		Values *[]models.ItemActivityCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivitiesComplete retrieves all the results into a single object
func (c GroupSiteAnalyticItemActivityStatActivityClient) ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivitiesComplete(ctx context.Context, id GroupSiteAnalyticItemActivityStatId) (ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivitiesCompleteResult, error) {
	return c.ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivitiesCompleteMatchingPredicate(ctx, id, models.ItemActivityCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivitiesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteAnalyticItemActivityStatActivityClient) ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivitiesCompleteMatchingPredicate(ctx context.Context, id GroupSiteAnalyticItemActivityStatId, predicate models.ItemActivityCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivitiesCompleteResult, err error) {
	items := make([]models.ItemActivityCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivities(ctx, id)
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

	result = ListGroupByIdSiteByIdAnalyticItemActivityStatByIdActivitiesCompleteResult{
		Items: items,
	}
	return
}
