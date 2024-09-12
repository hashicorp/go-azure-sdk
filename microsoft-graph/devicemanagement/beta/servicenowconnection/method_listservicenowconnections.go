package servicenowconnection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListServiceNowConnectionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.ServiceNowConnection
}

type ListServiceNowConnectionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.ServiceNowConnection
}

type ListServiceNowConnectionsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListServiceNowConnectionsOperationOptions() ListServiceNowConnectionsOperationOptions {
	return ListServiceNowConnectionsOperationOptions{}
}

func (o ListServiceNowConnectionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListServiceNowConnectionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Count != nil {
		out.Count = *o.Count
	}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.OrderBy != nil {
		out.OrderBy = *o.OrderBy
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	if o.Skip != nil {
		out.Skip = int(*o.Skip)
	}
	if o.Top != nil {
		out.Top = int(*o.Top)
	}
	return &out
}

func (o ListServiceNowConnectionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListServiceNowConnectionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListServiceNowConnectionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListServiceNowConnections - Get serviceNowConnections from deviceManagement. A list of ServiceNowConnections
func (c ServiceNowConnectionClient) ListServiceNowConnections(ctx context.Context, options ListServiceNowConnectionsOperationOptions) (result ListServiceNowConnectionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListServiceNowConnectionsCustomPager{},
		Path:          "/deviceManagement/serviceNowConnections",
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
		Values *[]beta.ServiceNowConnection `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListServiceNowConnectionsComplete retrieves all the results into a single object
func (c ServiceNowConnectionClient) ListServiceNowConnectionsComplete(ctx context.Context, options ListServiceNowConnectionsOperationOptions) (ListServiceNowConnectionsCompleteResult, error) {
	return c.ListServiceNowConnectionsCompleteMatchingPredicate(ctx, options, ServiceNowConnectionOperationPredicate{})
}

// ListServiceNowConnectionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ServiceNowConnectionClient) ListServiceNowConnectionsCompleteMatchingPredicate(ctx context.Context, options ListServiceNowConnectionsOperationOptions, predicate ServiceNowConnectionOperationPredicate) (result ListServiceNowConnectionsCompleteResult, err error) {
	items := make([]beta.ServiceNowConnection, 0)

	resp, err := c.ListServiceNowConnections(ctx, options)
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

	result = ListServiceNowConnectionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
