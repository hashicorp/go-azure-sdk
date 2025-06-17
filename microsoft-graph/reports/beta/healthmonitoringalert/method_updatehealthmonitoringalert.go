package healthmonitoringalert

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateHealthMonitoringAlertOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateHealthMonitoringAlertOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateHealthMonitoringAlertOperationOptions() UpdateHealthMonitoringAlertOperationOptions {
	return UpdateHealthMonitoringAlertOperationOptions{}
}

func (o UpdateHealthMonitoringAlertOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateHealthMonitoringAlertOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateHealthMonitoringAlertOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateHealthMonitoringAlert - Update alert. Update the properties of a Microsoft Entra health monitoring alert
// object. For example, change an alert's state from active to resolved.
func (c HealthMonitoringAlertClient) UpdateHealthMonitoringAlert(ctx context.Context, id beta.ReportHealthMonitoringAlertId, input beta.HealthMonitoringAlert, options UpdateHealthMonitoringAlertOperationOptions) (result UpdateHealthMonitoringAlertOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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
