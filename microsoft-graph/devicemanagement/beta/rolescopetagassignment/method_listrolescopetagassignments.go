package rolescopetagassignment

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

type ListRoleScopeTagAssignmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.RoleScopeTagAutoAssignment
}

type ListRoleScopeTagAssignmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.RoleScopeTagAutoAssignment
}

type ListRoleScopeTagAssignmentsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListRoleScopeTagAssignmentsOperationOptions() ListRoleScopeTagAssignmentsOperationOptions {
	return ListRoleScopeTagAssignmentsOperationOptions{}
}

func (o ListRoleScopeTagAssignmentsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListRoleScopeTagAssignmentsOperationOptions) ToOData() *odata.Query {
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

func (o ListRoleScopeTagAssignmentsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListRoleScopeTagAssignmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRoleScopeTagAssignmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRoleScopeTagAssignments - Get assignments from deviceManagement. The list of assignments for this Role Scope Tag.
func (c RoleScopeTagAssignmentClient) ListRoleScopeTagAssignments(ctx context.Context, id beta.DeviceManagementRoleScopeTagId, options ListRoleScopeTagAssignmentsOperationOptions) (result ListRoleScopeTagAssignmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListRoleScopeTagAssignmentsCustomPager{},
		Path:          fmt.Sprintf("%s/assignments", id.ID()),
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
		Values *[]beta.RoleScopeTagAutoAssignment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRoleScopeTagAssignmentsComplete retrieves all the results into a single object
func (c RoleScopeTagAssignmentClient) ListRoleScopeTagAssignmentsComplete(ctx context.Context, id beta.DeviceManagementRoleScopeTagId, options ListRoleScopeTagAssignmentsOperationOptions) (ListRoleScopeTagAssignmentsCompleteResult, error) {
	return c.ListRoleScopeTagAssignmentsCompleteMatchingPredicate(ctx, id, options, RoleScopeTagAutoAssignmentOperationPredicate{})
}

// ListRoleScopeTagAssignmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RoleScopeTagAssignmentClient) ListRoleScopeTagAssignmentsCompleteMatchingPredicate(ctx context.Context, id beta.DeviceManagementRoleScopeTagId, options ListRoleScopeTagAssignmentsOperationOptions, predicate RoleScopeTagAutoAssignmentOperationPredicate) (result ListRoleScopeTagAssignmentsCompleteResult, err error) {
	items := make([]beta.RoleScopeTagAutoAssignment, 0)

	resp, err := c.ListRoleScopeTagAssignments(ctx, id, options)
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

	result = ListRoleScopeTagAssignmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
