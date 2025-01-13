package quota

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PTUQuotaListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]UsageAndQuotaDetails
}

type PTUQuotaListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []UsageAndQuotaDetails
}

type PTUQuotaListOperationOptions struct {
	Skip *string
}

func DefaultPTUQuotaListOperationOptions() PTUQuotaListOperationOptions {
	return PTUQuotaListOperationOptions{}
}

func (o PTUQuotaListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PTUQuotaListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PTUQuotaListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	return &out
}

type PTUQuotaListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PTUQuotaListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PTUQuotaList ...
func (c QuotaClient) PTUQuotaList(ctx context.Context, id LocationId, options PTUQuotaListOperationOptions) (result PTUQuotaListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &PTUQuotaListCustomPager{},
		Path:          fmt.Sprintf("%s/quotaAndUsage", id.ID()),
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
		Values *[]UsageAndQuotaDetails `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PTUQuotaListComplete retrieves all the results into a single object
func (c QuotaClient) PTUQuotaListComplete(ctx context.Context, id LocationId, options PTUQuotaListOperationOptions) (PTUQuotaListCompleteResult, error) {
	return c.PTUQuotaListCompleteMatchingPredicate(ctx, id, options, UsageAndQuotaDetailsOperationPredicate{})
}

// PTUQuotaListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c QuotaClient) PTUQuotaListCompleteMatchingPredicate(ctx context.Context, id LocationId, options PTUQuotaListOperationOptions, predicate UsageAndQuotaDetailsOperationPredicate) (result PTUQuotaListCompleteResult, err error) {
	items := make([]UsageAndQuotaDetails, 0)

	resp, err := c.PTUQuotaList(ctx, id, options)
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

	result = PTUQuotaListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
