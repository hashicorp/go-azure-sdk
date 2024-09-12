package mailboxsetting

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MailboxSettings
}

type GetMailboxSettingOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetMailboxSettingOperationOptions() GetMailboxSettingOperationOptions {
	return GetMailboxSettingOperationOptions{}
}

func (o GetMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetMailboxSetting - Get user mailbox settings. Get the user's mailboxSettings. You can view all mailbox settings, or
// get specific settings. Users can set the following settings for their mailboxes through an Outlook client: Users can
// set their preferred date and time formats using Outlook on the web. Users can choose one of the supported short date
// or short time formats. This GET operation returns the format the user has chosen. Users can set the time zone they
// prefer on any Outlook client, by choosing from the supported time zones that their administrator has set up for their
// mailbox server. The administrator can set up time zones in the Windows time zone format or Internet Assigned Numbers
// Authority (IANA) time zone (also known as Olson time zone) format. The Windows format is the default. This GET
// operation returns the user's preferred time zone in the format that the administrator has set up. If you want that
// time zone to be in a specific format (Windows or IANA), you can first update the preferred time zone in that format
// as a mailbox setting. Subsequently you will be able to get the time zone in that format. Alternatively, you can
// manage the format conversion separately in your app.
func (c MailboxSettingClient) GetMailboxSetting(ctx context.Context, options GetMailboxSettingOperationOptions) (result GetMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          "/me/mailboxSettings",
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

	var model beta.MailboxSettings
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
