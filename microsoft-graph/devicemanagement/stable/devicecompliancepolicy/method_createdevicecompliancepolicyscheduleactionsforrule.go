package devicecompliancepolicy

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

type CreateDeviceCompliancePolicyScheduleActionsForRuleOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateDeviceCompliancePolicyScheduleActionsForRuleOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDeviceCompliancePolicyScheduleActionsForRuleOperationOptions() CreateDeviceCompliancePolicyScheduleActionsForRuleOperationOptions {
	return CreateDeviceCompliancePolicyScheduleActionsForRuleOperationOptions{}
}

func (o CreateDeviceCompliancePolicyScheduleActionsForRuleOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceCompliancePolicyScheduleActionsForRuleOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceCompliancePolicyScheduleActionsForRuleOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceCompliancePolicyScheduleActionsForRule - Invoke action scheduleActionsForRules. Not yet documented
func (c DeviceCompliancePolicyClient) CreateDeviceCompliancePolicyScheduleActionsForRule(ctx context.Context, id stable.DeviceManagementDeviceCompliancePolicyId, input CreateDeviceCompliancePolicyScheduleActionsForRuleRequest, options CreateDeviceCompliancePolicyScheduleActionsForRuleOperationOptions) (result CreateDeviceCompliancePolicyScheduleActionsForRuleOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/scheduleActionsForRules", id.ID()),
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
