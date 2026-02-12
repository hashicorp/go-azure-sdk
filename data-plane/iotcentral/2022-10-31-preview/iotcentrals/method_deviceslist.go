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

type DevicesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Device
}

type DevicesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Device
}

type DevicesListOperationOptions struct {
	Expand      *string
	Filter      *string
	Maxpagesize *int64
	Orderby     *string
}

func DefaultDevicesListOperationOptions() DevicesListOperationOptions {
	return DevicesListOperationOptions{}
}

func (o DevicesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DevicesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DevicesListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Filter != nil {
		out.Append("filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Maxpagesize != nil {
		out.Append("maxpagesize", fmt.Sprintf("%v", *o.Maxpagesize))
	}
	if o.Orderby != nil {
		out.Append("orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	return &out
}

type DevicesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DevicesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DevicesList ...
func (c IotcentralsClient) DevicesList(ctx context.Context, options DevicesListOperationOptions) (result DevicesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &DevicesListCustomPager{},
		Path:          "/devices",
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
		Values *[]Device `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DevicesListComplete retrieves all the results into a single object
func (c IotcentralsClient) DevicesListComplete(ctx context.Context, options DevicesListOperationOptions) (DevicesListCompleteResult, error) {
	return c.DevicesListCompleteMatchingPredicate(ctx, options, DeviceOperationPredicate{})
}

// DevicesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DevicesListCompleteMatchingPredicate(ctx context.Context, options DevicesListOperationOptions, predicate DeviceOperationPredicate) (result DevicesListCompleteResult, err error) {
	items := make([]Device, 0)

	resp, err := c.DevicesList(ctx, options)
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

	result = DevicesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
