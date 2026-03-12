package webjoboperationgroup

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

type WebAppsListWebJobsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WebJob
}

type WebAppsListWebJobsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WebJob
}

type WebAppsListWebJobsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListWebJobsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListWebJobs ...
func (c WebJobOperationGroupClient) WebAppsListWebJobs(ctx context.Context, id commonids.AppServiceId) (result WebAppsListWebJobsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WebAppsListWebJobsCustomPager{},
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

// WebAppsListWebJobsComplete retrieves all the results into a single object
func (c WebJobOperationGroupClient) WebAppsListWebJobsComplete(ctx context.Context, id commonids.AppServiceId) (WebAppsListWebJobsCompleteResult, error) {
	return c.WebAppsListWebJobsCompleteMatchingPredicate(ctx, id, WebJobOperationPredicate{})
}

// WebAppsListWebJobsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebJobOperationGroupClient) WebAppsListWebJobsCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, predicate WebJobOperationPredicate) (result WebAppsListWebJobsCompleteResult, err error) {
	items := make([]WebJob, 0)

	resp, err := c.WebAppsListWebJobs(ctx, id)
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

	result = WebAppsListWebJobsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
