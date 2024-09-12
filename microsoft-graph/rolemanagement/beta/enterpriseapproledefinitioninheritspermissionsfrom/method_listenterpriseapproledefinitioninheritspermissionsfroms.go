package enterpriseapproledefinitioninheritspermissionsfrom

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

type ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.UnifiedRoleDefinition
}

type ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.UnifiedRoleDefinition
}

type ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationOptions struct {
	Count   *bool
	Expand  *odata.Expand
	Filter  *string
	OrderBy *odata.OrderBy
	Search  *string
	Select  *[]string
	Skip    *int64
	Top     *int64
}

func DefaultListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationOptions() ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationOptions {
	return ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationOptions{}
}

func (o ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationOptions) ToOData() *odata.Query {
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

func (o ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEnterpriseAppRoleDefinitionInheritsPermissionsFroms - Get inheritsPermissionsFrom from roleManagement. Read-only
// collection of role definitions that the given role definition inherits from. Only Microsoft Entra built-in roles
// support this attribute.
func (c EnterpriseAppRoleDefinitionInheritsPermissionsFromClient) ListEnterpriseAppRoleDefinitionInheritsPermissionsFroms(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleDefinitionId, options ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationOptions) (result ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsCustomPager{},
		Path:          fmt.Sprintf("%s/inheritsPermissionsFrom", id.ID()),
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

// ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsComplete retrieves all the results into a single object
func (c EnterpriseAppRoleDefinitionInheritsPermissionsFromClient) ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsComplete(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleDefinitionId, options ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationOptions) (ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsCompleteResult, error) {
	return c.ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx, id, options, UnifiedRoleDefinitionOperationPredicate{})
}

// ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EnterpriseAppRoleDefinitionInheritsPermissionsFromClient) ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsCompleteMatchingPredicate(ctx context.Context, id beta.RoleManagementEnterpriseAppIdRoleDefinitionId, options ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsOperationOptions, predicate UnifiedRoleDefinitionOperationPredicate) (result ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsCompleteResult, err error) {
	items := make([]beta.UnifiedRoleDefinition, 0)

	resp, err := c.ListEnterpriseAppRoleDefinitionInheritsPermissionsFroms(ctx, id, options)
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

	result = ListEnterpriseAppRoleDefinitionInheritsPermissionsFromsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
