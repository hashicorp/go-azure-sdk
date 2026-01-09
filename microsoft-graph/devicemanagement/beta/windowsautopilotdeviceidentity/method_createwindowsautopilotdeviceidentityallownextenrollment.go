package windowsautopilotdeviceidentity

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateWindowsAutopilotDeviceIdentityAllowNextEnrollmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateWindowsAutopilotDeviceIdentityAllowNextEnrollmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateWindowsAutopilotDeviceIdentityAllowNextEnrollmentOperationOptions() CreateWindowsAutopilotDeviceIdentityAllowNextEnrollmentOperationOptions {
	return CreateWindowsAutopilotDeviceIdentityAllowNextEnrollmentOperationOptions{}
}

func (o CreateWindowsAutopilotDeviceIdentityAllowNextEnrollmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateWindowsAutopilotDeviceIdentityAllowNextEnrollmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateWindowsAutopilotDeviceIdentityAllowNextEnrollmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateWindowsAutopilotDeviceIdentityAllowNextEnrollment - Invoke action allowNextEnrollment. Unblocks next autopilot
// enrollment.
func (c WindowsAutopilotDeviceIdentityClient) CreateWindowsAutopilotDeviceIdentityAllowNextEnrollment(ctx context.Context, id beta.DeviceManagementWindowsAutopilotDeviceIdentityId, options CreateWindowsAutopilotDeviceIdentityAllowNextEnrollmentOperationOptions) (result CreateWindowsAutopilotDeviceIdentityAllowNextEnrollmentOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/allowNextEnrollment", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	return
}
