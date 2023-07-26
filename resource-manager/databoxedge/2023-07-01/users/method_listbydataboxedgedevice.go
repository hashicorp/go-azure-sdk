package users

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByDataBoxEdgeDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]User
}

type ListByDataBoxEdgeDeviceCompleteResult struct {
	Items []User
}

type ListByDataBoxEdgeDeviceOperationOptions struct {
	Filter *string
}

func DefaultListByDataBoxEdgeDeviceOperationOptions() ListByDataBoxEdgeDeviceOperationOptions {
	return ListByDataBoxEdgeDeviceOperationOptions{}
}

func (o ListByDataBoxEdgeDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByDataBoxEdgeDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByDataBoxEdgeDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	return &out
}

// ListByDataBoxEdgeDevice ...
func (c UsersClient) ListByDataBoxEdgeDevice(ctx context.Context, id DataBoxEdgeDeviceId, options ListByDataBoxEdgeDeviceOperationOptions) (result ListByDataBoxEdgeDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/users", id.ID()),
		OptionsObject: options,
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
		Values *[]User `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByDataBoxEdgeDeviceComplete retrieves all the results into a single object
func (c UsersClient) ListByDataBoxEdgeDeviceComplete(ctx context.Context, id DataBoxEdgeDeviceId, options ListByDataBoxEdgeDeviceOperationOptions) (ListByDataBoxEdgeDeviceCompleteResult, error) {
	return c.ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx, id, options, UserOperationPredicate{})
}

// ListByDataBoxEdgeDeviceCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UsersClient) ListByDataBoxEdgeDeviceCompleteMatchingPredicate(ctx context.Context, id DataBoxEdgeDeviceId, options ListByDataBoxEdgeDeviceOperationOptions, predicate UserOperationPredicate) (result ListByDataBoxEdgeDeviceCompleteResult, err error) {
	items := make([]User, 0)

	resp, err := c.ListByDataBoxEdgeDevice(ctx, id, options)
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

	result = ListByDataBoxEdgeDeviceCompleteResult{
		Items: items,
	}
	return
}
