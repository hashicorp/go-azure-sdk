package monitoringalertrecord

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMonitoringAlertRecordDeviceManagementChangePortalNotificationAsSentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

// CreateMonitoringAlertRecordDeviceManagementChangePortalNotificationAsSent - Invoke action
// changeAlertRecordsPortalNotificationAsSent. Set the isPortalNotificationSent property of all portal notification
// resources associated with the specified alertRecord to true, marking them as sent. A maximum of 100 alertRecord IDs
// can be received at one time, and a maximum of 100 portal notification resources can be changed in the
// isPortalNotificationSent property status.
func (c MonitoringAlertRecordClient) CreateMonitoringAlertRecordDeviceManagementChangePortalNotificationAsSent(ctx context.Context, input CreateMonitoringAlertRecordDeviceManagementChangePortalNotificationAsSentRequest) (result CreateMonitoringAlertRecordDeviceManagementChangePortalNotificationAsSentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/monitoring/alertRecords/deviceManagement.changeAlertRecordsPortalNotificationAsSent",
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
