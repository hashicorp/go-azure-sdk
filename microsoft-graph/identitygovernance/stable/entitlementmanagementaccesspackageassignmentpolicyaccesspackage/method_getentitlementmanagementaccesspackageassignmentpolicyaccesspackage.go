package entitlementmanagementaccesspackageassignmentpolicyaccesspackage

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

type GetEntitlementManagementAccessPackageAssignmentPolicyAccessPackageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AccessPackage
}

type GetEntitlementManagementAccessPackageAssignmentPolicyAccessPackageOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEntitlementManagementAccessPackageAssignmentPolicyAccessPackageOperationOptions() GetEntitlementManagementAccessPackageAssignmentPolicyAccessPackageOperationOptions {
	return GetEntitlementManagementAccessPackageAssignmentPolicyAccessPackageOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageAssignmentPolicyAccessPackageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentPolicyAccessPackageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentPolicyAccessPackageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageAssignmentPolicyAccessPackage - Get accessPackage from identityGovernance.
// Access package containing this policy. Read-only. Supports $expand.
func (c EntitlementManagementAccessPackageAssignmentPolicyAccessPackageClient) GetEntitlementManagementAccessPackageAssignmentPolicyAccessPackage(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyId, options GetEntitlementManagementAccessPackageAssignmentPolicyAccessPackageOperationOptions) (result GetEntitlementManagementAccessPackageAssignmentPolicyAccessPackageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackage", id.ID()),
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

	var model stable.AccessPackage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
