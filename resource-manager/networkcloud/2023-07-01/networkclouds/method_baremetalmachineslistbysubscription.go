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

type BareMetalMachinesListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BareMetalMachine
}

type BareMetalMachinesListBySubscriptionCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []BareMetalMachine
}

// BareMetalMachinesListBySubscription ...
func (c NetworkcloudsClient) BareMetalMachinesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result BareMetalMachinesListBySubscriptionOperationResponse, err error) {
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

// BareMetalMachinesListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) BareMetalMachinesListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (BareMetalMachinesListBySubscriptionCompleteResult, error) {
	return c.BareMetalMachinesListBySubscriptionCompleteMatchingPredicate(ctx, id, BareMetalMachineOperationPredicate{})
}

// BareMetalMachinesListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) BareMetalMachinesListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate BareMetalMachineOperationPredicate) (result BareMetalMachinesListBySubscriptionCompleteResult, err error) {
	items := make([]BareMetalMachine, 0)

	resp, err := c.BareMetalMachinesListBySubscription(ctx, id)
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

	result = BareMetalMachinesListBySubscriptionCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
