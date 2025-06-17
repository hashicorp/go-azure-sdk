package authenticationhardwareoathmethoddeviceassigntomailboxsetting

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

type UpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSettingOperationOptions() UpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSettingOperationOptions {
	return UpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSettingOperationOptions{}
}

func (o UpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSetting - Update property mailboxSettings value.
func (c AuthenticationHardwareOathMethodDeviceAssignToMailboxSettingClient) UpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSetting(ctx context.Context, id beta.MeAuthenticationHardwareOathMethodId, input beta.MailboxSettings, options UpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSettingOperationOptions) (result UpdateAuthenticationHardwareOathMethodDeviceAssignToMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/device/assignTo/mailboxSettings", id.ID()),
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
