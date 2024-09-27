package comanageddevice

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

type CreateComanagedDeviceRotateBitLockerKeyOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateComanagedDeviceRotateBitLockerKeyOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateComanagedDeviceRotateBitLockerKeyOperationOptions() CreateComanagedDeviceRotateBitLockerKeyOperationOptions {
	return CreateComanagedDeviceRotateBitLockerKeyOperationOptions{}
}

func (o CreateComanagedDeviceRotateBitLockerKeyOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateComanagedDeviceRotateBitLockerKeyOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateComanagedDeviceRotateBitLockerKeyOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateComanagedDeviceRotateBitLockerKey - Invoke action rotateBitLockerKeys. Rotate BitLockerKeys
func (c ComanagedDeviceClient) CreateComanagedDeviceRotateBitLockerKey(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options CreateComanagedDeviceRotateBitLockerKeyOperationOptions) (result CreateComanagedDeviceRotateBitLockerKeyOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/rotateBitLockerKeys", id.ID()),
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

	return
}
