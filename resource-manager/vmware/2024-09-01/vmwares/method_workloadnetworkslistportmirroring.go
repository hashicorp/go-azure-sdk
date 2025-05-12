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

type WorkloadNetworksListPortMirroringOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WorkloadNetworkPortMirroring
}

type WorkloadNetworksListPortMirroringCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WorkloadNetworkPortMirroring
}

type WorkloadNetworksListPortMirroringCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *WorkloadNetworksListPortMirroringCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// WorkloadNetworksListPortMirroring ...
func (c VMwaresClient) WorkloadNetworksListPortMirroring(ctx context.Context, id PrivateCloudId) (result WorkloadNetworksListPortMirroringOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &WorkloadNetworksListPortMirroringCustomPager{},
		Path:       fmt.Sprintf("%s/workloadNetworks/default/portMirroringProfiles", id.ID()),
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
		Values *[]WorkloadNetworkPortMirroring `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// WorkloadNetworksListPortMirroringComplete retrieves all the results into a single object
func (c VMwaresClient) WorkloadNetworksListPortMirroringComplete(ctx context.Context, id PrivateCloudId) (WorkloadNetworksListPortMirroringCompleteResult, error) {
	return c.WorkloadNetworksListPortMirroringCompleteMatchingPredicate(ctx, id, WorkloadNetworkPortMirroringOperationPredicate{})
}

// WorkloadNetworksListPortMirroringCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VMwaresClient) WorkloadNetworksListPortMirroringCompleteMatchingPredicate(ctx context.Context, id PrivateCloudId, predicate WorkloadNetworkPortMirroringOperationPredicate) (result WorkloadNetworksListPortMirroringCompleteResult, err error) {
	items := make([]WorkloadNetworkPortMirroring, 0)

	resp, err := c.WorkloadNetworksListPortMirroring(ctx, id)
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

	result = WorkloadNetworksListPortMirroringCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
