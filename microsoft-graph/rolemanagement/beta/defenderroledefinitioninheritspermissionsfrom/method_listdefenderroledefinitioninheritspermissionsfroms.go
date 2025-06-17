package defenderroledefinitioninheritspermissionsfrom

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

type ListDefenderRoleDefinitionInheritsPermissionsFromsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleDefinition
}

type ListDefenderRoleDefinitionInheritsPermissionsFromsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleDefinition
}

type ListDefenderRoleDefinitionInheritsPermissionsFromsOperationOptions struct {
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

func DefaultListDefenderRoleDefinitionInheritsPermissionsFromsOperationOptions() ListDefenderRoleDefinitionInheritsPermissionsFromsOperationOptions {
	return ListDefenderRoleDefinitionInheritsPermissionsFromsOperationOptions{}
}

func (o ListDefenderRoleDefinitionInheritsPermissionsFromsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListDefenderRoleDefinitionInheritsPermissionsFromsOperationOptions) ToOData() *odata.Query {
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

func (o ListDefenderRoleDefinitionInheritsPermissionsFromsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListDefenderRoleDefinitionInheritsPermissionsFromsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListDefenderRoleDefinitionInheritsPermissionsFromsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListDefenderRoleDefinitionInheritsPermissionsFroms - Get inheritsPermissionsFrom from roleManagement. Read-only
// collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles
// support this attribute.
func (c DefenderRoleDefinitionInheritsPermissionsFromClient) ListDefenderRoleDefinitionInheritsPermissionsFroms(ctx context.Context, id beta.RoleManagementDefenderRoleDefinitionId, options ListDefenderRoleDefinitionInheritsPermissionsFromsOperationOptions) (result ListDefenderRoleDefinitionInheritsPermissionsFromsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListDefenderRoleDefinitionInheritsPermissionsFromsCustomPager{},
		Path:          fmt.Sprintf("%s/inheritsPermissionsFrom", id.ID()),
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
		Values *[]beta.UnifiedRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListDefenderRoleDefinitionInheritsPermissionsFromsComplete retrieves all the results into a single object
func (c DefenderRoleDefinitionInheritsPermissionsFromClient) ListDefenderRoleDefinitionInheritsPermissionsFromsComplete(ctx context.Context, id beta.RoleManagementDefenderRoleDefinitionId, options ListDefenderRoleDefinitionInheritsPermissionsFromsOperationOptions) (ListDefenderRoleDefinitionInheritsPermissionsFromsCompleteResult, error) {
	return c.ListDefenderRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx, id, options, UnifiedRoleDefinitionOperationPredicate{})
}

// ListDefenderRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c DefenderRoleDefinitionInheritsPermissionsFromClient) ListDefenderRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementDefenderRoleDefinitionId, options ListDefenderRoleDefinitionInheritsPermissionsFromsOperationOptions, predicate UnifiedRoleDefinitionOperationPredicate) (result ListDefenderRoleDefinitionInheritsPermissionsFromsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleDefinition, 0)

	resp, err := c.ListDefenderRoleDefinitionInheritsPermissionsFroms(ctx, id, options)
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

	result = ListDefenderRoleDefinitionInheritsPermissionsFromsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
