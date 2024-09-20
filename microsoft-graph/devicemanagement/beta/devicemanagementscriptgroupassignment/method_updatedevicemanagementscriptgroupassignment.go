package devicemanagementscriptgroupassignment

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceManagementScriptGroupAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceManagementScriptGroupAssignmentOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDeviceManagementScriptGroupAssignmentOperationOptions() UpdateDeviceManagementScriptGroupAssignmentOperationOptions {
	return UpdateDeviceManagementScriptGroupAssignmentOperationOptions{}
}

func (o UpdateDeviceManagementScriptGroupAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceManagementScriptGroupAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceManagementScriptGroupAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceManagementScriptGroupAssignment - Update the navigation property groupAssignments in deviceManagement
func (c DeviceManagementScriptGroupAssignmentClient) UpdateDeviceManagementScriptGroupAssignment(ctx context.Context, id beta.DeviceManagementDeviceManagementScriptIdGroupAssignmentId, input beta.DeviceManagementScriptGroupAssignment, options UpdateDeviceManagementScriptGroupAssignmentOperationOptions) (result UpdateDeviceManagementScriptGroupAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
