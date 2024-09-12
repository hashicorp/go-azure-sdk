package me

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGetmberObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type ListGetmberObjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type ListGetmberObjectsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListGetmberObjectsOperationOptions() ListGetmberObjectsOperationOptions {
	return ListGetmberObjectsOperationOptions{}
}

func (o ListGetmberObjectsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGetmberObjectsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListGetmberObjectsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGetmberObjectsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGetmberObjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGetmberObjects - Invoke action getMemberObjects. Return all IDs for the groups, administrative units, and
// directory roles that a user, group, service principal, organizational contact, device, or directory object is a
// member of. This function is transitive. Note: Only users and role-enabled groups can be members of directory roles.
func (c MeClient) ListGetmberObjects(ctx context.Context, input ListGetmberObjectsRequest, options ListGetmberObjectsOperationOptions) (result ListGetmberObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListGetmberObjectsCustomPager{},
		Path:          "/me/getMemberObjects",
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
		Values *[]string `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGetmberObjectsComplete retrieves all the results into a single object
func (c MeClient) ListGetmberObjectsComplete(ctx context.Context, input ListGetmberObjectsRequest, options ListGetmberObjectsOperationOptions) (result ListGetmberObjectsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.ListGetmberObjects(ctx, input, options)
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

	result = ListGetmberObjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
