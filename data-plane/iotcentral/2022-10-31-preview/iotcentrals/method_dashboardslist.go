package iotcentrals

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DashboardsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Dashboard
}

type DashboardsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Dashboard
}

type DashboardsListOperationOptions struct {
	Filter      *string
	Maxpagesize *int64
	Orderby     *string
}

func DefaultDashboardsListOperationOptions() DashboardsListOperationOptions {
	return DashboardsListOperationOptions{}
}

func (o DashboardsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DashboardsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DashboardsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Maxpagesize != nil {
		out.Append("maxpagesize", fmt.Sprintf("%v", *o.Maxpagesize))
	}
	if o.Orderby != nil {
		out.Append("orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	return &out
}

type DashboardsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DashboardsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DashboardsList ...
func (c IotcentralsClient) DashboardsList(ctx context.Context, options DashboardsListOperationOptions) (result DashboardsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &DashboardsListCustomPager{},
		Path:          "/dashboards",
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
		Values *[]Dashboard `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DashboardsListComplete retrieves all the results into a single object
func (c IotcentralsClient) DashboardsListComplete(ctx context.Context, options DashboardsListOperationOptions) (DashboardsListCompleteResult, error) {
	return c.DashboardsListCompleteMatchingPredicate(ctx, options, DashboardOperationPredicate{})
}

// DashboardsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DashboardsListCompleteMatchingPredicate(ctx context.Context, options DashboardsListOperationOptions, predicate DashboardOperationPredicate) (result DashboardsListCompleteResult, err error) {
	items := make([]Dashboard, 0)

	resp, err := c.DashboardsList(ctx, options)
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

	result = DashboardsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
