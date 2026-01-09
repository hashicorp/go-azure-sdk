package virtualendpointserviceplan

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListVirtualEndpointServicePlansOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCServicePlan
}

type ListVirtualEndpointServicePlansCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCServicePlan
}

type ListVirtualEndpointServicePlansOperationOptions struct {
	Count     *bool
	Expand    *odata.Expand
	Filter    *string
	Metadata  *odata.Metadata
	OrderBy   *odata.OrderBy
	RetryFunc client.RequestRetryFunc
	Search    *string
	Select    *[]string
	Skip      *int64
	Top       *int64
}

func DefaultListVirtualEndpointServicePlansOperationOptions() ListVirtualEndpointServicePlansOperationOptions {
	return ListVirtualEndpointServicePlansOperationOptions{}
}

func (o ListVirtualEndpointServicePlansOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVirtualEndpointServicePlansOperationOptions) ToOData() *odata.Query {
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
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
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

func (o ListVirtualEndpointServicePlansOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListVirtualEndpointServicePlansCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointServicePlansCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointServicePlans - List servicePlans. List the currently available service plans that an organization
// can purchase for their Cloud PCs. For examples of currently available service plans, see Windows 365 compare plans
// and pricing. Currently, Microsoft Graph API is available for Windows 365 Enterprise.
func (c VirtualEndpointServicePlanClient) ListVirtualEndpointServicePlans(ctx context.Context, options ListVirtualEndpointServicePlansOperationOptions) (result ListVirtualEndpointServicePlansOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListVirtualEndpointServicePlansCustomPager{},
		Path:          "/deviceManagement/virtualEndpoint/servicePlans",
		RetryFunc:     options.RetryFunc,
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
		Values *[]beta.CloudPCServicePlan `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointServicePlansComplete retrieves all the results into a single object
func (c VirtualEndpointServicePlanClient) ListVirtualEndpointServicePlansComplete(ctx context.Context, options ListVirtualEndpointServicePlansOperationOptions) (ListVirtualEndpointServicePlansCompleteResult, error) {
	return c.ListVirtualEndpointServicePlansCompleteMatchingPredicate(ctx, options, CloudPCServicePlanOperationPredicate{})
}

// ListVirtualEndpointServicePlansCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointServicePlanClient) ListVirtualEndpointServicePlansCompleteMatchingPredicate(ctx context.Context, options ListVirtualEndpointServicePlansOperationOptions, predicate CloudPCServicePlanOperationPredicate) (result ListVirtualEndpointServicePlansCompleteResult, err error) {
	items := make([]beta.CloudPCServicePlan, 0)

	resp, err := c.ListVirtualEndpointServicePlans(ctx, options)
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

	result = ListVirtualEndpointServicePlansCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
