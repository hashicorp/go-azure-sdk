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

type ScheduledJobsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ScheduledJob
}

type ScheduledJobsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ScheduledJob
}

type ScheduledJobsListOperationOptions struct {
	Maxpagesize *int64
}

func DefaultScheduledJobsListOperationOptions() ScheduledJobsListOperationOptions {
	return ScheduledJobsListOperationOptions{}
}

func (o ScheduledJobsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ScheduledJobsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ScheduledJobsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Maxpagesize != nil {
		out.Append("maxpagesize", fmt.Sprintf("%v", *o.Maxpagesize))
	}
	return &out
}

type ScheduledJobsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ScheduledJobsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ScheduledJobsList ...
func (c IotcentralsClient) ScheduledJobsList(ctx context.Context, options ScheduledJobsListOperationOptions) (result ScheduledJobsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ScheduledJobsListCustomPager{},
		Path:          "/scheduledJobs",
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
		Values *[]ScheduledJob `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ScheduledJobsListComplete retrieves all the results into a single object
func (c IotcentralsClient) ScheduledJobsListComplete(ctx context.Context, options ScheduledJobsListOperationOptions) (ScheduledJobsListCompleteResult, error) {
	return c.ScheduledJobsListCompleteMatchingPredicate(ctx, options, ScheduledJobOperationPredicate{})
}

// ScheduledJobsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) ScheduledJobsListCompleteMatchingPredicate(ctx context.Context, options ScheduledJobsListOperationOptions, predicate ScheduledJobOperationPredicate) (result ScheduledJobsListCompleteResult, err error) {
	items := make([]ScheduledJob, 0)

	resp, err := c.ScheduledJobsList(ctx, options)
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

	result = ScheduledJobsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
