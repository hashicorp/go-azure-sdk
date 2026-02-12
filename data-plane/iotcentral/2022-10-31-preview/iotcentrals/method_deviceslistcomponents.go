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

type DevicesListComponentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]interface{}
}

type DevicesListComponentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []interface{}
}

type DevicesListComponentsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DevicesListComponentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DevicesListComponents ...
func (c IotcentralsClient) DevicesListComponents(ctx context.Context, id DeviceId) (result DevicesListComponentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &DevicesListComponentsCustomPager{},
		Path:       fmt.Sprintf("%s/components", id.Path()),
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
		Values *[]interface{} `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DevicesListComponentsComplete retrieves all the results into a single object
func (c IotcentralsClient) DevicesListComponentsComplete(ctx context.Context, id DeviceId) (result DevicesListComponentsCompleteResult, err error) {
	items := make([]interface{}, 0)

	resp, err := c.DevicesListComponents(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			items = append(items, v)
		}
	}

	result = DevicesListComponentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
