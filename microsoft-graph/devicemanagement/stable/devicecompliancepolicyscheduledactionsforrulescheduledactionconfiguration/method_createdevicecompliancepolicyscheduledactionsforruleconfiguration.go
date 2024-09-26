package devicecompliancepolicyscheduledactionsforrulescheduledactionconfiguration

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

type CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.DeviceComplianceActionItem
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

// CreateDeviceCompliancePolicyScheduledActionsForRuleConfiguration - Create deviceComplianceActionItem. Create a new
// deviceComplianceActionItem object.
func (c DeviceCompliancePolicyScheduledActionsForRuleScheduledActionConfigurationClient) CreateDeviceCompliancePolicyScheduledActionsForRuleConfiguration(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyIdScheduledActionsForRuleId, input stable.DeviceComplianceActionItem, options CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationOptions) (result CreateDeviceCompliancePolicyScheduledActionsForRuleConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
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

	var model stable.DeviceComplianceActionItem
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
