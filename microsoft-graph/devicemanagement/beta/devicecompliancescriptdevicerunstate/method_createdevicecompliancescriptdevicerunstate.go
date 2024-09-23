package devicecompliancescriptdevicerunstate

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

type CreateDeviceComplianceScriptDeviceRunStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceComplianceScriptDeviceState
}

type CreateDeviceComplianceScriptDeviceRunStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDeviceComplianceScriptDeviceRunStateOperationOptions() CreateDeviceComplianceScriptDeviceRunStateOperationOptions {
	return CreateDeviceComplianceScriptDeviceRunStateOperationOptions{}
}

func (o CreateDeviceComplianceScriptDeviceRunStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceComplianceScriptDeviceRunStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceComplianceScriptDeviceRunStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceComplianceScriptDeviceRunState - Create new navigation property to deviceRunStates for deviceManagement
func (c DeviceComplianceScriptDeviceRunStateClient) CreateDeviceComplianceScriptDeviceRunState(ctx context.Context, id beta.DeviceManagementDeviceComplianceScriptId, input beta.DeviceComplianceScriptDeviceState, options CreateDeviceComplianceScriptDeviceRunStateOperationOptions) (result CreateDeviceComplianceScriptDeviceRunStateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceRunStates", id.ID()),
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

	var model beta.DeviceComplianceScriptDeviceState
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
