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

type DevicesGetCommandHistoryOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeviceCommand
}

type DevicesGetCommandHistoryCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeviceCommand
}

type DevicesGetCommandHistoryCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DevicesGetCommandHistoryCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DevicesGetCommandHistory ...
func (c IotcentralsClient) DevicesGetCommandHistory(ctx context.Context, id CommandId) (result DevicesGetCommandHistoryOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DevicesGetCommandHistoryCustomPager{},
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

// DevicesGetCommandHistoryComplete retrieves all the results into a single object
func (c IotcentralsClient) DevicesGetCommandHistoryComplete(ctx context.Context, id CommandId) (DevicesGetCommandHistoryCompleteResult, error) {
	return c.DevicesGetCommandHistoryCompleteMatchingPredicate(ctx, id, DeviceCommandOperationPredicate{})
}

// DevicesGetCommandHistoryCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DevicesGetCommandHistoryCompleteMatchingPredicate(ctx context.Context, id CommandId, predicate DeviceCommandOperationPredicate) (result DevicesGetCommandHistoryCompleteResult, err error) {
	items := make([]DeviceCommand, 0)

	resp, err := c.DevicesGetCommandHistory(ctx, id)
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

	result = DevicesGetCommandHistoryCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
