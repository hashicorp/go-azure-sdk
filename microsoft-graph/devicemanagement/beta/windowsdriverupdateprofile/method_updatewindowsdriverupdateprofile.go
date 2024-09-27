package windowsdriverupdateprofile

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateWindowsDriverUpdateProfileOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateWindowsDriverUpdateProfileOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateWindowsDriverUpdateProfileOperationOptions() UpdateWindowsDriverUpdateProfileOperationOptions {
	return UpdateWindowsDriverUpdateProfileOperationOptions{}
}

func (o UpdateWindowsDriverUpdateProfileOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateWindowsDriverUpdateProfileOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateWindowsDriverUpdateProfileOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateWindowsDriverUpdateProfile - Update the navigation property windowsDriverUpdateProfiles in deviceManagement
func (c WindowsDriverUpdateProfileClient) UpdateWindowsDriverUpdateProfile(ctx context.Context, id beta.DeviceManagementWindowsDriverUpdateProfileId, input beta.WindowsDriverUpdateProfile, options UpdateWindowsDriverUpdateProfileOperationOptions) (result UpdateWindowsDriverUpdateProfileOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
