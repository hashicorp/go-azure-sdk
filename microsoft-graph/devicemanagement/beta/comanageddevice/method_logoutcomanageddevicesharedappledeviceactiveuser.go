package comanageddevice

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LogoutComanagedDeviceSharedAppleDeviceActiveUserOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type LogoutComanagedDeviceSharedAppleDeviceActiveUserOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultLogoutComanagedDeviceSharedAppleDeviceActiveUserOperationOptions() LogoutComanagedDeviceSharedAppleDeviceActiveUserOperationOptions {
	return LogoutComanagedDeviceSharedAppleDeviceActiveUserOperationOptions{}
}

func (o LogoutComanagedDeviceSharedAppleDeviceActiveUserOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o LogoutComanagedDeviceSharedAppleDeviceActiveUserOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o LogoutComanagedDeviceSharedAppleDeviceActiveUserOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// LogoutComanagedDeviceSharedAppleDeviceActiveUser - Invoke action logoutSharedAppleDeviceActiveUser. Logout shared
// Apple device active user
func (c ComanagedDeviceClient) LogoutComanagedDeviceSharedAppleDeviceActiveUser(ctx context.Context, id beta.DeviceManagementComanagedDeviceId, options LogoutComanagedDeviceSharedAppleDeviceActiveUserOperationOptions) (result LogoutComanagedDeviceSharedAppleDeviceActiveUserOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/logoutSharedAppleDeviceActiveUser", id.ID()),
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
