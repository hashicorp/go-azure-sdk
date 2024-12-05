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

type ConnectionRaiBlocklistItemsListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]RaiBlocklistItemPropertiesBasicResource
}

type ConnectionRaiBlocklistItemsListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []RaiBlocklistItemPropertiesBasicResource
}

type ConnectionRaiBlocklistItemsListOperationOptions struct {
	ProxyApiVersion *string
}

func DefaultConnectionRaiBlocklistItemsListOperationOptions() ConnectionRaiBlocklistItemsListOperationOptions {
	return ConnectionRaiBlocklistItemsListOperationOptions{}
}

func (o ConnectionRaiBlocklistItemsListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ConnectionRaiBlocklistItemsListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ConnectionRaiBlocklistItemsListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ProxyApiVersion != nil {
		out.Append("proxy-api-version", fmt.Sprintf("%v", *o.ProxyApiVersion))
	}
	return &out
}

type ConnectionRaiBlocklistItemsListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ConnectionRaiBlocklistItemsListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ConnectionRaiBlocklistItemsList ...
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistItemsList(ctx context.Context, id RaiBlocklistId, options ConnectionRaiBlocklistItemsListOperationOptions) (result ConnectionRaiBlocklistItemsListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ConnectionRaiBlocklistItemsListCustomPager{},
		Path:          fmt.Sprintf("%s/raiBlocklistItems", id.ID()),
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
		Values *[]RaiBlocklistItemPropertiesBasicResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ConnectionRaiBlocklistItemsListComplete retrieves all the results into a single object
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistItemsListComplete(ctx context.Context, id RaiBlocklistId, options ConnectionRaiBlocklistItemsListOperationOptions) (ConnectionRaiBlocklistItemsListCompleteResult, error) {
	return c.ConnectionRaiBlocklistItemsListCompleteMatchingPredicate(ctx, id, options, RaiBlocklistItemPropertiesBasicResourceOperationPredicate{})
}

// ConnectionRaiBlocklistItemsListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c V2WorkspaceConnectionResourceClient) ConnectionRaiBlocklistItemsListCompleteMatchingPredicate(ctx context.Context, id RaiBlocklistId, options ConnectionRaiBlocklistItemsListOperationOptions, predicate RaiBlocklistItemPropertiesBasicResourceOperationPredicate) (result ConnectionRaiBlocklistItemsListCompleteResult, err error) {
	items := make([]RaiBlocklistItemPropertiesBasicResource, 0)

	resp, err := c.ConnectionRaiBlocklistItemsList(ctx, id, options)
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

	result = ConnectionRaiBlocklistItemsListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
