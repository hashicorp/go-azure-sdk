package workloadnetworks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSegmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkSegment
}

type ListSegmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadNetworkSegment
}

// ListSegments ...
func (c WorkloadNetworksClient) ListSegments(ctx context.Context, id PrivateCloudId) (result ListSegmentsOperationResponse, err error) {
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

// ListSegmentsComplete retrieves all the results into a single object
func (c WorkloadNetworksClient) ListSegmentsComplete(ctx context.Context, id PrivateCloudId) (ListSegmentsCompleteResult, error) {
	return c.ListSegmentsCompleteMatchingPredicate(ctx, id, WorkloadNetworkSegmentOperationPredicate{})
}

// ListSegmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkloadNetworksClient) ListSegmentsCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkSegmentOperationPredicate) (result ListSegmentsCompleteResult, err error) {
	items := make([]WorkloadNetworkSegment, 0)

	resp, err := c.ListSegments(ctx, id)
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

	result = ListSegmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
