package networkclouds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachinesListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BareMetalMachine
}

type BareMetalMachinesListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BareMetalMachine
}

// BareMetalMachinesListByResourceGroup ...
func (c NetworkcloudsClient) BareMetalMachinesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result BareMetalMachinesListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/bareMetalMachines", id.ID()),
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
		Values *[]BareMetalMachine `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BareMetalMachinesListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) BareMetalMachinesListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (BareMetalMachinesListByResourceGroupCompleteResult, error) {
	return c.BareMetalMachinesListByResourceGroupCompleteMatchingPredicate(ctx, id, BareMetalMachineOperationPredicate{})
}

// BareMetalMachinesListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) BareMetalMachinesListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate BareMetalMachineOperationPredicate) (result BareMetalMachinesListByResourceGroupCompleteResult, err error) {
	items := make([]BareMetalMachine, 0)

	resp, err := c.BareMetalMachinesListByResourceGroup(ctx, id)
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

	result = BareMetalMachinesListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
