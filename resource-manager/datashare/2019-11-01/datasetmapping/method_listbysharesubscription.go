package datasetmapping

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByShareSubscriptionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DataSetMapping
}

type ListByShareSubscriptionCompleteResult struct {
	Items []DataSetMapping
}

type ListByShareSubscriptionOperationOptions struct {
	Filter  *string
	Orderby *string
}

func DefaultListByShareSubscriptionOperationOptions() ListByShareSubscriptionOperationOptions {
	return ListByShareSubscriptionOperationOptions{}
}

func (o ListByShareSubscriptionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListByShareSubscriptionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListByShareSubscriptionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Orderby != nil {
		out.Append("$orderby", fmt.Sprintf("%v", *o.Orderby))
	}
	return &out
}

// ListByShareSubscription ...
func (c DataSetMappingClient) ListByShareSubscription(ctx context.Context, id ShareSubscriptionId, options ListByShareSubscriptionOperationOptions) (result ListByShareSubscriptionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/dataSetMappings", id.ID()),
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
		Values *[]DataSetMapping `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListByShareSubscriptionComplete retrieves all the results into a single object
func (c DataSetMappingClient) ListByShareSubscriptionComplete(ctx context.Context, id ShareSubscriptionId, options ListByShareSubscriptionOperationOptions) (ListByShareSubscriptionCompleteResult, error) {
	return c.ListByShareSubscriptionCompleteMatchingPredicate(ctx, id, options, DataSetMappingOperationPredicate{})
}

// ListByShareSubscriptionCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DataSetMappingClient) ListByShareSubscriptionCompleteMatchingPredicate(ctx context.Context, id ShareSubscriptionId, options ListByShareSubscriptionOperationOptions, predicate DataSetMappingOperationPredicate) (result ListByShareSubscriptionCompleteResult, err error) {
	items := make([]DataSetMapping, 0)

	resp, err := c.ListByShareSubscription(ctx, id, options)
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

	result = ListByShareSubscriptionCompleteResult{
		Items: items,
	}
	return
}
