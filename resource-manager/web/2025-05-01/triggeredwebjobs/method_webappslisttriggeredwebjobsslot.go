package triggeredwebjobs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListTriggeredWebJobsSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TriggeredWebJob
}

type WebAppsListTriggeredWebJobsSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TriggeredWebJob
}

type WebAppsListTriggeredWebJobsSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListTriggeredWebJobsSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListTriggeredWebJobsSlot ...
func (c TriggeredWebJobsClient) WebAppsListTriggeredWebJobsSlot(ctx context.Context, id SlotId) (result WebAppsListTriggeredWebJobsSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListTriggeredWebJobsSlotCustomPager{},
		Path:       fmt.Sprintf("%s/triggeredWebJobs", id.ID()),
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
		Values *[]TriggeredWebJob `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListTriggeredWebJobsSlotComplete retrieves all the results into a single object
func (c TriggeredWebJobsClient) WebAppsListTriggeredWebJobsSlotComplete(ctx context.Context, id SlotId) (WebAppsListTriggeredWebJobsSlotCompleteResult, error) {
	return c.WebAppsListTriggeredWebJobsSlotCompleteMatchingPredicate(ctx, id, TriggeredWebJobOperationPredicate{})
}

// WebAppsListTriggeredWebJobsSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TriggeredWebJobsClient) WebAppsListTriggeredWebJobsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate TriggeredWebJobOperationPredicate) (result WebAppsListTriggeredWebJobsSlotCompleteResult, err error) {
	items := make([]TriggeredWebJob, 0)

	resp, err := c.WebAppsListTriggeredWebJobsSlot(ctx, id)
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

	result = WebAppsListTriggeredWebJobsSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
