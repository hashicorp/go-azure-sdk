package entitlementmanagementaccesspackageassignmentaccesspackageassignmentresourceroleaccesspackageresourcescopeaccesspackageresourceaccesspackageresourcerole

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

type ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AccessPackageResourceRole
}

type ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AccessPackageResourceRole
}

type ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationOptions struct {
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

func DefaultListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationOptions() ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationOptions {
	return ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationOptions{}
}

func (o ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationOptions) ToOData() *odata.Query {
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

func (o ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

type ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoles - Get
// accessPackageResourceRoles from identityGovernance. Read-only. Nullable. Supports $expand.
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient) ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoles(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentResourceRoleId, options ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationOptions) (result ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Pager:         &ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesCustomPager{},
		Path:          fmt.Sprintf("%s/accessPackageResourceScope/accessPackageResource/accessPackageResourceRoles", id.ID()),
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
		Values *[]beta.AccessPackageResourceRole `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesComplete retrieves all the results into a single object
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient) ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesComplete(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentResourceRoleId, options ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationOptions) (ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesCompleteResult, error) {
	return c.ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesCompleteMatchingPredicate(ctx, id, options, AccessPackageResourceRoleOperationPredicate{})
}

// ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c EntitlementManagementAccessPackageAssignmentAccessPackageAssignmentResourceRoleAccessPackageResourceScopeAccessPackageResourceAccessPackageResourceRoleClient) ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesCompleteMatchingPredicate(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentIdAccessPackageAssignmentResourceRoleId, options ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesOperationOptions, predicate AccessPackageResourceRoleOperationPredicate) (result ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesCompleteResult, err error) {
	items := make([]beta.AccessPackageResourceRole, 0)

	resp, err := c.ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRoles(ctx, id, options)
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

	result = ListEntitlementManagementAccessPackageAssignmentResourceRoleResourceScopeResourceResourceRolesCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
