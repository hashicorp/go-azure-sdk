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

type DevicesListRelationshipsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeviceRelationship
}

type DevicesListRelationshipsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeviceRelationship
}

type DevicesListRelationshipsOperationOptions struct {
	Maxpagesize *int64
}

func DefaultDevicesListRelationshipsOperationOptions() DevicesListRelationshipsOperationOptions {
	return DevicesListRelationshipsOperationOptions{}
}

func (o DevicesListRelationshipsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DevicesListRelationshipsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DevicesListRelationshipsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Maxpagesize != nil {
		out.Append("maxpagesize", fmt.Sprintf("%v", *o.Maxpagesize))
	}
	return &out
}

type DevicesListRelationshipsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DevicesListRelationshipsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DevicesListRelationships ...
func (c IotcentralsClient) DevicesListRelationships(ctx context.Context, id DeviceId, options DevicesListRelationshipsOperationOptions) (result DevicesListRelationshipsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &DevicesListRelationshipsCustomPager{},
		Path:          fmt.Sprintf("%s/relationships", id.Path()),
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
		Values *[]DeviceRelationship `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DevicesListRelationshipsComplete retrieves all the results into a single object
func (c IotcentralsClient) DevicesListRelationshipsComplete(ctx context.Context, id DeviceId, options DevicesListRelationshipsOperationOptions) (DevicesListRelationshipsCompleteResult, error) {
	return c.DevicesListRelationshipsCompleteMatchingPredicate(ctx, id, options, DeviceRelationshipOperationPredicate{})
}

// DevicesListRelationshipsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) DevicesListRelationshipsCompleteMatchingPredicate(ctx context.Context, id DeviceId, options DevicesListRelationshipsOperationOptions, predicate DeviceRelationshipOperationPredicate) (result DevicesListRelationshipsCompleteResult, err error) {
	items := make([]DeviceRelationship, 0)

	resp, err := c.DevicesListRelationships(ctx, id, options)
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

	result = DevicesListRelationshipsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
