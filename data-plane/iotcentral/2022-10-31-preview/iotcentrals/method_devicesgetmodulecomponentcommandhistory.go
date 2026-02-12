package iotcentrals

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DevicesGetModuleComponentCommandHistoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeviceCommand
}

type DevicesGetModuleComponentCommandHistoryCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeviceCommand
}

type DevicesGetModuleComponentCommandHistoryCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DevicesGetModuleComponentCommandHistoryCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DevicesGetModuleComponentCommandHistory ...
func (c IotcentralsClient) DevicesGetModuleComponentCommandHistory(ctx context.Context, id ModuleComponentCommandId) (result DevicesGetModuleComponentCommandHistoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DevicesGetModuleComponentCommandHistoryCustomPager{},
		Path:       id.Path(),
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
		Values *[]DeviceCommand `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DevicesGetModuleComponentCommandHistoryComplete retrieves all the results into a single object
func (c IotcentralsClient) DevicesGetModuleComponentCommandHistoryComplete(ctx context.Context, id ModuleComponentCommandId) (DevicesGetModuleComponentCommandHistoryCompleteResult, error) {
	return c.DevicesGetModuleComponentCommandHistoryCompleteMatchingPredicate(ctx, id, DeviceCommandOperationPredicate{})
}

// DevicesGetModuleComponentCommandHistoryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DevicesGetModuleComponentCommandHistoryCompleteMatchingPredicate(ctx context.Context, id ModuleComponentCommandId, predicate DeviceCommandOperationPredicate) (result DevicesGetModuleComponentCommandHistoryCompleteResult, err error) {
	items := make([]DeviceCommand, 0)

	resp, err := c.DevicesGetModuleComponentCommandHistory(ctx, id)
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

	result = DevicesGetModuleComponentCommandHistoryCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
