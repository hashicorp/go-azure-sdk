package webjobs

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListWebJobsSlotOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WebJob
}

type WebAppsListWebJobsSlotCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WebJob
}

type WebAppsListWebJobsSlotCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListWebJobsSlotCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListWebJobsSlot ...
func (c WebJobsClient) WebAppsListWebJobsSlot(ctx context.Context, id SlotId) (result WebAppsListWebJobsSlotOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListWebJobsSlotCustomPager{},
		Path:       fmt.Sprintf("%s/webJobs", id.ID()),
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
		Values *[]WebJob `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListWebJobsSlotComplete retrieves all the results into a single object
func (c WebJobsClient) WebAppsListWebJobsSlotComplete(ctx context.Context, id SlotId) (WebAppsListWebJobsSlotCompleteResult, error) {
	return c.WebAppsListWebJobsSlotCompleteMatchingPredicate(ctx, id, WebJobOperationPredicate{})
}

// WebAppsListWebJobsSlotCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebJobsClient) WebAppsListWebJobsSlotCompleteMatchingPredicate(ctx context.Context, id SlotId, predicate WebJobOperationPredicate) (result WebAppsListWebJobsSlotCompleteResult, err error) {
	items := make([]WebJob, 0)

	resp, err := c.WebAppsListWebJobsSlot(ctx, id)
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

	result = WebAppsListWebJobsSlotCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
