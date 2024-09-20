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

type UpdateDeviceLocalCredentialOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDeviceLocalCredentialOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDeviceLocalCredentialOperationOptions() UpdateDeviceLocalCredentialOperationOptions {
	return UpdateDeviceLocalCredentialOperationOptions{}
}

func (o UpdateDeviceLocalCredentialOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDeviceLocalCredentialOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDeviceLocalCredentialOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDeviceLocalCredential - Update the navigation property deviceLocalCredentials in directory
func (c DeviceLocalCredentialClient) UpdateDeviceLocalCredential(ctx context.Context, id beta.DirectoryDeviceLocalCredentialId, input beta.DeviceLocalCredentialInfo, options UpdateDeviceLocalCredentialOperationOptions) (result UpdateDeviceLocalCredentialOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
