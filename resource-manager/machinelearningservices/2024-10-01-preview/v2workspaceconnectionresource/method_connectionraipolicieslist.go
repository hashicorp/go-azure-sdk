package v2workspaceconnectionresource

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionRaiPoliciesListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RaiPolicyPropertiesBasicResource
}

type ConnectionRaiPoliciesListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RaiPolicyPropertiesBasicResource
}

type ConnectionRaiPoliciesListOperationOptions struct {
	ProxyApiVersion *string
}

func DefaultConnectionRaiPoliciesListOperationOptions() ConnectionRaiPoliciesListOperationOptions {
	return ConnectionRaiPoliciesListOperationOptions{}
}

func (o ConnectionRaiPoliciesListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ConnectionRaiPoliciesListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ConnectionRaiPoliciesListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ProxyApiVersion != nil {
		out.Append("proxy-api-version", fmt.Sprintf("%v", *o.ProxyApiVersion))
	}
	return &out
}

type ConnectionRaiPoliciesListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ConnectionRaiPoliciesListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ConnectionRaiPoliciesList ...
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiPoliciesList(ctx context.Context, id ConnectionId, options ConnectionRaiPoliciesListOperationOptions) (result ConnectionRaiPoliciesListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ConnectionRaiPoliciesListCustomPager{},
		Path:          fmt.Sprintf("%s/raiPolicies", id.ID()),
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
		Values *[]RaiPolicyPropertiesBasicResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ConnectionRaiPoliciesListComplete retrieves all the results into a single object
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiPoliciesListComplete(ctx context.Context, id ConnectionId, options ConnectionRaiPoliciesListOperationOptions) (ConnectionRaiPoliciesListCompleteResult, error) {
	return c.ConnectionRaiPoliciesListCompleteMatchingPredicate(ctx, id, options, RaiPolicyPropertiesBasicResourceOperationPredicate{})
}

// ConnectionRaiPoliciesListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiPoliciesListCompleteMatchingPredicate(ctx context.Context, id ConnectionId, options ConnectionRaiPoliciesListOperationOptions, predicate RaiPolicyPropertiesBasicResourceOperationPredicate) (result ConnectionRaiPoliciesListCompleteResult, err error) {
	items := make([]RaiPolicyPropertiesBasicResource, 0)

	resp, err := c.ConnectionRaiPoliciesList(ctx, id, options)
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

	result = ConnectionRaiPoliciesListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
