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

type DeleteMonitoringAlertRecordOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteMonitoringAlertRecordOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteMonitoringAlertRecordOperationOptions() DeleteMonitoringAlertRecordOperationOptions {
	return DeleteMonitoringAlertRecordOperationOptions{}
}

func (o DeleteMonitoringAlertRecordOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteMonitoringAlertRecordOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteMonitoringAlertRecordOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteMonitoringAlertRecord - Delete navigation property alertRecords for deviceManagement
func (c MonitoringAlertRecordClient) DeleteMonitoringAlertRecord(ctx context.Context, id beta.DeviceManagementMonitoringAlertRecordId, options DeleteMonitoringAlertRecordOperationOptions) (result DeleteMonitoringAlertRecordOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
