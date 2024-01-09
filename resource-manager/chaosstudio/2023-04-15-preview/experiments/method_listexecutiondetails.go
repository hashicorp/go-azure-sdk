package experiments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListExecutionDetailsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ExperimentExecutionDetails
}

type ListExecutionDetailsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []ExperimentExecutionDetails
}

// ListExecutionDetails ...
func (c ExperimentsClient) ListExecutionDetails(ctx context.Context, id ExperimentId) (result ListExecutionDetailsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/executionDetails", id.ID()),
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
		Values *[]ExperimentExecutionDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExecutionDetailsComplete retrieves all the results into a single object
func (c ExperimentsClient) ListExecutionDetailsComplete(ctx context.Context, id ExperimentId) (ListExecutionDetailsCompleteResult, error) {
	return c.ListExecutionDetailsCompleteMatchingPredicate(ctx, id, ExperimentExecutionDetailsOperationPredicate{})
}

// ListExecutionDetailsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExperimentsClient) ListExecutionDetailsCompleteMatchingPredicate(ctx context.Context, id ExperimentId, predicate ExperimentExecutionDetailsOperationPredicate) (result ListExecutionDetailsCompleteResult, err error) {
	items := make([]ExperimentExecutionDetails, 0)

	resp, err := c.ListExecutionDetails(ctx, id)
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

	result = ListExecutionDetailsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
