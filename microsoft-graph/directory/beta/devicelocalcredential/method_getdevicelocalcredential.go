package devicelocalcredential

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDeviceLocalCredentialOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceLocalCredentialInfo
}

type GetDeviceLocalCredentialOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDeviceLocalCredentialOperationOptions() GetDeviceLocalCredentialOperationOptions {
	return GetDeviceLocalCredentialOperationOptions{}
}

func (o GetDeviceLocalCredentialOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDeviceLocalCredentialOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDeviceLocalCredentialOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDeviceLocalCredential - Get deviceLocalCredentialInfo. Retrieve the properties of a deviceLocalCredentialInfo for
// a specified device object.
func (c DeviceLocalCredentialClient) GetDeviceLocalCredential(ctx context.Context, id beta.DirectoryDeviceLocalCredentialId, options GetDeviceLocalCredentialOperationOptions) (result GetDeviceLocalCredentialOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model beta.DeviceLocalCredentialInfo
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
