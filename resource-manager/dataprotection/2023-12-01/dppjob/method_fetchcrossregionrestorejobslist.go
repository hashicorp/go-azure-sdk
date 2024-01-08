package dppjob

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FetchCrossRegionRestoreJobsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AzureBackupJobResource
}

type FetchCrossRegionRestoreJobsListCompleteResult struct {
	Items []AzureBackupJobResource
}

type FetchCrossRegionRestoreJobsListOperationOptions struct {
	Filter *string
}

func DefaultFetchCrossRegionRestoreJobsListOperationOptions() FetchCrossRegionRestoreJobsListOperationOptions {
	return FetchCrossRegionRestoreJobsListOperationOptions{}
}

func (o FetchCrossRegionRestoreJobsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o FetchCrossRegionRestoreJobsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o FetchCrossRegionRestoreJobsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// FetchCrossRegionRestoreJobsList ...
func (c DppJobClient) FetchCrossRegionRestoreJobsList(ctx context.Context, id ProviderLocationId, input CrossRegionRestoreJobsRequest, options FetchCrossRegionRestoreJobsListOperationOptions) (result FetchCrossRegionRestoreJobsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/fetchCrossRegionRestoreJobs", id.ID()),
		OptionsObject: options,
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
		Values *[]AzureBackupJobResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// FetchCrossRegionRestoreJobsListComplete retrieves all the results into a single object
func (c DppJobClient) FetchCrossRegionRestoreJobsListComplete(ctx context.Context, id ProviderLocationId, input CrossRegionRestoreJobsRequest, options FetchCrossRegionRestoreJobsListOperationOptions) (FetchCrossRegionRestoreJobsListCompleteResult, error) {
	return c.FetchCrossRegionRestoreJobsListCompleteMatchingPredicate(ctx, id, input, options, AzureBackupJobResourceOperationPredicate{})
}

// FetchCrossRegionRestoreJobsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DppJobClient) FetchCrossRegionRestoreJobsListCompleteMatchingPredicate(ctx context.Context, id ProviderLocationId, input CrossRegionRestoreJobsRequest, options FetchCrossRegionRestoreJobsListOperationOptions, predicate AzureBackupJobResourceOperationPredicate) (result FetchCrossRegionRestoreJobsListCompleteResult, err error) {
	items := make([]AzureBackupJobResource, 0)

	resp, err := c.FetchCrossRegionRestoreJobsList(ctx, id, input, options)
	if err != nil {
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

	result = FetchCrossRegionRestoreJobsListCompleteResult{
		Items: items,
	}
	return
}
