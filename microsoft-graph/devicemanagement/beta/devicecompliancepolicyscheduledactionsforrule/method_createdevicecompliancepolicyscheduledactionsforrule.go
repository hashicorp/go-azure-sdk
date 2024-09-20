package devicecompliancepolicyscheduledactionsforrule

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

type CreateDeviceCompliancePolicyScheduledActionsForRuleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceComplianceScheduledActionForRule
}

type CreateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions() CreateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions {
	return CreateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions{}
}

func (o CreateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceCompliancePolicyScheduledActionsForRule - Create new navigation property to scheduledActionsForRule for
// deviceManagement
func (c DeviceCompliancePolicyScheduledActionsForRuleClient) CreateDeviceCompliancePolicyScheduledActionsForRule(ctx context.Context, id beta.DeviceManagementDeviceCompliancePolicyId, input beta.DeviceComplianceScheduledActionForRule, options CreateDeviceCompliancePolicyScheduledActionsForRuleOperationOptions) (result CreateDeviceCompliancePolicyScheduledActionsForRuleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/scheduledActionsForRule", id.ID()),
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

	var model beta.DeviceComplianceScheduledActionForRule
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
