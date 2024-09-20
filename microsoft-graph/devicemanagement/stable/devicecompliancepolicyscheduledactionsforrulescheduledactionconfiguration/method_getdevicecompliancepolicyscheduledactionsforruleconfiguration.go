package devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DeviceComplianceActionItem
}

type GetDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions() GetDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions {
	return GetDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions{}
}

func (o GetDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToOData() *odata.Query {
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

func (o GetDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceCompliancePolicyScheduledActionsForRuleConfiguration - Get deviceComplianceActionItem. Read properties and
// relationships of the deviceComplianceActionItem object.
func (c DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) GetDeviceCompliancePolicyScheduledActionsForRuleConfiguration(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId, options GetDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) (result GetDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationResponse, err error) {
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

	var model stable.DeviceComplianceActionItem
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
