package comanageddevicedevicehealthscriptstatedeviceiddeviceidididpolicyidpolicyid

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

type UpdateComanagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceHealthScriptPolicyState
}

// UpdateComanagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyId ...
func (c ComanagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdClient) UpdateComanagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyId(ctx context.Context, id DeviceManagementComanagedDeviceId, input beta.DeviceHealthScriptPolicyState) (result UpdateComanagedDeviceDeviceHealthScriptStateDeviceIdDeviceIdIdIdPolicyIdPolicyIdOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPatch,
		Path:       fmt.Sprintf("%s/deviceHealthScriptStates/deviceId='{deviceId}',id='{id}',policyId='{policyId}'", id.ID()),
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

	var model beta.DeviceHealthScriptPolicyState
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
