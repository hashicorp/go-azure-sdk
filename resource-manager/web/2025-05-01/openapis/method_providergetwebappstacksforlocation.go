package openapis

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProviderGetWebAppStacksForLocationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]WebAppStack
}

type ProviderGetWebAppStacksForLocationCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []WebAppStack
}

type ProviderGetWebAppStacksForLocationOperationOptions struct {
	StackOsType *ProviderStackOsType
}

func DefaultProviderGetWebAppStacksForLocationOperationOptions() ProviderGetWebAppStacksForLocationOperationOptions {
	return ProviderGetWebAppStacksForLocationOperationOptions{}
}

func (o ProviderGetWebAppStacksForLocationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ProviderGetWebAppStacksForLocationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ProviderGetWebAppStacksForLocationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.StackOsType != nil {
		out.Append("stackOsType", fmt.Sprintf("%v", *o.StackOsType))
	}
	return &out
}

type ProviderGetWebAppStacksForLocationCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ProviderGetWebAppStacksForLocationCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ProviderGetWebAppStacksForLocation ...
func (c OpenapisClient) ProviderGetWebAppStacksForLocation(ctx context.Context, id LocationId, options ProviderGetWebAppStacksForLocationOperationOptions) (result ProviderGetWebAppStacksForLocationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ProviderGetWebAppStacksForLocationCustomPager{},
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

// ProviderGetWebAppStacksForLocationComplete retrieves all the results into a single object
func (c OpenapisClient) ProviderGetWebAppStacksForLocationComplete(ctx context.Context, id LocationId, options ProviderGetWebAppStacksForLocationOperationOptions) (ProviderGetWebAppStacksForLocationCompleteResult, error) {
	return c.ProviderGetWebAppStacksForLocationCompleteMatchingPredicate(ctx, id, options, WebAppStackOperationPredicate{})
}

// ProviderGetWebAppStacksForLocationCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) ProviderGetWebAppStacksForLocationCompleteMatchingPredicate(ctx context.Context, id LocationId, options ProviderGetWebAppStacksForLocationOperationOptions, predicate WebAppStackOperationPredicate) (result ProviderGetWebAppStacksForLocationCompleteResult, err error) {
	items := make([]WebAppStack, 0)

	resp, err := c.ProviderGetWebAppStacksForLocation(ctx, id, options)
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

	result = ProviderGetWebAppStacksForLocationCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
