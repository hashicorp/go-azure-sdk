package workloadnetworksegments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworksListSegmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkSegment
}

type WorkloadNetworksListSegmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadNetworkSegment
}

// WorkloadNetworksListSegments ...
func (c WorkloadNetworkSegmentsClient) WorkloadNetworksListSegments(ctx context.Context, id PrivateCloudId) (result WorkloadNetworksListSegmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/workloadNetworks/default/segments", id.ID()),
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
		Values *[]WorkloadNetworkSegment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkloadNetworksListSegmentsComplete retrieves all the results into a single object
func (c WorkloadNetworkSegmentsClient) WorkloadNetworksListSegmentsComplete(ctx context.Context, id PrivateCloudId) (WorkloadNetworksListSegmentsCompleteResult, error) {
	return c.WorkloadNetworksListSegmentsCompleteMatchingPredicate(ctx, id, WorkloadNetworkSegmentOperationPredicate{})
}

// WorkloadNetworksListSegmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkloadNetworkSegmentsClient) WorkloadNetworksListSegmentsCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkSegmentOperationPredicate) (result WorkloadNetworksListSegmentsCompleteResult, err error) {
	items := make([]WorkloadNetworkSegment, 0)

	resp, err := c.WorkloadNetworksListSegments(ctx, id)
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

	result = WorkloadNetworksListSegmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
