package networkclouds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricsConfigurationsListByClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ClusterMetricsConfiguration
}

type MetricsConfigurationsListByClusterCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ClusterMetricsConfiguration
}

type MetricsConfigurationsListByClusterOperationOptions struct {
	Top *int64
}

func DefaultMetricsConfigurationsListByClusterOperationOptions() MetricsConfigurationsListByClusterOperationOptions {
	return MetricsConfigurationsListByClusterOperationOptions{}
}

func (o MetricsConfigurationsListByClusterOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o MetricsConfigurationsListByClusterOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o MetricsConfigurationsListByClusterOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type MetricsConfigurationsListByClusterCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *MetricsConfigurationsListByClusterCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// MetricsConfigurationsListByCluster ...
func (c NetworkcloudsClient) MetricsConfigurationsListByCluster(ctx context.Context, id ClusterId, options MetricsConfigurationsListByClusterOperationOptions) (result MetricsConfigurationsListByClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &MetricsConfigurationsListByClusterCustomPager{},
		Path:          fmt.Sprintf("%s/metricsConfigurations", id.ID()),
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
		Values *[]ClusterMetricsConfiguration `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// MetricsConfigurationsListByClusterComplete retrieves all the results into a single object
func (c NetworkcloudsClient) MetricsConfigurationsListByClusterComplete(ctx context.Context, id ClusterId, options MetricsConfigurationsListByClusterOperationOptions) (MetricsConfigurationsListByClusterCompleteResult, error) {
	return c.MetricsConfigurationsListByClusterCompleteMatchingPredicate(ctx, id, options, ClusterMetricsConfigurationOperationPredicate{})
}

// MetricsConfigurationsListByClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) MetricsConfigurationsListByClusterCompleteMatchingPredicate(ctx context.Context, id ClusterId, options MetricsConfigurationsListByClusterOperationOptions, predicate ClusterMetricsConfigurationOperationPredicate) (result MetricsConfigurationsListByClusterCompleteResult, err error) {
	items := make([]ClusterMetricsConfiguration, 0)

	resp, err := c.MetricsConfigurationsListByCluster(ctx, id, options)
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

	result = MetricsConfigurationsListByClusterCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
