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

type JobsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Job
}

type JobsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Job
}

type JobsListOperationOptions struct {
	Maxpagesize *int64
}

func DefaultJobsListOperationOptions() JobsListOperationOptions {
	return JobsListOperationOptions{}
}

func (o JobsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o JobsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o JobsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Maxpagesize != nil {
		out.Append("maxpagesize", fmt.Sprintf("%v", *o.Maxpagesize))
	}
	return &out
}

type JobsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *JobsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// JobsList ...
func (c IotcentralsClient) JobsList(ctx context.Context, options JobsListOperationOptions) (result JobsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &JobsListCustomPager{},
		Path:          "/jobs",
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
		Values *[]Job `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// JobsListComplete retrieves all the results into a single object
func (c IotcentralsClient) JobsListComplete(ctx context.Context, options JobsListOperationOptions) (JobsListCompleteResult, error) {
	return c.JobsListCompleteMatchingPredicate(ctx, options, JobOperationPredicate{})
}

// JobsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) JobsListCompleteMatchingPredicate(ctx context.Context, options JobsListOperationOptions, predicate JobOperationPredicate) (result JobsListCompleteResult, err error) {
	items := make([]Job, 0)

	resp, err := c.JobsList(ctx, options)
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

	result = JobsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
