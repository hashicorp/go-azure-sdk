package iotcentrals

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsersListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]User
}

type UsersListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []User
}

type UsersListOperationOptions struct {
	Maxpagesize *int64
}

func DefaultUsersListOperationOptions() UsersListOperationOptions {
	return UsersListOperationOptions{}
}

func (o UsersListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UsersListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o UsersListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Maxpagesize != nil {
		out.Append("maxpagesize", fmt.Sprintf("%v", *o.Maxpagesize))
	}
	return &out
}

type UsersListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *UsersListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// UsersList ...
func (c IotcentralsClient) UsersList(ctx context.Context, options UsersListOperationOptions) (result UsersListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &UsersListCustomPager{},
		Path:          "/users",
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]User, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := UnmarshalUserImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for User (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// UsersListComplete retrieves all the results into a single object
func (c IotcentralsClient) UsersListComplete(ctx context.Context, options UsersListOperationOptions) (UsersListCompleteResult, error) {
	return c.UsersListCompleteMatchingPredicate(ctx, options, UserOperationPredicate{})
}

// UsersListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) UsersListCompleteMatchingPredicate(ctx context.Context, options UsersListOperationOptions, predicate UserOperationPredicate) (result UsersListCompleteResult, err error) {
	items := make([]User, 0)

	resp, err := c.UsersList(ctx, options)
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

	result = UsersListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
