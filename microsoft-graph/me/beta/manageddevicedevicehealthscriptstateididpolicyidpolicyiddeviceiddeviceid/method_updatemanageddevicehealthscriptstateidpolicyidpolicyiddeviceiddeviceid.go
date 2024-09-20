package manageddevicedevicehealthscriptstateididpolicyidpolicyiddeviceiddeviceid

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

type UpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions() UpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions {
	return UpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions{}
}

func (o UpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceId - Update the navigation property
// deviceHealthScriptStates in me
func (c ManagedDeviceDeviceHealthScriptStateIdIdPolicyIdPolicyIdDeviceIdDeviceIdClient) UpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceId(ctx context.Context, id beta.MeManagedDeviceId, input beta.DeviceHealthScriptPolicyState, options UpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationOptions) (result UpdateManagedDeviceHealthScriptStateIdPolicyIdPolicyIdDeviceIdDeviceIdOperationResponse, err error) {
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
