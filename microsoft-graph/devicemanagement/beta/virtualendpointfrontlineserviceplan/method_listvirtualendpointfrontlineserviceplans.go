package virtualendpointfrontlineserviceplan

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

type ListVirtualEndpointFrontLineServicePlansOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCFrontLineServicePlan
}

type ListVirtualEndpointFrontLineServicePlansCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCFrontLineServicePlan
}

type ListVirtualEndpointFrontLineServicePlansOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListVirtualEndpointFrontLineServicePlansOperationOptions() ListVirtualEndpointFrontLineServicePlansOperationOptions {
	return ListVirtualEndpointFrontLineServicePlansOperationOptions{}
}

func (o ListVirtualEndpointFrontLineServicePlansOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVirtualEndpointFrontLineServicePlansOperationOptions) ToOData() *odata.Query {
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

func (o ListVirtualEndpointFrontLineServicePlansOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListVirtualEndpointFrontLineServicePlansCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointFrontLineServicePlansCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointFrontLineServicePlans - List cloudPcFrontLineServicePlans. Get a list of the
// cloudPcFrontLineServicePlan objects and their properties.
func (c VirtualEndpointFrontLineServicePlanClient) ListVirtualEndpointFrontLineServicePlans(ctx context.Context, options ListVirtualEndpointFrontLineServicePlansOperationOptions) (result ListVirtualEndpointFrontLineServicePlansOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListVirtualEndpointFrontLineServicePlansCustomPager{},
		Path:          "/deviceManagement/virtualEndpoint/frontLineServicePlans",
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
		Values *[]beta.CloudPCFrontLineServicePlan `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointFrontLineServicePlansComplete retrieves all the results into a single object
func (c VirtualEndpointFrontLineServicePlanClient) ListVirtualEndpointFrontLineServicePlansComplete(ctx context.Context, options ListVirtualEndpointFrontLineServicePlansOperationOptions) (ListVirtualEndpointFrontLineServicePlansCompleteResult, error) {
	return c.ListVirtualEndpointFrontLineServicePlansCompleteMatchingPredicate(ctx, options, CloudPCFrontLineServicePlanOperationPredicate{})
}

// ListVirtualEndpointFrontLineServicePlansCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointFrontLineServicePlanClient) ListVirtualEndpointFrontLineServicePlansCompleteMatchingPredicate(ctx context.Context, options ListVirtualEndpointFrontLineServicePlansOperationOptions, predicate CloudPCFrontLineServicePlanOperationPredicate) (result ListVirtualEndpointFrontLineServicePlansCompleteResult, err error) {
	items := make([]beta.CloudPCFrontLineServicePlan, 0)

	resp, err := c.ListVirtualEndpointFrontLineServicePlans(ctx, options)
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

	result = ListVirtualEndpointFrontLineServicePlansCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
