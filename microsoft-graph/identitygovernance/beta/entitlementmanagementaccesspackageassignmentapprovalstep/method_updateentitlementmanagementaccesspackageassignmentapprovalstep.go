package entitlementmanagementaccesspackageassignmentapprovalstep

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementAccessPackageAssignmentApprovalStepOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageAssignmentApprovalStepOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementAccessPackageAssignmentApprovalStepOperationOptions() UpdateEntitlementManagementAccessPackageAssignmentApprovalStepOperationOptions {
	return UpdateEntitlementManagementAccessPackageAssignmentApprovalStepOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageAssignmentApprovalStepOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentApprovalStepOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentApprovalStepOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageAssignmentApprovalStep - Update approvalStep. Apply approve or deny decision
// on an approvalStep object.
func (c EntitlementManagementAccessPackageAssignmentApprovalStepClient) UpdateEntitlementManagementAccessPackageAssignmentApprovalStep(ctx context.Context, id beta.IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalIdStepId, input beta.ApprovalStep, options UpdateEntitlementManagementAccessPackageAssignmentApprovalStepOperationOptions) (result UpdateEntitlementManagementAccessPackageAssignmentApprovalStepOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
