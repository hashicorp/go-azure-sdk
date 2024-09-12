package entitlementmanagementaccesspackageassignmentresourceroleaccesspackagesubject

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

type GetEntitlementManagementAccessPackageAssignmentResourceRoleSubjectOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageSubject
}

type GetEntitlementManagementAccessPackageAssignmentResourceRoleSubjectOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEntitlementManagementAccessPackageAssignmentResourceRoleSubjectOperationOptions() GetEntitlementManagementAccessPackageAssignmentResourceRoleSubjectOperationOptions {
	return GetEntitlementManagementAccessPackageAssignmentResourceRoleSubjectOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageAssignmentResourceRoleSubjectOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentResourceRoleSubjectOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentResourceRoleSubjectOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageAssignmentResourceRoleSubject - Get accessPackageSubject from
// identityGovernance. Read-only. Nullable. Supports $filter (eq) on objectId and $expand query parameters.
func (c EntitlementManagementAccessPackageAssignmentResourceRoleAccessPackageSubjectClient) GetEntitlementManagementAccessPackageAssignmentResourceRoleSubject(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentResourceRoleId, options GetEntitlementManagementAccessPackageAssignmentResourceRoleSubjectOperationOptions) (result GetEntitlementManagementAccessPackageAssignmentResourceRoleSubjectOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageSubject", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var model beta.AccessPackageSubject
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
