package authenticationwindowshelloforbusinessmethoddevice

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

type GetAuthenticationWindowsHelloForBusinessMethodDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Device
}

type GetAuthenticationWindowsHelloForBusinessMethodDeviceOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetAuthenticationWindowsHelloForBusinessMethodDeviceOperationOptions() GetAuthenticationWindowsHelloForBusinessMethodDeviceOperationOptions {
	return GetAuthenticationWindowsHelloForBusinessMethodDeviceOperationOptions{}
}

func (o GetAuthenticationWindowsHelloForBusinessMethodDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetAuthenticationWindowsHelloForBusinessMethodDeviceOperationOptions) ToOData() *odata.Query {
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

func (o GetAuthenticationWindowsHelloForBusinessMethodDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetAuthenticationWindowsHelloForBusinessMethodDevice - Get device from me. The registered device on which this
// Windows Hello for Business key resides. Supports $expand. When you get a user's Windows Hello for Business
// registration information, this property is returned only on a single GET and when you specify ?$expand. For example,
// GET /users/admin@contoso.com/authentication/windowsHelloForBusinessMethods/_jpuR-TGZtk6aQCLF3BQjA2?$expand=device.
func (c AuthenticationWindowsHelloForBusinessMethodDeviceClient) GetAuthenticationWindowsHelloForBusinessMethodDevice(ctx context.Context, id beta.MeAuthenticationWindowsHelloForBusinessMethodId, options GetAuthenticationWindowsHelloForBusinessMethodDeviceOperationOptions) (result GetAuthenticationWindowsHelloForBusinessMethodDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/device", id.ID()),
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

	var model beta.Device
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
