package roleassignmentrolescopetag

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

type ListRoleAssignmentRoleScopeTagsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RoleScopeTag
}

type ListRoleAssignmentRoleScopeTagsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RoleScopeTag
}

type ListRoleAssignmentRoleScopeTagsOperationOptions struct {
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

func DefaultListRoleAssignmentRoleScopeTagsOperationOptions() ListRoleAssignmentRoleScopeTagsOperationOptions {
	return ListRoleAssignmentRoleScopeTagsOperationOptions{}
}

func (o ListRoleAssignmentRoleScopeTagsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRoleAssignmentRoleScopeTagsOperationOptions) ToOData() *odata.Query {
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

func (o ListRoleAssignmentRoleScopeTagsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRoleAssignmentRoleScopeTagsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleAssignmentRoleScopeTagsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleAssignmentRoleScopeTags - Get roleScopeTags from deviceManagement. Indicates the set of scope tags for the
// role assignment. These scope tags will limit the visibility of any Intune resources to those that match any of the
// scope tags in this collection.
func (c RoleAssignmentRoleScopeTagClient) ListRoleAssignmentRoleScopeTags(ctx context.Context, id beta.DeviceManagementRoleAssignmentId, options ListRoleAssignmentRoleScopeTagsOperationOptions) (result ListRoleAssignmentRoleScopeTagsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRoleAssignmentRoleScopeTagsCustomPager{},
		Path:          fmt.Sprintf("%s/roleScopeTags", id.ID()),
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
		Values *[]beta.RoleScopeTag `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleAssignmentRoleScopeTagsComplete retrieves all the results into a single object
func (c RoleAssignmentRoleScopeTagClient) ListRoleAssignmentRoleScopeTagsComplete(ctx context.Context, id beta.DeviceManagementRoleAssignmentId, options ListRoleAssignmentRoleScopeTagsOperationOptions) (ListRoleAssignmentRoleScopeTagsCompleteResult, error) {
	return c.ListRoleAssignmentRoleScopeTagsCompleteMatchingPredicate(ctx, id, options, RoleScopeTagOperationPredicate{})
}

// ListRoleAssignmentRoleScopeTagsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleAssignmentRoleScopeTagClient) ListRoleAssignmentRoleScopeTagsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementRoleAssignmentId, options ListRoleAssignmentRoleScopeTagsOperationOptions, predicate RoleScopeTagOperationPredicate) (result ListRoleAssignmentRoleScopeTagsCompleteResult, err error) {
	items := make([]beta.RoleScopeTag, 0)

	resp, err := c.ListRoleAssignmentRoleScopeTags(ctx, id, options)
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

	result = ListRoleAssignmentRoleScopeTagsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
