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

type DeviceGroupsGetDevicesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Device
}

type DeviceGroupsGetDevicesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Device
}

type DeviceGroupsGetDevicesOperationOptions struct {
	Maxpagesize *int64
}

func DefaultDeviceGroupsGetDevicesOperationOptions() DeviceGroupsGetDevicesOperationOptions {
	return DeviceGroupsGetDevicesOperationOptions{}
}

func (o DeviceGroupsGetDevicesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeviceGroupsGetDevicesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeviceGroupsGetDevicesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Maxpagesize != nil {
		out.Append("maxpagesize", fmt.Sprintf("%v", *o.Maxpagesize))
	}
	return &out
}

type DeviceGroupsGetDevicesCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DeviceGroupsGetDevicesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DeviceGroupsGetDevices ...
func (c IotcentralsClient) DeviceGroupsGetDevices(ctx context.Context, id DeviceGroupId, options DeviceGroupsGetDevicesOperationOptions) (result DeviceGroupsGetDevicesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &DeviceGroupsGetDevicesCustomPager{},
		Path:          fmt.Sprintf("%s/devices", id.Path()),
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

// DeviceGroupsGetDevicesComplete retrieves all the results into a single object
func (c IotcentralsClient) DeviceGroupsGetDevicesComplete(ctx context.Context, id DeviceGroupId, options DeviceGroupsGetDevicesOperationOptions) (DeviceGroupsGetDevicesCompleteResult, error) {
	return c.DeviceGroupsGetDevicesCompleteMatchingPredicate(ctx, id, options, DeviceOperationPredicate{})
}

// DeviceGroupsGetDevicesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DeviceGroupsGetDevicesCompleteMatchingPredicate(ctx context.Context, id DeviceGroupId, options DeviceGroupsGetDevicesOperationOptions, predicate DeviceOperationPredicate) (result DeviceGroupsGetDevicesCompleteResult, err error) {
	items := make([]Device, 0)

	resp, err := c.DeviceGroupsGetDevices(ctx, id, options)
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

	result = DeviceGroupsGetDevicesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
