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

type DeviceGroupsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeviceGroup
}

type DeviceGroupsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeviceGroup
}

type DeviceGroupsListOperationOptions struct {
	Filter      *string
	Maxpagesize *int64
	Orderby     *string
}

func DefaultDeviceGroupsListOperationOptions() DeviceGroupsListOperationOptions {
	return DeviceGroupsListOperationOptions{}
}

func (o DeviceGroupsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeviceGroupsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeviceGroupsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
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

type DeviceGroupsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DeviceGroupsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DeviceGroupsList ...
func (c IotcentralsClient) DeviceGroupsList(ctx context.Context, options DeviceGroupsListOperationOptions) (result DeviceGroupsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &DeviceGroupsListCustomPager{},
		Path:          "/deviceGroups",
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
		Values *[]DeviceGroup `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DeviceGroupsListComplete retrieves all the results into a single object
func (c IotcentralsClient) DeviceGroupsListComplete(ctx context.Context, options DeviceGroupsListOperationOptions) (DeviceGroupsListCompleteResult, error) {
	return c.DeviceGroupsListCompleteMatchingPredicate(ctx, options, DeviceGroupOperationPredicate{})
}

// DeviceGroupsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DeviceGroupsListCompleteMatchingPredicate(ctx context.Context, options DeviceGroupsListOperationOptions, predicate DeviceGroupOperationPredicate) (result DeviceGroupsListCompleteResult, err error) {
	items := make([]DeviceGroup, 0)

	resp, err := c.DeviceGroupsList(ctx, options)
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

	result = DeviceGroupsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
