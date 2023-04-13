package operations

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

type MapsListSubscriptionOperationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]OperationDetail
}

type MapsListSubscriptionOperationsCompleteResult struct {
	Items []OperationDetail
}

// MapsListSubscriptionOperations ...
func (c OperationsClient) MapsListSubscriptionOperations(ctx context.Context, id commonids.SubscriptionId) (result MapsListSubscriptionOperationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/providers/Microsoft.Maps/operations", id.ID()),
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
		Values *[]OperationDetail `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// MapsListSubscriptionOperationsComplete retrieves all the results into a single object
func (c OperationsClient) MapsListSubscriptionOperationsComplete(ctx context.Context, id commonids.SubscriptionId) (MapsListSubscriptionOperationsCompleteResult, error) {
	return c.MapsListSubscriptionOperationsCompleteMatchingPredicate(ctx, id, OperationDetailOperationPredicate{})
}

// MapsListSubscriptionOperationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OperationsClient) MapsListSubscriptionOperationsCompleteMatchingPredicate(ctx context.Context, id commonids.SubscriptionId, predicate OperationDetailOperationPredicate) (result MapsListSubscriptionOperationsCompleteResult, err error) {
	items := make([]OperationDetail, 0)

	resp, err := c.MapsListSubscriptionOperations(ctx, id)
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

	result = MapsListSubscriptionOperationsCompleteResult{
		Items: items,
	}
	return
}
