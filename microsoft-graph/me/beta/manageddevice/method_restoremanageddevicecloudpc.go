package manageddevice

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

type RestoreManagedDeviceCloudPCOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RestoreManagedDeviceCloudPCOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRestoreManagedDeviceCloudPCOperationOptions() RestoreManagedDeviceCloudPCOperationOptions {
	return RestoreManagedDeviceCloudPCOperationOptions{}
}

func (o RestoreManagedDeviceCloudPCOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RestoreManagedDeviceCloudPCOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RestoreManagedDeviceCloudPCOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RestoreManagedDeviceCloudPC - Invoke action restoreCloudPc. Restore a Cloud PC device to a previous state with an
// Intune managed device ID.
func (c ManagedDeviceClient) RestoreManagedDeviceCloudPC(ctx context.Context, id beta.MeManagedDeviceId, input RestoreManagedDeviceCloudPCRequest, options RestoreManagedDeviceCloudPCOperationOptions) (result RestoreManagedDeviceCloudPCOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/restoreCloudPc", id.ID()),
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
