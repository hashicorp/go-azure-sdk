package alerts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSubscriptionLevelAlertsByRegionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Alert
}

type ListSubscriptionLevelAlertsByRegionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Alert
}

// ListSubscriptionLevelAlertsByRegion ...
func (c AlertsClient) ListSubscriptionLevelAlertsByRegion(ctx context.Context, id LocationId) (result ListSubscriptionLevelAlertsByRegionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/alerts", id.ID()),
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
		Values *[]Alert `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSubscriptionLevelAlertsByRegionComplete retrieves all the results into a single object
func (c AlertsClient) ListSubscriptionLevelAlertsByRegionComplete(ctx context.Context, id LocationId) (ListSubscriptionLevelAlertsByRegionCompleteResult, error) {
	return c.ListSubscriptionLevelAlertsByRegionCompleteMatchingPredicate(ctx, id, AlertOperationPredicate{})
}

// ListSubscriptionLevelAlertsByRegionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AlertsClient) ListSubscriptionLevelAlertsByRegionCompleteMatchingPredicate(ctx context.Context, id LocationId, predicate AlertOperationPredicate) (result ListSubscriptionLevelAlertsByRegionCompleteResult, err error) {
	items := make([]Alert, 0)

	resp, err := c.ListSubscriptionLevelAlertsByRegion(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
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

	result = ListSubscriptionLevelAlertsByRegionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
