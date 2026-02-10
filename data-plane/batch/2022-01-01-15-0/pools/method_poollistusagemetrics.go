package pools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoolListUsageMetricsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PoolUsageMetrics
}

type PoolListUsageMetricsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PoolUsageMetrics
}

type PoolListUsageMetricsOperationOptions struct {
	ClientRequestId       *string
	Endtime               *string
	Filter                *string
	Maxresults            *int64
	OcpDate               *string
	ReturnClientRequestId *bool
	Starttime             *string
	Timeout               *int64
}

func DefaultPoolListUsageMetricsOperationOptions() PoolListUsageMetricsOperationOptions {
	return PoolListUsageMetricsOperationOptions{}
}

func (o PoolListUsageMetricsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.ClientRequestId != nil {
		out.Append("client-request-id", fmt.Sprintf("%v", *o.ClientRequestId))
	}
	if o.OcpDate != nil {
		out.Append("ocp-date", fmt.Sprintf("%v", *o.OcpDate))
	}
	if o.ReturnClientRequestId != nil {
		out.Append("return-client-request-id", fmt.Sprintf("%v", *o.ReturnClientRequestId))
	}
	return &out
}

func (o PoolListUsageMetricsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PoolListUsageMetricsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Endtime != nil {
		out.Append("endtime", fmt.Sprintf("%v", *o.Endtime))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Maxresults != nil {
		out.Append("maxresults", fmt.Sprintf("%v", *o.Maxresults))
	}
	if o.Starttime != nil {
		out.Append("starttime", fmt.Sprintf("%v", *o.Starttime))
	}
	if o.Timeout != nil {
		out.Append("timeout", fmt.Sprintf("%v", *o.Timeout))
	}
	return &out
}

type PoolListUsageMetricsCustomPager struct {
	NextLink *odata.Link `json:"odata.nextLink"`
}

func (p *PoolListUsageMetricsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PoolListUsageMetrics ...
func (c PoolsClient) PoolListUsageMetrics(ctx context.Context, options PoolListUsageMetricsOperationOptions) (result PoolListUsageMetricsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &PoolListUsageMetricsCustomPager{},
		Path:          "/poolusagemetrics",
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
		Values *[]PoolUsageMetrics `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PoolListUsageMetricsComplete retrieves all the results into a single object
func (c PoolsClient) PoolListUsageMetricsComplete(ctx context.Context, options PoolListUsageMetricsOperationOptions) (PoolListUsageMetricsCompleteResult, error) {
	return c.PoolListUsageMetricsCompleteMatchingPredicate(ctx, options, PoolUsageMetricsOperationPredicate{})
}

// PoolListUsageMetricsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PoolsClient) PoolListUsageMetricsCompleteMatchingPredicate(ctx context.Context, options PoolListUsageMetricsOperationOptions, predicate PoolUsageMetricsOperationPredicate) (result PoolListUsageMetricsCompleteResult, err error) {
	items := make([]PoolUsageMetrics, 0)

	resp, err := c.PoolListUsageMetrics(ctx, options)
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

	result = PoolListUsageMetricsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
