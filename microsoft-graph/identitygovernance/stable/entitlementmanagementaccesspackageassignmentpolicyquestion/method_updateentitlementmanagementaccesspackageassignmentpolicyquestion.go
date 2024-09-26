package entitlementmanagementaccesspackageassignmentpolicyquestion

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions() UpdateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions {
	return UpdateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions{}
}

func (o UpdateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAccessPackageAssignmentPolicyQuestion - Update the navigation property questions in
// identityGovernance
func (c EntitlementManagementAccessPackageAssignmentPolicyQuestionClient) UpdateEntitlementManagementAccessPackageAssignmentPolicyQuestion(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAccessPackageIdAssignmentPolicyIdQuestionId, input stable.AccessPackageQuestion, options UpdateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationOptions) (result UpdateEntitlementManagementAccessPackageAssignmentPolicyQuestionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
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
