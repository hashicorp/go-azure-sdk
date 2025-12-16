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

type DimensionsByExternalCloudProviderTypeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Dimension
}

type DimensionsByExternalCloudProviderTypeCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Dimension
}

type DimensionsByExternalCloudProviderTypeOperationOptions struct {
	Expand *string
	Filter *string
	Top    *int64
}

func DefaultDimensionsByExternalCloudProviderTypeOperationOptions() DimensionsByExternalCloudProviderTypeOperationOptions {
	return DimensionsByExternalCloudProviderTypeOperationOptions{}
}

func (o DimensionsByExternalCloudProviderTypeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DimensionsByExternalCloudProviderTypeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DimensionsByExternalCloudProviderTypeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type DimensionsByExternalCloudProviderTypeCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *DimensionsByExternalCloudProviderTypeCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// DimensionsByExternalCloudProviderType ...
func (c OpenapisClient) DimensionsByExternalCloudProviderType(ctx context.Context, id ExternalCloudProviderTypeId, options DimensionsByExternalCloudProviderTypeOperationOptions) (result DimensionsByExternalCloudProviderTypeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &DimensionsByExternalCloudProviderTypeCustomPager{},
		Path:          fmt.Sprintf("%s/dimensions", id.ID()),
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
		Values *[]Dimension `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// DimensionsByExternalCloudProviderTypeComplete retrieves all the results into a single object
func (c OpenapisClient) DimensionsByExternalCloudProviderTypeComplete(ctx context.Context, id ExternalCloudProviderTypeId, options DimensionsByExternalCloudProviderTypeOperationOptions) (DimensionsByExternalCloudProviderTypeCompleteResult, error) {
	return c.DimensionsByExternalCloudProviderTypeCompleteMatchingPredicate(ctx, id, options, DimensionOperationPredicate{})
}

// DimensionsByExternalCloudProviderTypeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OpenapisClient) DimensionsByExternalCloudProviderTypeCompleteMatchingPredicate(ctx context.Context, id ExternalCloudProviderTypeId, options DimensionsByExternalCloudProviderTypeOperationOptions, predicate DimensionOperationPredicate) (result DimensionsByExternalCloudProviderTypeCompleteResult, err error) {
	items := make([]Dimension, 0)

	resp, err := c.DimensionsByExternalCloudProviderType(ctx, id, options)
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

	result = DimensionsByExternalCloudProviderTypeCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
