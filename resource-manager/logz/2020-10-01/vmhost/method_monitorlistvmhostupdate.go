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

type MonitorListVMHostUpdateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]VMResources
}

type MonitorListVMHostUpdateCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []VMResources
}

// MonitorListVMHostUpdate ...
func (c VMHostClient) MonitorListVMHostUpdate(ctx context.Context, id MonitorId, input VMHostUpdateRequest) (result MonitorListVMHostUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/vmHostUpdate", id.ID()),
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
		Values *[]VMResources `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// MonitorListVMHostUpdateComplete retrieves all the results into a single object
func (c VMHostClient) MonitorListVMHostUpdateComplete(ctx context.Context, id MonitorId, input VMHostUpdateRequest) (MonitorListVMHostUpdateCompleteResult, error) {
	return c.MonitorListVMHostUpdateCompleteMatchingPredicate(ctx, id, input, VMResourcesOperationPredicate{})
}

// MonitorListVMHostUpdateCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VMHostClient) MonitorListVMHostUpdateCompleteMatchingPredicate(ctx context.Context, id MonitorId, input VMHostUpdateRequest, predicate VMResourcesOperationPredicate) (result MonitorListVMHostUpdateCompleteResult, err error) {
	items := make([]VMResources, 0)

	resp, err := c.MonitorListVMHostUpdate(ctx, id, input)
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

	result = MonitorListVMHostUpdateCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
