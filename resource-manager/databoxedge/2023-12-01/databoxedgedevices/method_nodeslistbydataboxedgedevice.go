package databoxedgedevices

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NodesListByDataBoxEdgeDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Node
}

type NodesListByDataBoxEdgeDeviceCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Node
}

type NodesListByDataBoxEdgeDeviceCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *NodesListByDataBoxEdgeDeviceCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// NodesListByDataBoxEdgeDevice ...
func (c DataBoxEdgeDevicesClient) NodesListByDataBoxEdgeDevice(ctx context.Context, id DataBoxEdgeDeviceId) (result NodesListByDataBoxEdgeDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &NodesListByDataBoxEdgeDeviceCustomPager{},
		Path:       fmt.Sprintf("%s/nodes", id.ID()),
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
		Values *[]Node `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// NodesListByDataBoxEdgeDeviceComplete retrieves all the results into a single object
func (c DataBoxEdgeDevicesClient) NodesListByDataBoxEdgeDeviceComplete(ctx context.Context, id DataBoxEdgeDeviceId) (NodesListByDataBoxEdgeDeviceCompleteResult, error) {
	return c.NodesListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx, id, NodeOperationPredicate{})
}

// NodesListByDataBoxEdgeDeviceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataBoxEdgeDevicesClient) NodesListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx context.Context, id DataBoxEdgeDeviceId, predicate NodeOperationPredicate) (result NodesListByDataBoxEdgeDeviceCompleteResult, err error) {
	items := make([]Node, 0)

	resp, err := c.NodesListByDataBoxEdgeDevice(ctx, id)
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

	result = NodesListByDataBoxEdgeDeviceCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
