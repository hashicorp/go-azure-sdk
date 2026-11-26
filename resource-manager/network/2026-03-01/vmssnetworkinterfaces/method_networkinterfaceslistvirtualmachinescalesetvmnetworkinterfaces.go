package vmssnetworkinterfaces

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]NetworkInterface
}

type NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []NetworkInterface
}

type NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfaces ...
func (c VMSSNetworkInterfacesClient) NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfaces(ctx context.Context, id VirtualMachineId) (result NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesCustomPager{},
		Path:       fmt.Sprintf("%s/networkInterfaces", id.ID()),
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
		Values *[]NetworkInterface `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesComplete retrieves all the results into a single object
func (c VMSSNetworkInterfacesClient) NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesComplete(ctx context.Context, id VirtualMachineId) (NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesCompleteResult, error) {
	return c.NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesCompleteMatchingPredicate(ctx, id, NetworkInterfaceOperationPredicate{})
}

// NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VMSSNetworkInterfacesClient) NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesCompleteMatchingPredicate(ctx context.Context, id VirtualMachineId, predicate NetworkInterfaceOperationPredicate) (result NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesCompleteResult, err error) {
	items := make([]NetworkInterface, 0)

	resp, err := c.NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfaces(ctx, id)
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

	result = NetworkInterfacesListVirtualMachineScaleSetVMNetworkInterfacesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
