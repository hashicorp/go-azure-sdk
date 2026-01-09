package healthmonitoringalertconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateHealthMonitoringAlertConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateHealthMonitoringAlertConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateHealthMonitoringAlertConfigurationOperationOptions() UpdateHealthMonitoringAlertConfigurationOperationOptions {
	return UpdateHealthMonitoringAlertConfigurationOperationOptions{}
}

func (o UpdateHealthMonitoringAlertConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateHealthMonitoringAlertConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateHealthMonitoringAlertConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateHealthMonitoringAlertConfiguration - Update alertConfiguration. Update the properties of a Microsoft Entra
// health monitoring alertConfiguration object. You can use alertConfiguration settings to specify the distribution
// groups where alert notifications are to be sent. This API doesn't currently support group validation.
func (c HealthMonitoringAlertConfigurationClient) UpdateHealthMonitoringAlertConfiguration(ctx context.Context, id beta.ReportHealthMonitoringAlertConfigurationId, input beta.HealthMonitoringAlertConfiguration, options UpdateHealthMonitoringAlertConfigurationOperationOptions) (result UpdateHealthMonitoringAlertConfigurationOperationResponse, err error) {
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
