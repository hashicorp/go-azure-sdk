package keyvalues

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetKeyValuesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]KeyValueListResult
}

type GetKeyValuesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []KeyValueListResult
}

type GetKeyValuesOperationOptions struct {
	AcceptDatetime *string
	After          *string
	Key            *string
	Label          *string
	Select         *string
	SyncToken      *string
}

func DefaultGetKeyValuesOperationOptions() GetKeyValuesOperationOptions {
	return GetKeyValuesOperationOptions{}
}

func (o GetKeyValuesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.AcceptDatetime != nil {
		out.Append("Accept-Datetime", fmt.Sprintf("%v", *o.AcceptDatetime))
	}
	if o.SyncToken != nil {
		out.Append("Sync-Token", fmt.Sprintf("%v", *o.SyncToken))
	}
	return &out
}

func (o GetKeyValuesOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetKeyValuesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.After != nil {
		out.Append("After", fmt.Sprintf("%v", *o.After))
	}
	if o.Key != nil {
		out.Append("key", fmt.Sprintf("%v", *o.Key))
	}
	if o.Label != nil {
		out.Append("label", fmt.Sprintf("%v", *o.Label))
	}
	if o.Select != nil {
		out.Append("$Select", fmt.Sprintf("%v", *o.Select))
	}
	return &out
}

type GetKeyValuesCustomPager struct {
	NextLink *odata.Link `json:"@nextLink"`
}

func (p *GetKeyValuesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetKeyValues ...
func (c KeyValuesClient) GetKeyValues(ctx context.Context, options GetKeyValuesOperationOptions) (result GetKeyValuesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/vnd.microsoft.appconfig.kvset+json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &GetKeyValuesCustomPager{},
		Path:          "/kv",
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
		Values *[]KeyValueListResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetKeyValuesComplete retrieves all the results into a single object
func (c KeyValuesClient) GetKeyValuesComplete(ctx context.Context, options GetKeyValuesOperationOptions) (GetKeyValuesCompleteResult, error) {
	return c.GetKeyValuesCompleteMatchingPredicate(ctx, options, KeyValueListResultOperationPredicate{})
}

// GetKeyValuesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c KeyValuesClient) GetKeyValuesCompleteMatchingPredicate(ctx context.Context, options GetKeyValuesOperationOptions, predicate KeyValueListResultOperationPredicate) (result GetKeyValuesCompleteResult, err error) {
	items := make([]KeyValueListResult, 0)

	resp, err := c.GetKeyValues(ctx, options)
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

	result = GetKeyValuesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
