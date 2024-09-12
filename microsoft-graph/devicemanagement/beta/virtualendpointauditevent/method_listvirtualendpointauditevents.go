package virtualendpointauditevent

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

type ListVirtualEndpointAuditEventsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.CloudPCAuditEvent
}

type ListVirtualEndpointAuditEventsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.CloudPCAuditEvent
}

type ListVirtualEndpointAuditEventsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListVirtualEndpointAuditEventsOperationOptions() ListVirtualEndpointAuditEventsOperationOptions {
	return ListVirtualEndpointAuditEventsOperationOptions{}
}

func (o ListVirtualEndpointAuditEventsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListVirtualEndpointAuditEventsOperationOptions) ToOData() *odata.Query {
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

func (o ListVirtualEndpointAuditEventsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListVirtualEndpointAuditEventsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListVirtualEndpointAuditEventsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListVirtualEndpointAuditEvents - List auditEvents. List all the cloudPcAuditEvent objects for the tenant.
func (c VirtualEndpointAuditEventClient) ListVirtualEndpointAuditEvents(ctx context.Context, options ListVirtualEndpointAuditEventsOperationOptions) (result ListVirtualEndpointAuditEventsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListVirtualEndpointAuditEventsCustomPager{},
		Path:          "/deviceManagement/virtualEndpoint/auditEvents",
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
		Values *[]beta.CloudPCAuditEvent `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListVirtualEndpointAuditEventsComplete retrieves all the results into a single object
func (c VirtualEndpointAuditEventClient) ListVirtualEndpointAuditEventsComplete(ctx context.Context, options ListVirtualEndpointAuditEventsOperationOptions) (ListVirtualEndpointAuditEventsCompleteResult, error) {
	return c.ListVirtualEndpointAuditEventsCompleteMatchingPredicate(ctx, options, CloudPCAuditEventOperationPredicate{})
}

// ListVirtualEndpointAuditEventsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c VirtualEndpointAuditEventClient) ListVirtualEndpointAuditEventsCompleteMatchingPredicate(ctx context.Context, options ListVirtualEndpointAuditEventsOperationOptions, predicate CloudPCAuditEventOperationPredicate) (result ListVirtualEndpointAuditEventsCompleteResult, err error) {
	items := make([]beta.CloudPCAuditEvent, 0)

	resp, err := c.ListVirtualEndpointAuditEvents(ctx, options)
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

	result = ListVirtualEndpointAuditEventsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
