package keys

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetKeysOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]KeyListResult
}

type GetKeysCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []KeyListResult
}

type GetKeysOperationOptions struct {
	AcceptDatetime *string
	After          *string
	Name           *string
	SyncToken      *string
}

func DefaultGetKeysOperationOptions() GetKeysOperationOptions {
	return GetKeysOperationOptions{}
}

func (o GetKeysOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.AcceptDatetime != nil {
		out.Append("Accept-Datetime", fmt.Sprintf("%v", *o.AcceptDatetime))
	}
	if o.SyncToken != nil {
		out.Append("Sync-Token", fmt.Sprintf("%v", *o.SyncToken))
	}
	return &out
}

func (o GetKeysOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o GetKeysOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.After != nil {
		out.Append("After", fmt.Sprintf("%v", *o.After))
	}
	if o.Name != nil {
		out.Append("name", fmt.Sprintf("%v", *o.Name))
	}
	return &out
}

type GetKeysCustomPager struct {
	NextLink *odata.Link `json:"@nextLink"`
}

func (p *GetKeysCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// GetKeys ...
func (c KeysClient) GetKeys(ctx context.Context, options GetKeysOperationOptions) (result GetKeysOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/vnd.microsoft.appconfig.keyset+json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &GetKeysCustomPager{},
		Path:          "/keys",
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
		Values *[]KeyListResult `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// GetKeysComplete retrieves all the results into a single object
func (c KeysClient) GetKeysComplete(ctx context.Context, options GetKeysOperationOptions) (GetKeysCompleteResult, error) {
	return c.GetKeysCompleteMatchingPredicate(ctx, options, KeyListResultOperationPredicate{})
}

// GetKeysCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c KeysClient) GetKeysCompleteMatchingPredicate(ctx context.Context, options GetKeysOperationOptions, predicate KeyListResultOperationPredicate) (result GetKeysCompleteResult, err error) {
	items := make([]KeyListResult, 0)

	resp, err := c.GetKeys(ctx, options)
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

	result = GetKeysCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
