package comanageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid

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

type UpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions() UpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions {
	return UpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions{}
}

func (o UpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceId - Update the navigation property
// deviceHealthScriptStates in deviceManagement
func (c ComanagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient) UpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceId(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, input beta.DeviceHealthScriptPolicyState, options UpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions) (result UpdateComanagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceHealthScriptStates/id='{id}',policyId='{policyId}',deviceId='{deviceId}'", id.ID()),
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
