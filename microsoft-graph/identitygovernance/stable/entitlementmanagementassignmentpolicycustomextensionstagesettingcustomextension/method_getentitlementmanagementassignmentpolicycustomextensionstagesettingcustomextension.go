package entitlementmanagementassignmentpolicycustomextensionstagesettingcustomextension

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        stable.CustomCalloutExtension
}

type GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions() GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions {
	return GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions{}
}

func (o GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions) ToOData() *odata.Query {
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

func (o GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtension - Get customExtension from
// identityGovernance. Indicates the custom workflow extension that will be executed at this stage. Nullable. Supports
// $expand.
func (c EntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionClient) GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtension(ctx context.Context, id stable.IdentityGovernanceEntitlementManagementAssignmentPolicyIdCustomExtensionStageSettingId, options GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationOptions) (result GetEntitlementManagementAssignmentPolicyCustomExtensionStageSettingCustomExtensionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/customExtension", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := stable.UnmarshalCustomCalloutExtensionImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
