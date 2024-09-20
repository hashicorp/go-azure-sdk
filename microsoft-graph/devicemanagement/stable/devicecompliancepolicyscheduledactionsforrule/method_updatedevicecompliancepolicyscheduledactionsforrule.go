package devicecompliancepolicyscheduledactionsforrule

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceCompliancePolicyScheduledActionsForRuleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions() UpdateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions {
	return UpdateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions{}
}

func (o UpdateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceCompliancePolicyScheduledActionsForRule - Update deviceComplianceScheduledActionForRule. Update the
// properties of a deviceComplianceScheduledActionForRule object.
func (c DeviceCompliancePolicyScheduledActionsForRuleClient) UpdateDeviceCompliancePolicyScheduledActionsForRule(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId, input stable.DeviceComplianceScheduledActionForRule, options UpdateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions) (result UpdateDeviceCompliancePolicyScheduledActionsForRuleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
