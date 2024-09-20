package importeddeviceidentity

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateImportedDeviceIdentityOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateImportedDeviceIdentityOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateImportedDeviceIdentityOperationOptions() UpdateImportedDeviceIdentityOperationOptions {
	return UpdateImportedDeviceIdentityOperationOptions{}
}

func (o UpdateImportedDeviceIdentityOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateImportedDeviceIdentityOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateImportedDeviceIdentityOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateImportedDeviceIdentity - Update the navigation property importedDeviceIdentities in deviceManagement
func (c ImportedDeviceIdentityClient) UpdateImportedDeviceIdentity(ctx context.Context, id beta.DeviceManagementImportedDeviceIdentityId, input beta.ImportedDeviceIdentity, options UpdateImportedDeviceIdentityOperationOptions) (result UpdateImportedDeviceIdentityOperationResponse, err error) {
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
