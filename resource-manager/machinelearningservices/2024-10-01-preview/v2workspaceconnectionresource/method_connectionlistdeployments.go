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

type ConnectionListDeploymentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EndpointDeploymentResourcePropertiesBasicResource
}

type ConnectionListDeploymentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EndpointDeploymentResourcePropertiesBasicResource
}

type ConnectionListDeploymentsOperationOptions struct {
	ProxyApiVersion *string
}

func DefaultConnectionListDeploymentsOperationOptions() ConnectionListDeploymentsOperationOptions {
	return ConnectionListDeploymentsOperationOptions{}
}

func (o ConnectionListDeploymentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ConnectionListDeploymentsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ConnectionListDeploymentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ProxyApiVersion != nil {
		out.Append("proxy-api-version", fmt.Sprintf("%v", *o.ProxyApiVersion))
	}
	return &out
}

type ConnectionListDeploymentsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ConnectionListDeploymentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ConnectionListDeployments ...
func (c V2WorkspaceConnectionResourceClient) ConnectionListDeployments(ctx context.Context, id ConnectionId, options ConnectionListDeploymentsOperationOptions) (result ConnectionListDeploymentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ConnectionListDeploymentsCustomPager{},
		Path:          fmt.Sprintf("%s/deployments", id.ID()),
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
		Values *[]EndpointDeploymentResourcePropertiesBasicResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ConnectionListDeploymentsComplete retrieves all the results into a single object
func (c V2WorkspaceConnectionResourceClient) ConnectionListDeploymentsComplete(ctx context.Context, id ConnectionId, options ConnectionListDeploymentsOperationOptions) (ConnectionListDeploymentsCompleteResult, error) {
	return c.ConnectionListDeploymentsCompleteMatchingPredicate(ctx, id, options, EndpointDeploymentResourcePropertiesBasicResourceOperationPredicate{})
}

// ConnectionListDeploymentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c V2WorkspaceConnectionResourceClient) ConnectionListDeploymentsCompleteMatchingPredicate(ctx context.Context, id ConnectionId, options ConnectionListDeploymentsOperationOptions, predicate EndpointDeploymentResourcePropertiesBasicResourceOperationPredicate) (result ConnectionListDeploymentsCompleteResult, err error) {
	items := make([]EndpointDeploymentResourcePropertiesBasicResource, 0)

	resp, err := c.ConnectionListDeployments(ctx, id, options)
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

	result = ConnectionListDeploymentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
