package provider

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetWebAppStacksForLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WebAppStack
}

type GetWebAppStacksForLocationCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WebAppStack
}

type GetWebAppStacksForLocationOperationOptions struct {
	StackOsType *ProviderStackOsType
}

func DefaultGetWebAppStacksForLocationOperationOptions() GetWebAppStacksForLocationOperationOptions {
	return GetWebAppStacksForLocationOperationOptions{}
}

func (o GetWebAppStacksForLocationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetWebAppStacksForLocationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetWebAppStacksForLocationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.StackOsType != nil {
		out.Append("stackOsType", fmt.Sprintf("%v", *o.StackOsType))
	}
	return &out
}

type GetWebAppStacksForLocationCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *GetWebAppStacksForLocationCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetWebAppStacksForLocation ...
func (c ProviderClient) GetWebAppStacksForLocation(ctx context.Context, id LocationId, options GetWebAppStacksForLocationOperationOptions) (result GetWebAppStacksForLocationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &GetWebAppStacksForLocationCustomPager{},
		Path:          fmt.Sprintf("%s/webAppStacks", id.ID()),
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
		Values *[]WebAppStack `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetWebAppStacksForLocationComplete retrieves all the results into a single object
func (c ProviderClient) GetWebAppStacksForLocationComplete(ctx context.Context, id LocationId, options GetWebAppStacksForLocationOperationOptions) (GetWebAppStacksForLocationCompleteResult, error) {
	return c.GetWebAppStacksForLocationCompleteMatchingPredicate(ctx, id, options, WebAppStackOperationPredicate{})
}

// GetWebAppStacksForLocationCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ProviderClient) GetWebAppStacksForLocationCompleteMatchingPredicate(ctx context.Context, id LocationId, options GetWebAppStacksForLocationOperationOptions, predicate WebAppStackOperationPredicate) (result GetWebAppStacksForLocationCompleteResult, err error) {
	items := make([]WebAppStack, 0)

	resp, err := c.GetWebAppStacksForLocation(ctx, id, options)
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

	result = GetWebAppStacksForLocationCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
