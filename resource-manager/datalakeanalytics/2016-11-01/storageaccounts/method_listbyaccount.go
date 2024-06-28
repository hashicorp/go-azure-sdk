package storageaccounts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByAccountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]StorageAccountInformation
}

type ListByAccountCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []StorageAccountInformation
}

type ListByAccountOperationOptions struct {
	Count   *bool
	Filter  *string
	Orderby *string
	Select  *string
	Skip    *int64
	Top     *int64
}

func DefaultListByAccountOperationOptions() ListByAccountOperationOptions {
	return ListByAccountOperationOptions{}
}

func (o ListByAccountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByAccountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByAccountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Count != nil {
		out.Append("$count", fmt.Sprintf("%v", *o.Count))
	}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	if o.Select != nil {
		out.Append("$select", fmt.Sprintf("%v", *o.Select))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListByAccountCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListByAccountCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListByAccount ...
func (c StorageAccountsClient) ListByAccount(ctx context.Context, id AccountId, options ListByAccountOperationOptions) (result ListByAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListByAccountCustomPager{},
		Path:          fmt.Sprintf("%s/storageAccounts", id.ID()),
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
		Values *[]StorageAccountInformation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByAccountComplete retrieves all the results into a single object
func (c StorageAccountsClient) ListByAccountComplete(ctx context.Context, id AccountId, options ListByAccountOperationOptions) (ListByAccountCompleteResult, error) {
	return c.ListByAccountCompleteMatchingPredicate(ctx, id, options, StorageAccountInformationOperationPredicate{})
}

// ListByAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c StorageAccountsClient) ListByAccountCompleteMatchingPredicate(ctx context.Context, id AccountId, options ListByAccountOperationOptions, predicate StorageAccountInformationOperationPredicate) (result ListByAccountCompleteResult, err error) {
	items := make([]StorageAccountInformation, 0)

	resp, err := c.ListByAccount(ctx, id, options)
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

	result = ListByAccountCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
