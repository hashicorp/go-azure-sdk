package manageddevice

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateManagedDeviceDownloadAppDiagnosticOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type CreateManagedDeviceDownloadAppDiagnosticOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateManagedDeviceDownloadAppDiagnosticOperationOptions() CreateManagedDeviceDownloadAppDiagnosticOperationOptions {
	return CreateManagedDeviceDownloadAppDiagnosticOperationOptions{}
}

func (o CreateManagedDeviceDownloadAppDiagnosticOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateManagedDeviceDownloadAppDiagnosticOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateManagedDeviceDownloadAppDiagnosticOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateManagedDeviceDownloadAppDiagnostic - Invoke action downloadAppDiagnostics
func (c ManagedDeviceClient) CreateManagedDeviceDownloadAppDiagnostic(ctx context.Context, input CreateManagedDeviceDownloadAppDiagnosticRequest, options CreateManagedDeviceDownloadAppDiagnosticOperationOptions) (result CreateManagedDeviceDownloadAppDiagnosticOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/managedDevices/downloadAppDiagnostics",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
