package exchangeresourcenamespaceresourceaction

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

type ListExchangeResourceNamespaceResourceActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRbacResourceAction
}

type ListExchangeResourceNamespaceResourceActionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRbacResourceAction
}

type ListExchangeResourceNamespaceResourceActionsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListExchangeResourceNamespaceResourceActionsOperationOptions() ListExchangeResourceNamespaceResourceActionsOperationOptions {
	return ListExchangeResourceNamespaceResourceActionsOperationOptions{}
}

func (o ListExchangeResourceNamespaceResourceActionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListExchangeResourceNamespaceResourceActionsOperationOptions) ToOData() *odata.Query {
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

func (o ListExchangeResourceNamespaceResourceActionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListExchangeResourceNamespaceResourceActionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListExchangeResourceNamespaceResourceActionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListExchangeResourceNamespaceResourceActions - Get resourceActions from roleManagement. Operations that an authorized
// principal is allowed to perform.
func (c ExchangeResourceNamespaceResourceActionClient) ListExchangeResourceNamespaceResourceActions(ctx context.Context, id beta.RoleManagementExchangeResourceNamespaceId, options ListExchangeResourceNamespaceResourceActionsOperationOptions) (result ListExchangeResourceNamespaceResourceActionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListExchangeResourceNamespaceResourceActionsCustomPager{},
		Path:          fmt.Sprintf("%s/resourceActions", id.ID()),
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
		Values *[]beta.UnifiedRbacResourceAction `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListExchangeResourceNamespaceResourceActionsComplete retrieves all the results into a single object
func (c ExchangeResourceNamespaceResourceActionClient) ListExchangeResourceNamespaceResourceActionsComplete(ctx context.Context, id beta.RoleManagementExchangeResourceNamespaceId, options ListExchangeResourceNamespaceResourceActionsOperationOptions) (ListExchangeResourceNamespaceResourceActionsCompleteResult, error) {
	return c.ListExchangeResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx, id, options, UnifiedRbacResourceActionOperationPredicate{})
}

// ListExchangeResourceNamespaceResourceActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ExchangeResourceNamespaceResourceActionClient) ListExchangeResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementExchangeResourceNamespaceId, options ListExchangeResourceNamespaceResourceActionsOperationOptions, predicate UnifiedRbacResourceActionOperationPredicate) (result ListExchangeResourceNamespaceResourceActionsCompleteResult, err error) {
	items := make([]beta.UnifiedRbacResourceAction, 0)

	resp, err := c.ListExchangeResourceNamespaceResourceActions(ctx, id, options)
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

	result = ListExchangeResourceNamespaceResourceActionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
