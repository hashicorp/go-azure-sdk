package triggeredjobhistories

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListTriggeredWebJobHistorySlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TriggeredJobHistory
}

type WebAppsListTriggeredWebJobHistorySlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TriggeredJobHistory
}

type WebAppsListTriggeredWebJobHistorySlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListTriggeredWebJobHistorySlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListTriggeredWebJobHistorySlot ...
func (c TriggeredJobHistoriesClient) WebAppsListTriggeredWebJobHistorySlot(ctx context.Context, id SlotTriggeredWebJobId) (result WebAppsListTriggeredWebJobHistorySlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListTriggeredWebJobHistorySlotCustomPager{},
		Path:       fmt.Sprintf("%s/history", id.ID()),
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
		Values *[]TriggeredJobHistory `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListTriggeredWebJobHistorySlotComplete retrieves all the results into a single object
func (c TriggeredJobHistoriesClient) WebAppsListTriggeredWebJobHistorySlotComplete(ctx context.Context, id SlotTriggeredWebJobId) (WebAppsListTriggeredWebJobHistorySlotCompleteResult, error) {
	return c.WebAppsListTriggeredWebJobHistorySlotCompleteMatchingPredicate(ctx, id, TriggeredJobHistoryOperationPredicate{})
}

// WebAppsListTriggeredWebJobHistorySlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TriggeredJobHistoriesClient) WebAppsListTriggeredWebJobHistorySlotCompleteMatchingPredicate(ctx context.Context, id SlotTriggeredWebJobId, predicate TriggeredJobHistoryOperationPredicate) (result WebAppsListTriggeredWebJobHistorySlotCompleteResult, err error) {
	items := make([]TriggeredJobHistory, 0)

	resp, err := c.WebAppsListTriggeredWebJobHistorySlot(ctx, id)
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

	result = WebAppsListTriggeredWebJobHistorySlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
