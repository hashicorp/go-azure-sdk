package invitedusermailboxsetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateInvitedUserMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateInvitedUserMailboxSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateInvitedUserMailboxSettingOperationOptions() UpdateInvitedUserMailboxSettingOperationOptions {
	return UpdateInvitedUserMailboxSettingOperationOptions{}
}

func (o UpdateInvitedUserMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateInvitedUserMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateInvitedUserMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateInvitedUserMailboxSetting - Update property mailboxSettings value.
func (c InvitedUserMailboxSettingClient) UpdateInvitedUserMailboxSetting(ctx context.Context, input beta.MailboxSettings, options UpdateInvitedUserMailboxSettingOperationOptions) (result UpdateInvitedUserMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/invitations/invitedUser/mailboxSettings",
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
