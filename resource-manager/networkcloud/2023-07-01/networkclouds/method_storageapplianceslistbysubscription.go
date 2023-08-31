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

type StorageAppliancesListBySubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StorageAppliance
}

type StorageAppliancesListBySubscriptionCompleteResult struct {
	Items []StorageAppliance
}

// StorageAppliancesListBySubscription ...
func (c NetworkcloudsClient) StorageAppliancesListBySubscription(ctx context.Context, id commonids.SubscriptionId) (result StorageAppliancesListBySubscriptionOperationResponse, err error) {
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

// StorageAppliancesListBySubscriptionComplete retrieves all the results into a single object
func (c NetworkcloudsClient) StorageAppliancesListBySubscriptionComplete(ctx context.Context, id commonids.SubscriptionId) (StorageAppliancesListBySubscriptionCompleteResult, error) {
	return c.StorageAppliancesListBySubscriptionCompleteMatchingPredicate(ctx, id, StorageApplianceOperationPredicate{})
}

// StorageAppliancesListBySubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c NetworkcloudsClient) StorageAppliancesListBySubscriptionCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate StorageApplianceOperationPredicate) (result StorageAppliancesListBySubscriptionCompleteResult, err error) {
	items := make([]StorageAppliance, 0)

	resp, err := c.StorageAppliancesListBySubscription(ctx, id)
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

	result = StorageAppliancesListBySubscriptionCompleteResult{
		Items: items,
	}
	return
}
