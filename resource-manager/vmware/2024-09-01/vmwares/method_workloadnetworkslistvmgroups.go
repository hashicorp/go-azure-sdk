package vmwares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkloadNetworksListVMGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkVMGroup
}

type WorkloadNetworksListVMGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadNetworkVMGroup
}

type WorkloadNetworksListVMGroupsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WorkloadNetworksListVMGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WorkloadNetworksListVMGroups ...
func (c VMwaresClient) WorkloadNetworksListVMGroups(ctx context.Context, id PrivateCloudId) (result WorkloadNetworksListVMGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WorkloadNetworksListVMGroupsCustomPager{},
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

// WorkloadNetworksListVMGroupsComplete retrieves all the results into a single object
func (c VMwaresClient) WorkloadNetworksListVMGroupsComplete(ctx context.Context, id PrivateCloudId) (WorkloadNetworksListVMGroupsCompleteResult, error) {
	return c.WorkloadNetworksListVMGroupsCompleteMatchingPredicate(ctx, id, WorkloadNetworkVMGroupOperationPredicate{})
}

// WorkloadNetworksListVMGroupsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VMwaresClient) WorkloadNetworksListVMGroupsCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkVMGroupOperationPredicate) (result WorkloadNetworksListVMGroupsCompleteResult, err error) {
	items := make([]WorkloadNetworkVMGroup, 0)

	resp, err := c.WorkloadNetworksListVMGroups(ctx, id)
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

	result = WorkloadNetworksListVMGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
