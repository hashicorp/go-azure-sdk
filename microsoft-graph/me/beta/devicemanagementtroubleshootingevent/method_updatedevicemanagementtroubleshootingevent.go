package devicemanagementtroubleshootingevent

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateDeviceManagementTroubleshootingEventOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceManagementTroubleshootingEventOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDeviceManagementTroubleshootingEventOperationOptions() UpdateDeviceManagementTroubleshootingEventOperationOptions {
	return UpdateDeviceManagementTroubleshootingEventOperationOptions{}
}

func (o UpdateDeviceManagementTroubleshootingEventOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceManagementTroubleshootingEventOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceManagementTroubleshootingEventOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceManagementTroubleshootingEvent - Update the navigation property deviceManagementTroubleshootingEvents in
// me
func (c DeviceManagementTroubleshootingEventClient) UpdateDeviceManagementTroubleshootingEvent(ctx context.Context, id beta.MeDeviceManagementTroubleshootingEventId, input beta.DeviceManagementTroubleshootingEvent, options UpdateDeviceManagementTroubleshootingEventOperationOptions) (result UpdateDeviceManagementTroubleshootingEventOperationResponse, err error) {
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
