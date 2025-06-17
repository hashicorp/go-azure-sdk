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

type CreateHealthMonitoringAlertOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HealthMonitoringAlert
}

type CreateHealthMonitoringAlertOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateHealthMonitoringAlertOperationOptions() CreateHealthMonitoringAlertOperationOptions {
	return CreateHealthMonitoringAlertOperationOptions{}
}

func (o CreateHealthMonitoringAlertOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateHealthMonitoringAlertOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateHealthMonitoringAlertOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateHealthMonitoringAlert - Create new navigation property to alerts for reports
func (c HealthMonitoringAlertClient) CreateHealthMonitoringAlert(ctx context.Context, input beta.HealthMonitoringAlert, options CreateHealthMonitoringAlertOperationOptions) (result CreateHealthMonitoringAlertOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/reports/healthMonitoring/alerts",
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

	var model beta.HealthMonitoringAlert
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
