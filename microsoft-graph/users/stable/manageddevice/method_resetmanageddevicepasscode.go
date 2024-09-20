package manageddevice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetManagedDevicePasscodeOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResetManagedDevicePasscodeOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultResetManagedDevicePasscodeOperationOptions() ResetManagedDevicePasscodeOperationOptions {
	return ResetManagedDevicePasscodeOperationOptions{}
}

func (o ResetManagedDevicePasscodeOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetManagedDevicePasscodeOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResetManagedDevicePasscodeOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResetManagedDevicePasscode - Invoke action resetPasscode. Reset passcode
func (c ManagedDeviceClient) ResetManagedDevicePasscode(ctx context.Context, id stable.UserIdManagedDeviceId, options ResetManagedDevicePasscodeOperationOptions) (result ResetManagedDevicePasscodeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resetPasscode", id.ID()),
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

	return
}
