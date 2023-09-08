package groupsiteanalyticitemactivitystat

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

type ListGroupByIdSiteByIdAnalyticItemActivityStatsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ItemActivityStatCollectionResponse
}

type ListGroupByIdSiteByIdAnalyticItemActivityStatsCompleteResult struct {
	Items []models.ItemActivityStatCollectionResponse
}

// ListGroupByIdSiteByIdAnalyticItemActivityStats ...
func (c GroupSiteAnalyticItemActivityStatClient) ListGroupByIdSiteByIdAnalyticItemActivityStats(ctx context.Context, id GroupSiteId) (result ListGroupByIdSiteByIdAnalyticItemActivityStatsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/analytics/itemActivityStats", id.ID()),
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
		Values *[]models.ItemActivityStatCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdSiteByIdAnalyticItemActivityStatsComplete retrieves all the results into a single object
func (c GroupSiteAnalyticItemActivityStatClient) ListGroupByIdSiteByIdAnalyticItemActivityStatsComplete(ctx context.Context, id GroupSiteId) (ListGroupByIdSiteByIdAnalyticItemActivityStatsCompleteResult, error) {
	return c.ListGroupByIdSiteByIdAnalyticItemActivityStatsCompleteMatchingPredicate(ctx, id, models.ItemActivityStatCollectionResponseOperationPredicate{})
}

// ListGroupByIdSiteByIdAnalyticItemActivityStatsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupSiteAnalyticItemActivityStatClient) ListGroupByIdSiteByIdAnalyticItemActivityStatsCompleteMatchingPredicate(ctx context.Context, id GroupSiteId, predicate models.ItemActivityStatCollectionResponseOperationPredicate) (result ListGroupByIdSiteByIdAnalyticItemActivityStatsCompleteResult, err error) {
	items := make([]models.ItemActivityStatCollectionResponse, 0)

	resp, err := c.ListGroupByIdSiteByIdAnalyticItemActivityStats(ctx, id)
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

	result = ListGroupByIdSiteByIdAnalyticItemActivityStatsCompleteResult{
		Items: items,
	}
	return
}
