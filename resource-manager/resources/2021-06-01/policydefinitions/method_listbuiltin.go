package policydefinitions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListBuiltInOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]PolicyDefinition
}

type ListBuiltInCompleteResult struct {
	Items []PolicyDefinition
}

type ListBuiltInOperationOptions struct {
	Filter *string
	Top    *int64
}

func DefaultListBuiltInOperationOptions() ListBuiltInOperationOptions {
	return ListBuiltInOperationOptions{}
}

func (o ListBuiltInOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListBuiltInOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListBuiltInOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Filter != nil {
		out.Append("$filter", fmt.Sprintf("%v", *o.Filter))
	}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListBuiltIn ...
func (c PolicyDefinitionsClient) ListBuiltIn(ctx context.Context, options ListBuiltInOperationOptions) (result ListBuiltInOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          "/providers/Microsoft.Authorization/policyDefinitions",
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
		Values *[]PolicyDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBuiltInComplete retrieves all the results into a single object
func (c PolicyDefinitionsClient) ListBuiltInComplete(ctx context.Context, options ListBuiltInOperationOptions) (ListBuiltInCompleteResult, error) {
	return c.ListBuiltInCompleteMatchingPredicate(ctx, options, PolicyDefinitionOperationPredicate{})
}

// ListBuiltInCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyDefinitionsClient) ListBuiltInCompleteMatchingPredicate(ctx context.Context, options ListBuiltInOperationOptions, predicate PolicyDefinitionOperationPredicate) (result ListBuiltInCompleteResult, err error) {
	items := make([]PolicyDefinition, 0)

	resp, err := c.ListBuiltIn(ctx, options)
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

	result = ListBuiltInCompleteResult{
		Items: items,
	}
	return
}
