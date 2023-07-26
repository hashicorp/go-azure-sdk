package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByMobileNetworkOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Service
}

type ListByMobileNetworkCompleteResult struct {
	Items []Service
}

// ListByMobileNetwork ...
func (c ServicesClient) ListByMobileNetwork(ctx context.Context, id MobileNetworkId) (result ListByMobileNetworkOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       fmt.Sprintf("%s/services", id.ID()),
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
		Values *[]Service `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByMobileNetworkComplete retrieves all the results into a single object
func (c ServicesClient) ListByMobileNetworkComplete(ctx context.Context, id MobileNetworkId) (ListByMobileNetworkCompleteResult, error) {
	return c.ListByMobileNetworkCompleteMatchingPredicate(ctx, id, ServiceOperationPredicate{})
}

// ListByMobileNetworkCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServicesClient) ListByMobileNetworkCompleteMatchingPredicate(ctx context.Context, id MobileNetworkId, predicate ServiceOperationPredicate) (result ListByMobileNetworkCompleteResult, err error) {
	items := make([]Service, 0)

	resp, err := c.ListByMobileNetwork(ctx, id)
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

	result = ListByMobileNetworkCompleteResult{
		Items: items,
	}
	return
}
