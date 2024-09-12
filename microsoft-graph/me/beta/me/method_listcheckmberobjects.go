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

type ListCheckmberObjectsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]string
}

type ListCheckmberObjectsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []string
}

type ListCheckmberObjectsOperationOptions struct {
	Skip *int64
	Top  *int64
}

func DefaultListCheckmberObjectsOperationOptions() ListCheckmberObjectsOperationOptions {
	return ListCheckmberObjectsOperationOptions{}
}

func (o ListCheckmberObjectsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCheckmberObjectsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListCheckmberObjectsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCheckmberObjectsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCheckmberObjectsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCheckmberObjects - Invoke action checkMemberObjects. Check for membership in a list of group IDs, administrative
// unit IDs, or directory role IDs, for the IDs of the specified user, group, service principal, organizational contact,
// device, or directory object. This method is transitive.
func (c MeClient) ListCheckmberObjects(ctx context.Context, input ListCheckmberObjectsRequest, options ListCheckmberObjectsOperationOptions) (result ListCheckmberObjectsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Pager:         &ListCheckmberObjectsCustomPager{},
		Path:          "/me/checkMemberObjects",
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

// ListCheckmberObjectsComplete retrieves all the results into a single object
func (c MeClient) ListCheckmberObjectsComplete(ctx context.Context, input ListCheckmberObjectsRequest, options ListCheckmberObjectsOperationOptions) (result ListCheckmberObjectsCompleteResult, err error) {
	items := make([]string, 0)

	resp, err := c.ListCheckmberObjects(ctx, input, options)
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

	result = ListCheckmberObjectsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
