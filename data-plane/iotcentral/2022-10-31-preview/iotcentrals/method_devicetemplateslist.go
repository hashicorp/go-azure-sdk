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

type DeviceTemplatesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeviceTemplate
}

type DeviceTemplatesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeviceTemplate
}

type DeviceTemplatesListOperationOptions struct {
	Filter      *string
	Maxpagesize *int64
	Orderby     *string
}

func DefaultDeviceTemplatesListOperationOptions() DeviceTemplatesListOperationOptions {
	return DeviceTemplatesListOperationOptions{}
}

func (o DeviceTemplatesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeviceTemplatesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeviceTemplatesListOperationOptions) ToQuery() *client.QueryParams {
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

type DeviceTemplatesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DeviceTemplatesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DeviceTemplatesList ...
func (c IotcentralsClient) DeviceTemplatesList(ctx context.Context, options DeviceTemplatesListOperationOptions) (result DeviceTemplatesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &DeviceTemplatesListCustomPager{},
		Path:          "/deviceTemplates",
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
		Values *[]DeviceTemplate `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DeviceTemplatesListComplete retrieves all the results into a single object
func (c IotcentralsClient) DeviceTemplatesListComplete(ctx context.Context, options DeviceTemplatesListOperationOptions) (DeviceTemplatesListCompleteResult, error) {
	return c.DeviceTemplatesListCompleteMatchingPredicate(ctx, options, DeviceTemplateOperationPredicate{})
}

// DeviceTemplatesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DeviceTemplatesListCompleteMatchingPredicate(ctx context.Context, options DeviceTemplatesListOperationOptions, predicate DeviceTemplateOperationPredicate) (result DeviceTemplatesListCompleteResult, err error) {
	items := make([]DeviceTemplate, 0)

	resp, err := c.DeviceTemplatesList(ctx, options)
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

	result = DeviceTemplatesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
