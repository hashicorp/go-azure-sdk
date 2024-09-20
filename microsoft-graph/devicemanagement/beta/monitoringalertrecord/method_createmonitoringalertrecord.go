package monitoringalertrecord

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMonitoringAlertRecordOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceManagementAlertRecord
}

type CreateMonitoringAlertRecordOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateMonitoringAlertRecordOperationOptions() CreateMonitoringAlertRecordOperationOptions {
	return CreateMonitoringAlertRecordOperationOptions{}
}

func (o CreateMonitoringAlertRecordOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateMonitoringAlertRecordOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateMonitoringAlertRecordOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateMonitoringAlertRecord - Create new navigation property to alertRecords for deviceManagement
func (c MonitoringAlertRecordClient) CreateMonitoringAlertRecord(ctx context.Context, input beta.DeviceManagementAlertRecord, options CreateMonitoringAlertRecordOperationOptions) (result CreateMonitoringAlertRecordOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/monitoring/alertRecords",
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

	var model beta.DeviceManagementAlertRecord
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
