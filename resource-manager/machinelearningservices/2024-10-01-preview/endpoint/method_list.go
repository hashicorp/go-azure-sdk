package endpoint

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]EndpointResourcePropertiesBasicResource
}

type ListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []EndpointResourcePropertiesBasicResource
}

type ListOperationOptions struct {
	EndpointType               *EndpointType
	Expand                     *string
	IncludeConnections         *bool
	IncludeOnlineEndpoints     *bool
	IncludeServerlessEndpoints *bool
	Skip                       *string
}

func DefaultListOperationOptions() ListOperationOptions {
	return ListOperationOptions{}
}

func (o ListOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ListOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.EndpointType != nil {
		out.Append("endpointType", fmt.Sprintf("%v", *o.EndpointType))
	}
	if o.Expand != nil {
		out.Append("$expand", fmt.Sprintf("%v", *o.Expand))
	}
	if o.IncludeConnections != nil {
		out.Append("includeConnections", fmt.Sprintf("%v", *o.IncludeConnections))
	}
	if o.IncludeOnlineEndpoints != nil {
		out.Append("includeOnlineEndpoints", fmt.Sprintf("%v", *o.IncludeOnlineEndpoints))
	}
	if o.IncludeServerlessEndpoints != nil {
		out.Append("includeServerlessEndpoints", fmt.Sprintf("%v", *o.IncludeServerlessEndpoints))
	}
	if o.Skip != nil {
		out.Append("$skip", fmt.Sprintf("%v", *o.Skip))
	}
	return &out
}

type ListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// List ...
func (c EndpointClient) List(ctx context.Context, id WorkspaceId, options ListOperationOptions) (result ListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCustomPager{},
		Path:          fmt.Sprintf("%s/endpoints", id.ID()),
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
		Values *[]EndpointResourcePropertiesBasicResource `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListComplete retrieves all the results into a single object
func (c EndpointClient) ListComplete(ctx context.Context, id WorkspaceId, options ListOperationOptions) (ListCompleteResult, error) {
	return c.ListCompleteMatchingPredicate(ctx, id, options, EndpointResourcePropertiesBasicResourceOperationPredicate{})
}

// ListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EndpointClient) ListCompleteMatchingPredicate(ctx context.Context, id WorkspaceId, options ListOperationOptions, predicate EndpointResourcePropertiesBasicResourceOperationPredicate) (result ListCompleteResult, err error) {
	items := make([]EndpointResourcePropertiesBasicResource, 0)

	resp, err := c.List(ctx, id, options)
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

	result = ListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
