package deviceregistereduser

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

type AddDeviceRegisteredUserRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type AddDeviceRegisteredUserRefOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultAddDeviceRegisteredUserRefOperationOptions() AddDeviceRegisteredUserRefOperationOptions {
	return AddDeviceRegisteredUserRefOperationOptions{}
}

func (o AddDeviceRegisteredUserRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AddDeviceRegisteredUserRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o AddDeviceRegisteredUserRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// AddDeviceRegisteredUserRef - Create new navigation property ref to registeredUsers for me
func (c DeviceRegisteredUserClient) AddDeviceRegisteredUserRef(ctx context.Context, id beta.MeDeviceId, input beta.ReferenceCreate, options AddDeviceRegisteredUserRefOperationOptions) (result AddDeviceRegisteredUserRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/registeredUsers/$ref", id.ID()),
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
