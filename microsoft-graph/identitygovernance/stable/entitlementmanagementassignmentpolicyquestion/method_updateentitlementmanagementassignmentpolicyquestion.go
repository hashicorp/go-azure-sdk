package entitlementmanagementassignmentpolicyquestion

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateEntitlementManagementAssignmentPolicyQuestionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateEntitlementManagementAssignmentPolicyQuestionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateEntitlementManagementAssignmentPolicyQuestionOperationOptions() UpdateEntitlementManagementAssignmentPolicyQuestionOperationOptions {
	return UpdateEntitlementManagementAssignmentPolicyQuestionOperationOptions{}
}

func (o UpdateEntitlementManagementAssignmentPolicyQuestionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateEntitlementManagementAssignmentPolicyQuestionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateEntitlementManagementAssignmentPolicyQuestionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateEntitlementManagementAssignmentPolicyQuestion - Update the navigation property questions in identityGovernance
func (c EntitlementManagementAssignmentPolicyQuestionClient) UpdateEntitlementManagementAssignmentPolicyQuestion(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentPolicyIdQuestionId, input stable.AccessPackageQuestion, options UpdateEntitlementManagementAssignmentPolicyQuestionOperationOptions) (result UpdateEntitlementManagementAssignmentPolicyQuestionOperationResponse, err error) {
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
