package monitoringalertrecord

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

type CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationOptions() CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationOptions {
	return CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationOptions{}
}

func (o CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSent - Invoke action setPortalNotificationAsSent.
// Set the status of the notification associated with the specified alertRecord on the Microsoft EndPoint Manager admin
// center as sent, by setting the isPortalNotificationSent property of the portal notification to true.
func (c MonitoringAlertRecordClient) CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSent(ctx context.Context, id beta.DeviceManagementMonitoringAlertRecordId, options CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationOptions) (result CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deviceManagement.setPortalNotificationAsSent", id.ID()),
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
