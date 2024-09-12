package entitlementmanagementassignmentpolicycustomextensionstagesetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.CustomExtensionStageSetting
}

type GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingOperationOptions() GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingOperationOptions {
	return GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingOperationOptions{}
}

func (o GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAssignmentPolicyCustomExtensionStageSetting - Get customExtensionStageSettings from
// identityGovernance. The collection of stages when to execute one or more custom access package workflow extensions.
// Supports $expand.
func (c EntitlementManagementAssignmentPolicyCustomExtensionStageSettingClient) GetEntitlementManagementAssignmentPolicyCustomExtensionStageSetting(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId, options GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingOperationOptions) (result GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.CustomExtensionStageSetting
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
