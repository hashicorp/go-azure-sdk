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

type ConnectionGetModelsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EndpointModelProperties
}

type ConnectionGetModelsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EndpointModelProperties
}

type ConnectionGetModelsOperationOptions struct {
	ProxyApiVersion *string
}

func DefaultConnectionGetModelsOperationOptions() ConnectionGetModelsOperationOptions {
	return ConnectionGetModelsOperationOptions{}
}

func (o ConnectionGetModelsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ConnectionGetModelsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ConnectionGetModelsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ProxyApiVersion != nil {
		out.Append("proxy-api-version", fmt.Sprintf("%v", *o.ProxyApiVersion))
	}
	return &out
}

type ConnectionGetModelsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ConnectionGetModelsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ConnectionGetModels ...
func (c V2WorkspaceConnectionResourceClient) ConnectionGetModels(ctx context.Context, id ConnectionId, options ConnectionGetModelsOperationOptions) (result ConnectionGetModelsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ConnectionGetModelsCustomPager{},
		Path:          fmt.Sprintf("%s/models", id.ID()),
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
		Values *[]EndpointModelProperties `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ConnectionGetModelsComplete retrieves all the results into a single object
func (c V2WorkspaceConnectionResourceClient) ConnectionGetModelsComplete(ctx context.Context, id ConnectionId, options ConnectionGetModelsOperationOptions) (ConnectionGetModelsCompleteResult, error) {
	return c.ConnectionGetModelsCompleteMatchingPredicate(ctx, id, options, EndpointModelPropertiesOperationPredicate{})
}

// ConnectionGetModelsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c V2WorkspaceConnectionResourceClient) ConnectionGetModelsCompleteMatchingPredicate(ctx context.Context, id ConnectionId, options ConnectionGetModelsOperationOptions, predicate EndpointModelPropertiesOperationPredicate) (result ConnectionGetModelsCompleteResult, err error) {
	items := make([]EndpointModelProperties, 0)

	resp, err := c.ConnectionGetModels(ctx, id, options)
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

	result = ConnectionGetModelsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
