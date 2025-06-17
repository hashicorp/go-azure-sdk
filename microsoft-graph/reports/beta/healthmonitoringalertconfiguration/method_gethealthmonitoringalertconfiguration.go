package healthmonitoringalertconfiguration

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetHealthMonitoringAlertConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.HealthMonitoringAlertConfiguration
}

type GetHealthMonitoringAlertConfigurationOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetHealthMonitoringAlertConfigurationOperationOptions() GetHealthMonitoringAlertConfigurationOperationOptions {
	return GetHealthMonitoringAlertConfigurationOperationOptions{}
}

func (o GetHealthMonitoringAlertConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetHealthMonitoringAlertConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetHealthMonitoringAlertConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetHealthMonitoringAlertConfiguration - Get alertConfiguration. Read the properties and relationships of a Microsoft
// Entra health monitoring alertConfiguration object. The returned alertConfiguration object contains the settings for
// the distribution groups where alert notifications are to be sent.
func (c HealthMonitoringAlertConfigurationClient) GetHealthMonitoringAlertConfiguration(ctx context.Context, id beta.ReportHealthMonitoringAlertConfigurationId, options GetHealthMonitoringAlertConfigurationOperationOptions) (result GetHealthMonitoringAlertConfigurationOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.HealthMonitoringAlertConfiguration
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
