package diagnostics

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobsListDetectorsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Diagnostics
}

type JobsListDetectorsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Diagnostics
}

// JobsListDetectors ...
func (c DiagnosticsClient) JobsListDetectors(ctx context.Context, id JobId) (result JobsListDetectorsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/detectors", id.ID()),
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
		Values *[]Diagnostics `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// JobsListDetectorsComplete retrieves all the results into a single object
func (c DiagnosticsClient) JobsListDetectorsComplete(ctx context.Context, id JobId) (JobsListDetectorsCompleteResult, error) {
	return c.JobsListDetectorsCompleteMatchingPredicate(ctx, id, DiagnosticsOperationPredicate{})
}

// JobsListDetectorsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DiagnosticsClient) JobsListDetectorsCompleteMatchingPredicate(ctx context.Context, id JobId, predicate DiagnosticsOperationPredicate) (result JobsListDetectorsCompleteResult, err error) {
	items := make([]Diagnostics, 0)

	resp, err := c.JobsListDetectors(ctx, id)
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

	result = JobsListDetectorsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
