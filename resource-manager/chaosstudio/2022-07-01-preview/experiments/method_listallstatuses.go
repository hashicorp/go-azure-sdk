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

type ListAllStatusesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]ExperimentStatus
}

type ListAllStatusesCompleteResult struct {
	Items []ExperimentStatus
}

// ListAllStatuses ...
func (c ExperimentsClient) ListAllStatuses(ctx context.Context, id ExperimentId) (result ListAllStatusesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/statuses", id.ID()),
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
		Values *[]ExperimentStatus `json:"values"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAllStatusesComplete retrieves all the results into a single object
func (c ExperimentsClient) ListAllStatusesComplete(ctx context.Context, id ExperimentId) (ListAllStatusesCompleteResult, error) {
	return c.ListAllStatusesCompleteMatchingPredicate(ctx, id, ExperimentStatusOperationPredicate{})
}

// ListAllStatusesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExperimentsClient) ListAllStatusesCompleteMatchingPredicate(ctx context.Context, id ExperimentId, predicate ExperimentStatusOperationPredicate) (result ListAllStatusesCompleteResult, err error) {
	items := make([]ExperimentStatus, 0)

	resp, err := c.ListAllStatuses(ctx, id)
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

	result = ListAllStatusesCompleteResult{
		Items: items,
	}
	return
}
