package entitlementmanagementroledefinitioninheritspermissionsfrom

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.UnifiedRoleDefinition
}

type ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.UnifiedRoleDefinition
}

type ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationOptions struct {
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

func DefaultListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationOptions() ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationOptions {
	return ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationOptions{}
}

func (o ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementRoleDefinitionInheritsPermissionsFroms - Get inheritsPermissionsFrom from roleManagement.
// Read-only collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in
// roles (isBuiltIn is true) support this attribute. Supports $expand.
func (c EntitlementManagementRoleDefinitionInheritsPermissionsFromClient) ListEntitlementManagementRoleDefinitionInheritsPermissionsFroms(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleDefinitionId, options ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationOptions) (result ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsCustomPager{},
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
		Values *[]stable.UnifiedRoleDefinition `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsComplete retrieves all the results into a single object
func (c EntitlementManagementRoleDefinitionInheritsPermissionsFromClient) ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsComplete(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleDefinitionId, options ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationOptions) (ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsCompleteResult, error) {
	return c.ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx, id, options, UnifiedRoleDefinitionOperationPredicate{})
}

// ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementRoleDefinitionInheritsPermissionsFromClient) ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx context.Context, id stable.RoleManagementEntitlementManagementRoleDefinitionId, options ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsOperationOptions, predicate UnifiedRoleDefinitionOperationPredicate) (result ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsCompleteResult, err error) {
	items := make([]stable.UnifiedRoleDefinition, 0)

	resp, err := c.ListEntitlementManagementRoleDefinitionInheritsPermissionsFroms(ctx, id, options)
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

	result = ListEntitlementManagementRoleDefinitionInheritsPermissionsFromsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
