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

type ConnectionRaiBlocklistsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RaiBlocklistPropertiesBasicResource
}

type ConnectionRaiBlocklistsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RaiBlocklistPropertiesBasicResource
}

type ConnectionRaiBlocklistsListOperationOptions struct {
	ProxyApiVersion *string
}

func DefaultConnectionRaiBlocklistsListOperationOptions() ConnectionRaiBlocklistsListOperationOptions {
	return ConnectionRaiBlocklistsListOperationOptions{}
}

func (o ConnectionRaiBlocklistsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ConnectionRaiBlocklistsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ConnectionRaiBlocklistsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ProxyApiVersion != nil {
		out.Append("proxy-api-version", fmt.Sprintf("%v", *o.ProxyApiVersion))
	}
	return &out
}

type ConnectionRaiBlocklistsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ConnectionRaiBlocklistsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ConnectionRaiBlocklistsList ...
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistsList(ctx context.Context, id ConnectionId, options ConnectionRaiBlocklistsListOperationOptions) (result ConnectionRaiBlocklistsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ConnectionRaiBlocklistsListCustomPager{},
		Path:          fmt.Sprintf("%s/raiBlocklists", id.ID()),
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
		Values *[]RaiBlocklistPropertiesBasicResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ConnectionRaiBlocklistsListComplete retrieves all the results into a single object
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistsListComplete(ctx context.Context, id ConnectionId, options ConnectionRaiBlocklistsListOperationOptions) (ConnectionRaiBlocklistsListCompleteResult, error) {
	return c.ConnectionRaiBlocklistsListCompleteMatchingPredicate(ctx, id, options, RaiBlocklistPropertiesBasicResourceOperationPredicate{})
}

// ConnectionRaiBlocklistsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistsListCompleteMatchingPredicate(ctx context.Context, id ConnectionId, options ConnectionRaiBlocklistsListOperationOptions, predicate RaiBlocklistPropertiesBasicResourceOperationPredicate) (result ConnectionRaiBlocklistsListCompleteResult, err error) {
	items := make([]RaiBlocklistPropertiesBasicResource, 0)

	resp, err := c.ConnectionRaiBlocklistsList(ctx, id, options)
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

	result = ConnectionRaiBlocklistsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
