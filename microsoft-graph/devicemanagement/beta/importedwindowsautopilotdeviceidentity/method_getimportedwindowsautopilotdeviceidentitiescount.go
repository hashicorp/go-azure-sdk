package importedwindowsautopilotdeviceidentity

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetImportedWindowsAutopilotDeviceIdentitiesCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetImportedWindowsAutopilotDeviceIdentitiesCountOperationOptions struct {
	Filter    *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Search    *string
}

func DefaultGetImportedWindowsAutopilotDeviceIdentitiesCountOperationOptions() GetImportedWindowsAutopilotDeviceIdentitiesCountOperationOptions {
	return GetImportedWindowsAutopilotDeviceIdentitiesCountOperationOptions{}
}

func (o GetImportedWindowsAutopilotDeviceIdentitiesCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetImportedWindowsAutopilotDeviceIdentitiesCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetImportedWindowsAutopilotDeviceIdentitiesCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetImportedWindowsAutopilotDeviceIdentitiesCount - Get the number of the resource
func (c ImportedWindowsAutopilotDeviceIdentityClient) GetImportedWindowsAutopilotDeviceIdentitiesCount(ctx context.Context, options GetImportedWindowsAutopilotDeviceIdentitiesCountOperationOptions) (result GetImportedWindowsAutopilotDeviceIdentitiesCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/deviceManagement/importedWindowsAutopilotDeviceIdentities/$count",
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
