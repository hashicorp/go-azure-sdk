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

type PTUQuotaListAvailableOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]AvailableQuota
}

type PTUQuotaListAvailableCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []AvailableQuota
}

type PTUQuotaListAvailableOperationOptions struct {
	Skip *string
}

func DefaultPTUQuotaListAvailableOperationOptions() PTUQuotaListAvailableOperationOptions {
	return PTUQuotaListAvailableOperationOptions{}
}

func (o PTUQuotaListAvailableOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o PTUQuotaListAvailableOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o PTUQuotaListAvailableOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	return &out
}

type PTUQuotaListAvailableCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *PTUQuotaListAvailableCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// PTUQuotaListAvailable ...
func (c QuotaClient) PTUQuotaListAvailable(ctx context.Context, id LocationId, options PTUQuotaListAvailableOperationOptions) (result PTUQuotaListAvailableOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &PTUQuotaListAvailableCustomPager{},
		Path:          fmt.Sprintf("%s/availableQuota", id.ID()),
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
		Values *[]AvailableQuota `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// PTUQuotaListAvailableComplete retrieves all the results into a single object
func (c QuotaClient) PTUQuotaListAvailableComplete(ctx context.Context, id LocationId, options PTUQuotaListAvailableOperationOptions) (PTUQuotaListAvailableCompleteResult, error) {
	return c.PTUQuotaListAvailableCompleteMatchingPredicate(ctx, id, options, AvailableQuotaOperationPredicate{})
}

// PTUQuotaListAvailableCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c QuotaClient) PTUQuotaListAvailableCompleteMatchingPredicate(ctx context.Context, id LocationId, options PTUQuotaListAvailableOperationOptions, predicate AvailableQuotaOperationPredicate) (result PTUQuotaListAvailableCompleteResult, err error) {
	items := make([]AvailableQuota, 0)

	resp, err := c.PTUQuotaListAvailable(ctx, id, options)
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

	result = PTUQuotaListAvailableCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
