package attacheddatanetworks

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByPacketCoreDataPlaneOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AttachedDataNetwork
}

type ListByPacketCoreDataPlaneCompleteResult struct {
	Items []AttachedDataNetwork
}

// ListByPacketCoreDataPlane ...
func (c AttachedDataNetworksClient) ListByPacketCoreDataPlane(ctx context.Context, id PacketCoreDataPlaneId) (result ListByPacketCoreDataPlaneOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/attachedDataNetworks", id.ID()),
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
		Values *[]AttachedDataNetwork `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByPacketCoreDataPlaneComplete retrieves all the results into a single object
func (c AttachedDataNetworksClient) ListByPacketCoreDataPlaneComplete(ctx context.Context, id PacketCoreDataPlaneId) (ListByPacketCoreDataPlaneCompleteResult, error) {
	return c.ListByPacketCoreDataPlaneCompleteMatchingPredicate(ctx, id, AttachedDataNetworkOperationPredicate{})
}

// ListByPacketCoreDataPlaneCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c AttachedDataNetworksClient) ListByPacketCoreDataPlaneCompleteMatchingPredicate(ctx context.Context, id PacketCoreDataPlaneId, predicate AttachedDataNetworkOperationPredicate) (result ListByPacketCoreDataPlaneCompleteResult, err error) {
	items := make([]AttachedDataNetwork, 0)

	resp, err := c.ListByPacketCoreDataPlane(ctx, id)
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

	result = ListByPacketCoreDataPlaneCompleteResult{
		Items: items,
	}
	return
}
