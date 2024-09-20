package devicemanagementtroubleshootingevent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDeviceManagementTroubleshootingEventOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DeviceManagementTroubleshootingEvent
}

type CreateDeviceManagementTroubleshootingEventOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateDeviceManagementTroubleshootingEventOperationOptions() CreateDeviceManagementTroubleshootingEventOperationOptions {
	return CreateDeviceManagementTroubleshootingEventOperationOptions{}
}

func (o CreateDeviceManagementTroubleshootingEventOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceManagementTroubleshootingEventOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceManagementTroubleshootingEventOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceManagementTroubleshootingEvent - Create new navigation property to deviceManagementTroubleshootingEvents
// for users
func (c DeviceManagementTroubleshootingEventClient) CreateDeviceManagementTroubleshootingEvent(ctx context.Context, id beta.UserId, input beta.DeviceManagementTroubleshootingEvent, options CreateDeviceManagementTroubleshootingEventOperationOptions) (result CreateDeviceManagementTroubleshootingEventOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceManagementTroubleshootingEvents", id.ID()),
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalDeviceManagementTroubleshootingEventImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
