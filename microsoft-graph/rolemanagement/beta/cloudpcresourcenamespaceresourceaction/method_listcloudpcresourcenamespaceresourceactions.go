package cloudpcresourcenamespaceresourceaction

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

type ListCloudPCResourceNamespaceResourceActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRbacResourceAction
}

type ListCloudPCResourceNamespaceResourceActionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRbacResourceAction
}

type ListCloudPCResourceNamespaceResourceActionsOperationOptions struct {
	Count    *bool
	Expand   *odata.Expand
	Filter   *string
	Metadata *odata.Metadata
	OrderBy  *odata.OrderBy
	Search   *string
	Select   *[]string
	Skip     *int64
	Top      *int64
}

func DefaultListCloudPCResourceNamespaceResourceActionsOperationOptions() ListCloudPCResourceNamespaceResourceActionsOperationOptions {
	return ListCloudPCResourceNamespaceResourceActionsOperationOptions{}
}

func (o ListCloudPCResourceNamespaceResourceActionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListCloudPCResourceNamespaceResourceActionsOperationOptions) ToOData() *odata.Query {
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

func (o ListCloudPCResourceNamespaceResourceActionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListCloudPCResourceNamespaceResourceActionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListCloudPCResourceNamespaceResourceActionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListCloudPCResourceNamespaceResourceActions - Get resourceActions from roleManagement. Operations that an authorized
// principal is allowed to perform.
func (c CloudPCResourceNamespaceResourceActionClient) ListCloudPCResourceNamespaceResourceActions(ctx context.Context, id beta.RoleManagementCloudPCResourceNamespaceId, options ListCloudPCResourceNamespaceResourceActionsOperationOptions) (result ListCloudPCResourceNamespaceResourceActionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListCloudPCResourceNamespaceResourceActionsCustomPager{},
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

// ListCloudPCResourceNamespaceResourceActionsComplete retrieves all the results into a single object
func (c CloudPCResourceNamespaceResourceActionClient) ListCloudPCResourceNamespaceResourceActionsComplete(ctx context.Context, id beta.RoleManagementCloudPCResourceNamespaceId, options ListCloudPCResourceNamespaceResourceActionsOperationOptions) (ListCloudPCResourceNamespaceResourceActionsCompleteResult, error) {
	return c.ListCloudPCResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx, id, options, UnifiedRbacResourceActionOperationPredicate{})
}

// ListCloudPCResourceNamespaceResourceActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c CloudPCResourceNamespaceResourceActionClient) ListCloudPCResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementCloudPCResourceNamespaceId, options ListCloudPCResourceNamespaceResourceActionsOperationOptions, predicate UnifiedRbacResourceActionOperationPredicate) (result ListCloudPCResourceNamespaceResourceActionsCompleteResult, err error) {
	items := make([]beta.UnifiedRbacResourceAction, 0)

	resp, err := c.ListCloudPCResourceNamespaceResourceActions(ctx, id, options)
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

	result = ListCloudPCResourceNamespaceResourceActionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
