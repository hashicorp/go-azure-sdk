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

type ListAtTenantScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeploymentOperation
}

type ListAtTenantScopeCompleteResult struct {
	Items []DeploymentOperation
}

type ListAtTenantScopeOperationOptions struct {
	Top *int64
}

func DefaultListAtTenantScopeOperationOptions() ListAtTenantScopeOperationOptions {
	return ListAtTenantScopeOperationOptions{}
}

func (o ListAtTenantScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAtTenantScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListAtTenantScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListAtTenantScope ...
func (c DeploymentOperationsClient) ListAtTenantScope(ctx context.Context, id DeploymentId, options ListAtTenantScopeOperationOptions) (result ListAtTenantScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
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

// ListAtTenantScopeComplete retrieves all the results into a single object
func (c DeploymentOperationsClient) ListAtTenantScopeComplete(ctx context.Context, id DeploymentId, options ListAtTenantScopeOperationOptions) (ListAtTenantScopeCompleteResult, error) {
	return c.ListAtTenantScopeCompleteMatchingPredicate(ctx, id, options, DeploymentOperationOperationPredicate{})
}

// ListAtTenantScopeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeploymentOperationsClient) ListAtTenantScopeCompleteMatchingPredicate(ctx context.Context, id DeploymentId, options ListAtTenantScopeOperationOptions, predicate DeploymentOperationOperationPredicate) (result ListAtTenantScopeCompleteResult, err error) {
	items := make([]DeploymentOperation, 0)

	resp, err := c.ListAtTenantScope(ctx, id, options)
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

	result = ListAtTenantScopeCompleteResult{
		Items: items,
	}
	return
}
