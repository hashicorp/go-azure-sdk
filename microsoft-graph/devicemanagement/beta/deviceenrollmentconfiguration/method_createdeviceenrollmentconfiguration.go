package deviceenrollmentconfiguration

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateDeviceEnrollmentConfigurationOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        beta.DeviceEnrollmentConfiguration
}

type CreateDeviceEnrollmentConfigurationOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDeviceEnrollmentConfigurationOperationOptions() CreateDeviceEnrollmentConfigurationOperationOptions {
	return CreateDeviceEnrollmentConfigurationOperationOptions{}
}

func (o CreateDeviceEnrollmentConfigurationOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceEnrollmentConfigurationOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceEnrollmentConfigurationOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceEnrollmentConfiguration - Create new navigation property to deviceEnrollmentConfigurations for
// deviceManagement
func (c DeviceEnrollmentConfigurationClient) CreateDeviceEnrollmentConfiguration(ctx context.Context, input beta.DeviceEnrollmentConfiguration, options CreateDeviceEnrollmentConfigurationOperationOptions) (result CreateDeviceEnrollmentConfigurationOperationResponse, err error) {
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
		Path:          "/deviceManagement/deviceEnrollmentConfigurations",
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

	var respObj json.RawMessage
	if err = resp.Unmarshal(&respObj); err != nil {
		return
	}
	model, err := beta.UnmarshalDeviceEnrollmentConfigurationImplementation(respObj)
	if err != nil {
		return
	}
	result.Model = model

	return
}
