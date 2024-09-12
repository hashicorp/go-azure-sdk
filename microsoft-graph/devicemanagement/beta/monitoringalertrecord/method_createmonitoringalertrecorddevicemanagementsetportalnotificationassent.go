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

// CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSent - Invoke action setPortalNotificationAsSent.
// Set the status of the notification associated with the specified alertRecord on the Microsoft EndPoint Manager admin
// center as sent, by setting the isPortalNotificationSent property of the portal notification to true.
func (c MonitoringAlertRecordClient) CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSent(ctx context.Context, id beta.DeviceManagementMonitoringAlertRecordId) (result CreateMonitoringAlertRecordDeviceManagementSetPortalNotificationAsSentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/deviceManagement.setPortalNotificationAsSent", id.ID()),
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
