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

type StorageAppliancesListByResourceGroupOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StorageAppliance
}

type StorageAppliancesListByResourceGroupCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StorageAppliance
}

// StorageAppliancesListByResourceGroup ...
func (c NetworkcloudsClient) StorageAppliancesListByResourceGroup(ctx context.Context, id commonids.ResourceGroupId) (result StorageAppliancesListByResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.NetworkCloud/storageAppliances", id.ID()),
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
		Values *[]StorageAppliance `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// StorageAppliancesListByResourceGroupComplete retrieves all the results into a single object
func (c NetworkcloudsClient) StorageAppliancesListByResourceGroupComplete(ctx context.Context, id commonids.ResourceGroupId) (StorageAppliancesListByResourceGroupCompleteResult, error) {
	return c.StorageAppliancesListByResourceGroupCompleteMatchingPredicate(ctx, id, StorageApplianceOperationPredicate{})
}

// StorageAppliancesListByResourceGroupCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) StorageAppliancesListByResourceGroupCompleteMatchingPredicate(ctx context.Context, id commonids.ResourceGroupId, predicate StorageApplianceOperationPredicate) (result StorageAppliancesListByResourceGroupCompleteResult, err error) {
	items := make([]StorageAppliance, 0)

	resp, err := c.StorageAppliancesListByResourceGroup(ctx, id)
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

	result = StorageAppliancesListByResourceGroupCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
