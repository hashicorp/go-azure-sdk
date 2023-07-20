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

type ListVMGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkVMGroup
}

type ListVMGroupsCompleteResult struct {
	Items []WorkloadNetworkVMGroup
}

// ListVMGroups ...
func (c WorkloadNetworksClient) ListVMGroups(ctx context.Context, id PrivateCloudId) (result ListVMGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/workloadNetworks/default/vmGroups", id.ID()),
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
		Values *[]WorkloadNetworkVMGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVMGroupsComplete retrieves all the results into a single object
func (c WorkloadNetworksClient) ListVMGroupsComplete(ctx context.Context, id PrivateCloudId) (ListVMGroupsCompleteResult, error) {
	return c.ListVMGroupsCompleteMatchingPredicate(ctx, id, WorkloadNetworkVMGroupOperationPredicate{})
}

// ListVMGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c WorkloadNetworksClient) ListVMGroupsCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkVMGroupOperationPredicate) (result ListVMGroupsCompleteResult, err error) {
	items := make([]WorkloadNetworkVMGroup, 0)

	resp, err := c.ListVMGroups(ctx, id)
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

	result = ListVMGroupsCompleteResult{
		Items: items,
	}
	return
}
