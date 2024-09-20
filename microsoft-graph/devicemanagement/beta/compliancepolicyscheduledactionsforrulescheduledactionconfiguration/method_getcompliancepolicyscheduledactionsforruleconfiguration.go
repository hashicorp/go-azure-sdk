package compliancepolicyscheduledactionsforrulescheduledactionconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetCompliancePolicyScheduledActionsForRuleConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceManagementComplianceActionItem
}

type GetCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions() GetCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions {
	return GetCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions{}
}

func (o GetCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToOData() *odata.Query {
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

func (o GetCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetCompliancePolicyScheduledActionsForRuleConfiguration - Get scheduledActionConfigurations from deviceManagement.
// The list of scheduled action configurations for this compliance policy. This collection can contain a maximum of 100
// elements.
func (c CompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) GetCompliancePolicyScheduledActionsForRuleConfiguration(ctx context.Context, id beta.DeviceManagementCompliancePolicyIdScheduledActionsForRuleIdScheduledActionConfigurationId, options GetCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) (result GetCompliancePolicyScheduledActionsForRuleConfigurationOperationResponse, err error) {
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

	var model beta.DeviceManagementComplianceActionItem
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
