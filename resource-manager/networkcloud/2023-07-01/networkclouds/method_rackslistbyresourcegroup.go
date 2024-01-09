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

type RacksListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Rack
}

type RacksListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Rack
}

// RacksListByResourceGroup ...
func (c NetworkcloudsClient) RacksListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result RacksListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/racks", id.ID()),
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
		Values *[]Rack `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// RacksListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) RacksListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (RacksListByResourceGroupCompleteResult, error) {
	return c.RacksListByResourceGroupCompleteMatchingPredicate(ctx, id, RackOperationPredicate{})
}

// RacksListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) RacksListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate RackOperationPredicate) (result RacksListByResourceGroupCompleteResult, err error) {
	items := make([]Rack, 0)

	resp, err := c.RacksListByResourceGroup(ctx, id)
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

	result = RacksListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
