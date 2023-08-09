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

type ListAtSubscriptionScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeploymentOperation
}

type ListAtSubscriptionScopeCompleteResult struct {
	Items []DeploymentOperation
}

type ListAtSubscriptionScopeOperationOptions struct {
	Top *int64
}

func DefaultListAtSubscriptionScopeOperationOptions() ListAtSubscriptionScopeOperationOptions {
	return ListAtSubscriptionScopeOperationOptions{}
}

func (o ListAtSubscriptionScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAtSubscriptionScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListAtSubscriptionScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

// ListAtSubscriptionScope ...
func (c DeploymentOperationsClient) ListAtSubscriptionScope(ctx context.Context, id ProviderDeploymentId, options ListAtSubscriptionScopeOperationOptions) (result ListAtSubscriptionScopeOperationResponse, err error) {
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

// ListAtSubscriptionScopeComplete retrieves all the results into a single object
func (c DeploymentOperationsClient) ListAtSubscriptionScopeComplete(ctx context.Context, id ProviderDeploymentId, options ListAtSubscriptionScopeOperationOptions) (ListAtSubscriptionScopeCompleteResult, error) {
	return c.ListAtSubscriptionScopeCompleteMatchingPredicate(ctx, id, options, DeploymentOperationOperationPredicate{})
}

// ListAtSubscriptionScopeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeploymentOperationsClient) ListAtSubscriptionScopeCompleteMatchingPredicate(ctx context.Context, id ProviderDeploymentId, options ListAtSubscriptionScopeOperationOptions, predicate DeploymentOperationOperationPredicate) (result ListAtSubscriptionScopeCompleteResult, err error) {
	items := make([]DeploymentOperation, 0)

	resp, err := c.ListAtSubscriptionScope(ctx, id, options)
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

	result = ListAtSubscriptionScopeCompleteResult{
		Items: items,
	}
	return
}
