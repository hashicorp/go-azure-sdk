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

type DevicesGetModuleCommandHistoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeviceCommand
}

type DevicesGetModuleCommandHistoryCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeviceCommand
}

type DevicesGetModuleCommandHistoryCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DevicesGetModuleCommandHistoryCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DevicesGetModuleCommandHistory ...
func (c IotcentralsClient) DevicesGetModuleCommandHistory(ctx context.Context, id ModuleCommandId) (result DevicesGetModuleCommandHistoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DevicesGetModuleCommandHistoryCustomPager{},
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

// DevicesGetModuleCommandHistoryComplete retrieves all the results into a single object
func (c IotcentralsClient) DevicesGetModuleCommandHistoryComplete(ctx context.Context, id ModuleCommandId) (DevicesGetModuleCommandHistoryCompleteResult, error) {
	return c.DevicesGetModuleCommandHistoryCompleteMatchingPredicate(ctx, id, DeviceCommandOperationPredicate{})
}

// DevicesGetModuleCommandHistoryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DevicesGetModuleCommandHistoryCompleteMatchingPredicate(ctx context.Context, id ModuleCommandId, predicate DeviceCommandOperationPredicate) (result DevicesGetModuleCommandHistoryCompleteResult, err error) {
	items := make([]DeviceCommand, 0)

	resp, err := c.DevicesGetModuleCommandHistory(ctx, id)
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

	result = DevicesGetModuleCommandHistoryCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
