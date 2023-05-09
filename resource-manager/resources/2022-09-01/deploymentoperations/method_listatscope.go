package deploymentoperations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListAtScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeploymentOperation
}

type ListAtScopeCompleteResult struct {
	Items []DeploymentOperation
}

type ListAtScopeOperationOptions struct {
	Top *int64
}

func DefaultListAtScopeOperationOptions() ListAtScopeOperationOptions {
	return ListAtScopeOperationOptions{}
}

func (o ListAtScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAtScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListAtScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListAtScope ...
func (c DeploymentOperationsClient) ListAtScope(ctx context.Context, id ScopedDeploymentId, options ListAtScopeOperationOptions) (result ListAtScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		Path:          fmt.Sprintf("%s/operations", id.ID()),
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
		Values *[]DeploymentOperation `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListAtScopeComplete retrieves all the results into a single object
func (c DeploymentOperationsClient) ListAtScopeComplete(ctx context.Context, id ScopedDeploymentId, options ListAtScopeOperationOptions) (ListAtScopeCompleteResult, error) {
	return c.ListAtScopeCompleteMatchingPredicate(ctx, id, options, DeploymentOperationOperationPredicate{})
}

// ListAtScopeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeploymentOperationsClient) ListAtScopeCompleteMatchingPredicate(ctx context.Context, id ScopedDeploymentId, options ListAtScopeOperationOptions, predicate DeploymentOperationOperationPredicate) (result ListAtScopeCompleteResult, err error) {
	items := make([]DeploymentOperation, 0)

	resp, err := c.ListAtScope(ctx, id, options)
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

	result = ListAtScopeCompleteResult{
		Items: items,
	}
	return
}
