package entitlementmanagementaccesspackageassignmentapprovalstage

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

type CreateEntitlementManagementAccessPackageAssignmentApprovalStageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ApprovalStage
}

type CreateEntitlementManagementAccessPackageAssignmentApprovalStageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateEntitlementManagementAccessPackageAssignmentApprovalStageOperationOptions() CreateEntitlementManagementAccessPackageAssignmentApprovalStageOperationOptions {
	return CreateEntitlementManagementAccessPackageAssignmentApprovalStageOperationOptions{}
}

func (o CreateEntitlementManagementAccessPackageAssignmentApprovalStageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentApprovalStageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateEntitlementManagementAccessPackageAssignmentApprovalStageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateEntitlementManagementAccessPackageAssignmentApprovalStage - Create new navigation property to stages for
// identityGovernance
func (c EntitlementManagementAccessPackageAssignmentApprovalStageClient) CreateEntitlementManagementAccessPackageAssignmentApprovalStage(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageAssignmentApprovalId, input stable.ApprovalStage, options CreateEntitlementManagementAccessPackageAssignmentApprovalStageOperationOptions) (result CreateEntitlementManagementAccessPackageAssignmentApprovalStageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/stages", id.ID()),
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

	var model stable.ApprovalStage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
