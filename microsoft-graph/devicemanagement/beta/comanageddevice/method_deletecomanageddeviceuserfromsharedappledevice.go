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

type DeleteComanagedDeviceUserFromSharedAppleDeviceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteComanagedDeviceUserFromSharedAppleDeviceOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultDeleteComanagedDeviceUserFromSharedAppleDeviceOperationOptions() DeleteComanagedDeviceUserFromSharedAppleDeviceOperationOptions {
	return DeleteComanagedDeviceUserFromSharedAppleDeviceOperationOptions{}
}

func (o DeleteComanagedDeviceUserFromSharedAppleDeviceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeleteComanagedDeviceUserFromSharedAppleDeviceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteComanagedDeviceUserFromSharedAppleDeviceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteComanagedDeviceUserFromSharedAppleDevice - Invoke action deleteUserFromSharedAppleDevice. Delete user from
// shared Apple device
func (c ComanagedDeviceClient) DeleteComanagedDeviceUserFromSharedAppleDevice(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, input DeleteComanagedDeviceUserFromSharedAppleDeviceRequest, options DeleteComanagedDeviceUserFromSharedAppleDeviceOperationOptions) (result DeleteComanagedDeviceUserFromSharedAppleDeviceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/deleteUserFromSharedAppleDevice", id.ID()),
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
