package authenticationmethoddevicehardwareoathdeviceassigntomailboxsetting

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

type UpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingOperationOptions() UpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingOperationOptions {
	return UpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingOperationOptions{}
}

func (o UpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSetting - Update property mailboxSettings value.
func (c AuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingClient) UpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSetting(ctx context.Context, id beta.DirectoryAuthenticationMethodDeviceHardwareOathDeviceId, input beta.MailboxSettings, options UpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingOperationOptions) (result UpdateAuthenticationMethodDeviceHardwareOathDeviceAssignToMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/assignTo/mailboxSettings", id.ID()),
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
