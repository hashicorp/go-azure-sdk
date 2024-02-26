package vmhost

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorListVMHostsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VMResourcesListResponse
}

type MonitorListVMHostsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []VMResourcesListResponse
}

// MonitorListVMHosts ...
func (c VMHostClient) MonitorListVMHosts(ctx context.Context, id MonitorId) (result MonitorListVMHostsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/listVMHosts", id.ID()),
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
		Values *[]VMResourcesListResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// MonitorListVMHostsComplete retrieves all the results into a single object
func (c VMHostClient) MonitorListVMHostsComplete(ctx context.Context, id MonitorId) (MonitorListVMHostsCompleteResult, error) {
	return c.MonitorListVMHostsCompleteMatchingPredicate(ctx, id, VMResourcesListResponseOperationPredicate{})
}

// MonitorListVMHostsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VMHostClient) MonitorListVMHostsCompleteMatchingPredicate(ctx context.Context, id MonitorId, predicate VMResourcesListResponseOperationPredicate) (result MonitorListVMHostsCompleteResult, err error) {
	items := make([]VMResourcesListResponse, 0)

	resp, err := c.MonitorListVMHosts(ctx, id)
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

	result = MonitorListVMHostsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
