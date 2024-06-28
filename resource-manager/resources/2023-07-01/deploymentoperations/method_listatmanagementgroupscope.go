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

type ListAtManagementGroupScopeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]DeploymentOperation
}

type ListAtManagementGroupScopeCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []DeploymentOperation
}

type ListAtManagementGroupScopeOperationOptions struct {
	Top *int64
}

func DefaultListAtManagementGroupScopeOperationOptions() ListAtManagementGroupScopeOperationOptions {
	return ListAtManagementGroupScopeOperationOptions{}
}

func (o ListAtManagementGroupScopeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListAtManagementGroupScopeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o ListAtManagementGroupScopeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Top != nil {
		out.Append("$top", fmt.Sprintf("%v", *o.Top))
	}
	return &out
}

type ListAtManagementGroupScopeCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListAtManagementGroupScopeCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListAtManagementGroupScope ...
func (c DeploymentOperationsClient) ListAtManagementGroupScope(ctx context.Context, id Providers2DeploymentId, options ListAtManagementGroupScopeOperationOptions) (result ListAtManagementGroupScopeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListAtManagementGroupScopeCustomPager{},
		Path:          fmt.Sprintf("%s/operations", id.ID()),
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

// ListAtManagementGroupScopeComplete retrieves all the results into a single object
func (c DeploymentOperationsClient) ListAtManagementGroupScopeComplete(ctx context.Context, id Providers2DeploymentId, options ListAtManagementGroupScopeOperationOptions) (ListAtManagementGroupScopeCompleteResult, error) {
	return c.ListAtManagementGroupScopeCompleteMatchingPredicate(ctx, id, options, DeploymentOperationOperationPredicate{})
}

// ListAtManagementGroupScopeCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DeploymentOperationsClient) ListAtManagementGroupScopeCompleteMatchingPredicate(ctx context.Context, id Providers2DeploymentId, options ListAtManagementGroupScopeOperationOptions, predicate DeploymentOperationOperationPredicate) (result ListAtManagementGroupScopeCompleteResult, err error) {
	items := make([]DeploymentOperation, 0)

	resp, err := c.ListAtManagementGroupScope(ctx, id, options)
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

	result = ListAtManagementGroupScopeCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
