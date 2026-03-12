package triggeredwebjoboperationgroup

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsListTriggeredWebJobsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]TriggeredWebJob
}

type WebAppsListTriggeredWebJobsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []TriggeredWebJob
}

type WebAppsListTriggeredWebJobsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListTriggeredWebJobsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListTriggeredWebJobs ...
func (c TriggeredWebJobOperationGroupClient) WebAppsListTriggeredWebJobs(ctx context.Context, id commonids.AppServiceId) (result WebAppsListTriggeredWebJobsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListTriggeredWebJobsCustomPager{},
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

// WebAppsListTriggeredWebJobsComplete retrieves all the results into a single object
func (c TriggeredWebJobOperationGroupClient) WebAppsListTriggeredWebJobsComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListTriggeredWebJobsCompleteResult, error) {
	return c.WebAppsListTriggeredWebJobsCompleteMatchingPredicate(ctx, id, TriggeredWebJobOperationPredicate{})
}

// WebAppsListTriggeredWebJobsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TriggeredWebJobOperationGroupClient) WebAppsListTriggeredWebJobsCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate TriggeredWebJobOperationPredicate) (result WebAppsListTriggeredWebJobsCompleteResult, err error) {
	items := make([]TriggeredWebJob, 0)

	resp, err := c.WebAppsListTriggeredWebJobs(ctx, id)
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

	result = WebAppsListTriggeredWebJobsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
