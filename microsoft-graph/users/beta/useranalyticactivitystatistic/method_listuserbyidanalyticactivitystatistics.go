package useranalyticactivitystatistic

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

type ListUserByIdAnalyticActivityStatisticsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.ActivityStatisticsCollectionResponse
}

type ListUserByIdAnalyticActivityStatisticsCompleteResult struct {
	Items []models.ActivityStatisticsCollectionResponse
}

// ListUserByIdAnalyticActivityStatistics ...
func (c UserAnalyticActivityStatisticClient) ListUserByIdAnalyticActivityStatistics(ctx context.Context, id UserId) (result ListUserByIdAnalyticActivityStatisticsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/analytics/activityStatistics", id.ID()),
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
		Values *[]models.ActivityStatisticsCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListUserByIdAnalyticActivityStatisticsComplete retrieves all the results into a single object
func (c UserAnalyticActivityStatisticClient) ListUserByIdAnalyticActivityStatisticsComplete(ctx context.Context, id UserId) (ListUserByIdAnalyticActivityStatisticsCompleteResult, error) {
	return c.ListUserByIdAnalyticActivityStatisticsCompleteMatchingPredicate(ctx, id, models.ActivityStatisticsCollectionResponseOperationPredicate{})
}

// ListUserByIdAnalyticActivityStatisticsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserAnalyticActivityStatisticClient) ListUserByIdAnalyticActivityStatisticsCompleteMatchingPredicate(ctx context.Context, id UserId, predicate models.ActivityStatisticsCollectionResponseOperationPredicate) (result ListUserByIdAnalyticActivityStatisticsCompleteResult, err error) {
	items := make([]models.ActivityStatisticsCollectionResponse, 0)

	resp, err := c.ListUserByIdAnalyticActivityStatistics(ctx, id)
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

	result = ListUserByIdAnalyticActivityStatisticsCompleteResult{
		Items: items,
	}
	return
}
