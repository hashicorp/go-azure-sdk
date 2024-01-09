package policydefinitionversions

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
	Model        *[]PolicyDefinitionVersion
}

type ListBuiltInCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []PolicyDefinitionVersion
}

type ListBuiltInOperationOptions struct {
	Top *int64
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
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListBuiltIn ...
func (c PolicyDefinitionVersionsClient) ListBuiltIn(ctx context.Context, id PolicyDefinitionId, options ListBuiltInOperationOptions) (result ListBuiltInOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/versions", id.ID()),
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
		Values *[]PolicyDefinitionVersion `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListBuiltInComplete retrieves all the results into a single object
func (c PolicyDefinitionVersionsClient) ListBuiltInComplete(ctx context.Context, id PolicyDefinitionId, options ListBuiltInOperationOptions) (ListBuiltInCompleteResult, error) {
	return c.ListBuiltInCompleteMatchingPredicate(ctx, id, options, PolicyDefinitionVersionOperationPredicate{})
}

// ListBuiltInCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c PolicyDefinitionVersionsClient) ListBuiltInCompleteMatchingPredicate(ctx context.Context, id PolicyDefinitionId, options ListBuiltInOperationOptions, predicate PolicyDefinitionVersionOperationPredicate) (result ListBuiltInCompleteResult, err error) {
	items := make([]PolicyDefinitionVersion, 0)

	resp, err := c.ListBuiltIn(ctx, id, options)
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
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
