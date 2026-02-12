package privatelinkresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListSupportedOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PrivateLinkResource
}

type ListSupportedCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PrivateLinkResource
}

type ListSupportedOperationOptions struct {
	XMsClientRequestId *string
}

func DefaultListSupportedOperationOptions() ListSupportedOperationOptions {
	return ListSupportedOperationOptions{}
}

func (o ListSupportedOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.XMsClientRequestId != nil {
		out.Append("x-ms-client-request-id", fmt.Sprintf("%v", *o.XMsClientRequestId))
	}
	return &out
}

func (o ListSupportedOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListSupportedOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListSupportedCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListSupportedCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListSupported ...
func (c PrivateLinkResourcesClient) ListSupported(ctx context.Context, id SearchServiceId, options ListSupportedOperationOptions) (result ListSupportedOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListSupportedCustomPager{},
		Path:          fmt.Sprintf("%s/privateLinkResources", id.ID()),
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
		Values *[]PrivateLinkResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListSupportedComplete retrieves all the results into a single object
func (c PrivateLinkResourcesClient) ListSupportedComplete(ctx context.Context, id SearchServiceId, options ListSupportedOperationOptions) (ListSupportedCompleteResult, error) {
	return c.ListSupportedCompleteMatchingPredicate(ctx, id, options, PrivateLinkResourceOperationPredicate{})
}

// ListSupportedCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PrivateLinkResourcesClient) ListSupportedCompleteMatchingPredicate(ctx context.Context, id SearchServiceId, options ListSupportedOperationOptions, predicate PrivateLinkResourceOperationPredicate) (result ListSupportedCompleteResult, err error) {
	items := make([]PrivateLinkResource, 0)

	resp, err := c.ListSupported(ctx, id, options)
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

	result = ListSupportedCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
