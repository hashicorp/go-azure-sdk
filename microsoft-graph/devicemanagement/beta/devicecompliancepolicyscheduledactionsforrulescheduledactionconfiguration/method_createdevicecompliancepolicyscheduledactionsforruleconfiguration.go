package devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration

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

type CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceComplianceActionItem
}

type CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions() CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions {
	return CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions{}
}

func (o CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceCompliancePolicyScheduledActionsForRuleConfiguration - Create new navigation property to
// scheduledActionConfigurations for deviceManagement
func (c DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) CreateDeviceCompliancePolicyScheduledActionsForRuleConfiguration(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId, input beta.DeviceComplianceActionItem, options CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) (result CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/scheduledActionConfigurations", id.ID()),
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

	var model beta.DeviceComplianceActionItem
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
