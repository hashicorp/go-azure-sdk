package sites

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

type WebAppsListPerfMonCountersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PerfMonResponse
}

type WebAppsListPerfMonCountersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PerfMonResponse
}

type WebAppsListPerfMonCountersOperationOptions struct {
	Filter *string
}

func DefaultWebAppsListPerfMonCountersOperationOptions() WebAppsListPerfMonCountersOperationOptions {
	return WebAppsListPerfMonCountersOperationOptions{}
}

func (o WebAppsListPerfMonCountersOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WebAppsListPerfMonCountersOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o WebAppsListPerfMonCountersOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

type WebAppsListPerfMonCountersCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WebAppsListPerfMonCountersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WebAppsListPerfMonCounters ...
func (c SitesClient) WebAppsListPerfMonCounters(ctx context.Context, id commonids.AppServiceId, options WebAppsListPerfMonCountersOperationOptions) (result WebAppsListPerfMonCountersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &WebAppsListPerfMonCountersCustomPager{},
		Path:          fmt.Sprintf("%s/perfcounters", id.ID()),
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
		Values *[]PerfMonResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WebAppsListPerfMonCountersComplete retrieves all the results into a single object
func (c SitesClient) WebAppsListPerfMonCountersComplete(ctx context.Context, id commonids.AppServiceId, options WebAppsListPerfMonCountersOperationOptions) (WebAppsListPerfMonCountersCompleteResult, error) {
	return c.WebAppsListPerfMonCountersCompleteMatchingPredicate(ctx, id, options, PerfMonResponseOperationPredicate{})
}

// WebAppsListPerfMonCountersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c SitesClient) WebAppsListPerfMonCountersCompleteMatchingPredicate(ctx context.Context, id commonids.AppServiceId, options WebAppsListPerfMonCountersOperationOptions, predicate PerfMonResponseOperationPredicate) (result WebAppsListPerfMonCountersCompleteResult, err error) {
	items := make([]PerfMonResponse, 0)

	resp, err := c.WebAppsListPerfMonCounters(ctx, id, options)
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

	result = WebAppsListPerfMonCountersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
