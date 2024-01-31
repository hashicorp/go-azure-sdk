package workloadclassifiers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByWorkloadGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadClassifier
}

type ListByWorkloadGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadClassifier
}

// ListByWorkloadGroup ...
func (c WorkloadClassifiersClient) ListByWorkloadGroup(ctx context.Context, id WorkloadGroupId) (result ListByWorkloadGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/workloadClassifiers", id.ID()),
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
		Values *[]WorkloadClassifier `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByWorkloadGroupComplete retrieves all the results into a single object
func (c WorkloadClassifiersClient) ListByWorkloadGroupComplete(ctx context.Context, id WorkloadGroupId) (ListByWorkloadGroupCompleteResult, error) {
	return c.ListByWorkloadGroupCompleteMatchingPredicate(ctx, id, WorkloadClassifierOperationPredicate{})
}

// ListByWorkloadGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkloadClassifiersClient) ListByWorkloadGroupCompleteMatchingPredicate(ctx context.Context, id WorkloadGroupId, predicate WorkloadClassifierOperationPredicate) (result ListByWorkloadGroupCompleteResult, err error) {
	items := make([]WorkloadClassifier, 0)

	resp, err := c.ListByWorkloadGroup(ctx, id)
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

	result = ListByWorkloadGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
