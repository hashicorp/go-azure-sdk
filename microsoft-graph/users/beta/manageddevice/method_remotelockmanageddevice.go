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

type RemoteLockManagedDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoteLockManagedDeviceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultRemoteLockManagedDeviceOperationOptions() RemoteLockManagedDeviceOperationOptions {
	return RemoteLockManagedDeviceOperationOptions{}
}

func (o RemoteLockManagedDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o RemoteLockManagedDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoteLockManagedDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoteLockManagedDevice - Invoke action remoteLock. Remote lock
func (c ManagedDeviceClient) RemoteLockManagedDevice(ctx context.Context, id beta.UserIdManagedDeviceId, options RemoteLockManagedDeviceOperationOptions) (result RemoteLockManagedDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/remoteLock", id.ID()),
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
