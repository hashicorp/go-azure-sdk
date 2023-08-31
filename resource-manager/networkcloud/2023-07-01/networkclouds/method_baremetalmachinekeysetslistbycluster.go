package networkclouds

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BareMetalMachineKeySetsListByClusterOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]BareMetalMachineKeySet
}

type BareMetalMachineKeySetsListByClusterCompleteResult struct {
	Items []BareMetalMachineKeySet
}

// BareMetalMachineKeySetsListByCluster ...
func (c NetworkcloudsClient) BareMetalMachineKeySetsListByCluster(ctx context.Context, id ClusterId) (result BareMetalMachineKeySetsListByClusterOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/bareMetalMachineKeySets", id.ID()),
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
		Values *[]BareMetalMachineKeySet `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// BareMetalMachineKeySetsListByClusterComplete retrieves all the results into a single object
func (c NetworkcloudsClient) BareMetalMachineKeySetsListByClusterComplete(ctx context.Context, id ClusterId) (BareMetalMachineKeySetsListByClusterCompleteResult, error) {
	return c.BareMetalMachineKeySetsListByClusterCompleteMatchingPredicate(ctx, id, BareMetalMachineKeySetOperationPredicate{})
}

// BareMetalMachineKeySetsListByClusterCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) BareMetalMachineKeySetsListByClusterCompleteMatchingPredicate(ctx context.Context, id ClusterId, predicate BareMetalMachineKeySetOperationPredicate) (result BareMetalMachineKeySetsListByClusterCompleteResult, err error) {
	items := make([]BareMetalMachineKeySet, 0)

	resp, err := c.BareMetalMachineKeySetsListByCluster(ctx, id)
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

	result = BareMetalMachineKeySetsListByClusterCompleteResult{
		Items: items,
	}
	return
}
