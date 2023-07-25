package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListContinuousWebJobsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ContinuousWebJob
}

type ListContinuousWebJobsCompleteResult struct {
	Items []ContinuousWebJob
}

// ListContinuousWebJobs ...
func (c WebAppsClient) ListContinuousWebJobs(ctx context.Context, id SiteId) (result ListContinuousWebJobsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/continuousWebJobs", id.ID()),
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
		Values *[]ContinuousWebJob `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListContinuousWebJobsComplete retrieves all the results into a single object
func (c WebAppsClient) ListContinuousWebJobsComplete(ctx context.Context, id SiteId) (ListContinuousWebJobsCompleteResult, error) {
	return c.ListContinuousWebJobsCompleteMatchingPredicate(ctx, id, ContinuousWebJobOperationPredicate{})
}

// ListContinuousWebJobsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WebAppsClient) ListContinuousWebJobsCompleteMatchingPredicate(ctx context.Context, id SiteId, predicate ContinuousWebJobOperationPredicate) (result ListContinuousWebJobsCompleteResult, err error) {
	items := make([]ContinuousWebJob, 0)

	resp, err := c.ListContinuousWebJobs(ctx, id)
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

	result = ListContinuousWebJobsCompleteResult{
		Items: items,
	}
	return
}
