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

type ListGetmberGroupsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type ListGetmberGroupsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type ListGetmberGroupsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListGetmberGroupsOperationOptions() ListGetmberGroupsOperationOptions {
	return ListGetmberGroupsOperationOptions{}
}

func (o ListGetmberGroupsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListGetmberGroupsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListGetmberGroupsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListGetmberGroupsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListGetmberGroupsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListGetmberGroups - Invoke action getMemberGroups. Return all the group IDs for the groups that the specified user,
// group, service principal, organizational contact, device, or directory object is a member of. This function is
// transitive. This API returns up to 11,000 group IDs. If more than 11,000 results are available, it returns a 400 Bad
// Request error with the DirectoryResultSizeLimitExceeded error code. If you get the DirectoryResultSizeLimitExceeded
// error code, use the List group transitive memberOf API instead.
func (c MeClient) ListGetmberGroups(ctx context.Context, input ListGetmberGroupsRequest, options ListGetmberGroupsOperationOptions) (result ListGetmberGroupsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListGetmberGroupsCustomPager{},
		Path:          "/me/getMemberGroups",
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

// ListGetmberGroupsComplete retrieves all the results into a single object
func (c MeClient) ListGetmberGroupsComplete(ctx context.Context, input ListGetmberGroupsRequest, options ListGetmberGroupsOperationOptions) (result ListGetmberGroupsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.ListGetmberGroups(ctx, input, options)
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

	result = ListGetmberGroupsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
