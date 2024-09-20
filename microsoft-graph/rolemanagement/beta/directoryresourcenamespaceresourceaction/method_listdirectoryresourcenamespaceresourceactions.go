package directoryresourcenamespaceresourceaction

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

type ListDirectoryResourceNamespaceResourceActionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRbacResourceAction
}

type ListDirectoryResourceNamespaceResourceActionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRbacResourceAction
}

type ListDirectoryResourceNamespaceResourceActionsOperationOptions struct {
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

func DefaultListDirectoryResourceNamespaceResourceActionsOperationOptions() ListDirectoryResourceNamespaceResourceActionsOperationOptions {
	return ListDirectoryResourceNamespaceResourceActionsOperationOptions{}
}

func (o ListDirectoryResourceNamespaceResourceActionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDirectoryResourceNamespaceResourceActionsOperationOptions) ToOData() *odata.Query {
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

func (o ListDirectoryResourceNamespaceResourceActionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDirectoryResourceNamespaceResourceActionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDirectoryResourceNamespaceResourceActionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDirectoryResourceNamespaceResourceActions - List resourceActions. Get a list of the unifiedRbacResourceAction
// objects and their properties.
func (c DirectoryResourceNamespaceResourceActionClient) ListDirectoryResourceNamespaceResourceActions(ctx context.Context, id beta.RoleManagementDirectoryResourceNamespaceId, options ListDirectoryResourceNamespaceResourceActionsOperationOptions) (result ListDirectoryResourceNamespaceResourceActionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDirectoryResourceNamespaceResourceActionsCustomPager{},
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

// ListDirectoryResourceNamespaceResourceActionsComplete retrieves all the results into a single object
func (c DirectoryResourceNamespaceResourceActionClient) ListDirectoryResourceNamespaceResourceActionsComplete(ctx context.Context, id beta.RoleManagementDirectoryResourceNamespaceId, options ListDirectoryResourceNamespaceResourceActionsOperationOptions) (ListDirectoryResourceNamespaceResourceActionsCompleteResult, error) {
	return c.ListDirectoryResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx, id, options, UnifiedRbacResourceActionOperationPredicate{})
}

// ListDirectoryResourceNamespaceResourceActionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DirectoryResourceNamespaceResourceActionClient) ListDirectoryResourceNamespaceResourceActionsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementDirectoryResourceNamespaceId, options ListDirectoryResourceNamespaceResourceActionsOperationOptions, predicate UnifiedRbacResourceActionOperationPredicate) (result ListDirectoryResourceNamespaceResourceActionsCompleteResult, err error) {
	items := make([]beta.UnifiedRbacResourceAction, 0)

	resp, err := c.ListDirectoryResourceNamespaceResourceActions(ctx, id, options)
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

	result = ListDirectoryResourceNamespaceResourceActionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
