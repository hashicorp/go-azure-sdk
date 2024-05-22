package usages

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByInstancePoolOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Usage
}

type ListByInstancePoolCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Usage
}

type ListByInstancePoolOperationOptions struct {
	ExpandChildren *bool
}

func DefaultListByInstancePoolOperationOptions() ListByInstancePoolOperationOptions {
	return ListByInstancePoolOperationOptions{}
}

func (o ListByInstancePoolOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByInstancePoolOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByInstancePoolOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ExpandChildren != nil {
		out.Append("expandChildren", fmt.Sprintf("%v", *o.ExpandChildren))
	}
	return &out
}

// ListByInstancePool ...
func (c UsagesClient) ListByInstancePool(ctx context.Context, id InstancePoolId, options ListByInstancePoolOperationOptions) (result ListByInstancePoolOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/usages", id.ID()),
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
		Values *[]Usage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByInstancePoolComplete retrieves all the results into a single object
func (c UsagesClient) ListByInstancePoolComplete(ctx context.Context, id InstancePoolId, options ListByInstancePoolOperationOptions) (ListByInstancePoolCompleteResult, error) {
	return c.ListByInstancePoolCompleteMatchingPredicate(ctx, id, options, UsageOperationPredicate{})
}

// ListByInstancePoolCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UsagesClient) ListByInstancePoolCompleteMatchingPredicate(ctx context.Context, id InstancePoolId, options ListByInstancePoolOperationOptions, predicate UsageOperationPredicate) (result ListByInstancePoolCompleteResult, err error) {
	items := make([]Usage, 0)

	resp, err := c.ListByInstancePool(ctx, id, options)
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

	result = ListByInstancePoolCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
