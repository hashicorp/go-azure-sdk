package datalakestoreaccounts

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
	Model        *[]DataLakeStoreAccountInformation
}

type ListByAccountCompleteResult struct {
	Items []DataLakeStoreAccountInformation
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

// ListByAccount ...
func (c DataLakeStoreAccountsClient) ListByAccount(ctx context.Context, id AccountId, options ListByAccountOperationOptions) (result ListByAccountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/dataLakeStoreAccounts", id.ID()),
		OptionsObject: options,
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
		Values *[]DataLakeStoreAccountInformation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByAccountComplete retrieves all the results into a single object
func (c DataLakeStoreAccountsClient) ListByAccountComplete(ctx context.Context, id AccountId, options ListByAccountOperationOptions) (ListByAccountCompleteResult, error) {
	return c.ListByAccountCompleteMatchingPredicate(ctx, id, options, DataLakeStoreAccountInformationOperationPredicate{})
}

// ListByAccountCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataLakeStoreAccountsClient) ListByAccountCompleteMatchingPredicate(ctx context.Context, id AccountId, options ListByAccountOperationOptions, predicate DataLakeStoreAccountInformationOperationPredicate) (result ListByAccountCompleteResult, err error) {
	items := make([]DataLakeStoreAccountInformation, 0)

	resp, err := c.ListByAccount(ctx, id, options)
	if err != nil {
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
		Items: items,
	}
	return
}
