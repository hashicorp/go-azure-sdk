package deviceconfiguration

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

type CreateDeviceConfigurationWindowsPrivacyAccessControlOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateDeviceConfigurationWindowsPrivacyAccessControlOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateDeviceConfigurationWindowsPrivacyAccessControlOperationOptions() CreateDeviceConfigurationWindowsPrivacyAccessControlOperationOptions {
	return CreateDeviceConfigurationWindowsPrivacyAccessControlOperationOptions{}
}

func (o CreateDeviceConfigurationWindowsPrivacyAccessControlOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateDeviceConfigurationWindowsPrivacyAccessControlOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateDeviceConfigurationWindowsPrivacyAccessControlOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateDeviceConfigurationWindowsPrivacyAccessControl - Invoke action windowsPrivacyAccessControls
func (c DeviceConfigurationClient) CreateDeviceConfigurationWindowsPrivacyAccessControl(ctx context.Context, id beta.DeviceManagementDeviceConfigurationId, input CreateDeviceConfigurationWindowsPrivacyAccessControlRequest, options CreateDeviceConfigurationWindowsPrivacyAccessControlOperationOptions) (result CreateDeviceConfigurationWindowsPrivacyAccessControlOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/windowsPrivacyAccessControls", id.ID()),
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
