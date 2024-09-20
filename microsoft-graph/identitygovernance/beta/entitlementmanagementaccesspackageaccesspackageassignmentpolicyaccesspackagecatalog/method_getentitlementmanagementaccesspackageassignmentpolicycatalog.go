package entitlementmanagementaccesspackageaccesspackageassignmentpolicyaccesspackagecatalog

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

type GetEntitlementManagementAccessPackageAssignmentPolicyCatalogOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.AccessPackageCatalog
}

type GetEntitlementManagementAccessPackageAssignmentPolicyCatalogOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetEntitlementManagementAccessPackageAssignmentPolicyCatalogOperationOptions() GetEntitlementManagementAccessPackageAssignmentPolicyCatalogOperationOptions {
	return GetEntitlementManagementAccessPackageAssignmentPolicyCatalogOperationOptions{}
}

func (o GetEntitlementManagementAccessPackageAssignmentPolicyCatalogOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentPolicyCatalogOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAccessPackageAssignmentPolicyCatalogOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAccessPackageAssignmentPolicyCatalog - Get accessPackageCatalog from identityGovernance
func (c EntitlementManagementAccessPackageAccessPackageAssignmentPolicyAccessPackageCatalogClient) GetEntitlementManagementAccessPackageAssignmentPolicyCatalog(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageIdAccessPackageAssignmentPolicyId, options GetEntitlementManagementAccessPackageAssignmentPolicyCatalogOperationOptions) (result GetEntitlementManagementAccessPackageAssignmentPolicyCatalogOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackageCatalog", id.ID()),
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

	var model beta.AccessPackageCatalog
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
