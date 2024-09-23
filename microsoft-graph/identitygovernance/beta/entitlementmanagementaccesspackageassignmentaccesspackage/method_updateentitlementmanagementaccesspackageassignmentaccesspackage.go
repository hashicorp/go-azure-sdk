package entitlementmanagementaccesspackageassignmentaccesspackage

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

type UpdateEntitlementManagementAccessPackageAssignmentAccessPackageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageAssignmentAccessPackageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementAccessPackageAssignmentAccessPackageOperationOptions() UpdateEntitlementManagementAccessPackageAssignmentAccessPackageOperationOptions {
	return UpdateEntitlementManagementAccessPackageAssignmentAccessPackageOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageAssignmentAccessPackageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentAccessPackageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentAccessPackageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageAssignmentAccessPackage - Update the navigation property accessPackage in
// identityGovernance
func (c EntitlementManagementAccessPackageAssignmentAccessPackageClient) UpdateEntitlementManagementAccessPackageAssignmentAccessPackage(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, input beta.AccessPackage, options UpdateEntitlementManagementAccessPackageAssignmentAccessPackageOperationOptions) (result UpdateEntitlementManagementAccessPackageAssignmentAccessPackageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/accessPackage", id.ID()),
		RetryFunc:     options.RetryFunc,
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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

	return
}
