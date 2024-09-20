package entitlementmanagementaccesspackageassignmenttarget

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

type UpdateEntitlementManagementAccessPackageAssignmentTargetOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageAssignmentTargetOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateEntitlementManagementAccessPackageAssignmentTargetOperationOptions() UpdateEntitlementManagementAccessPackageAssignmentTargetOperationOptions {
	return UpdateEntitlementManagementAccessPackageAssignmentTargetOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageAssignmentTargetOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentTargetOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentTargetOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageAssignmentTarget - Update the navigation property target in
// identityGovernance
func (c EntitlementManagementAccessPackageAssignmentTargetClient) UpdateEntitlementManagementAccessPackageAssignmentTarget(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentId, input beta.AccessPackageSubject, options UpdateEntitlementManagementAccessPackageAssignmentTargetOperationOptions) (result UpdateEntitlementManagementAccessPackageAssignmentTargetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/target", id.ID()),
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
