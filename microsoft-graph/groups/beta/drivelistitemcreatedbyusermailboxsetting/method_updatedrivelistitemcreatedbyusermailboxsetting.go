package drivelistitemcreatedbyusermailboxsetting

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

type UpdateDriveListItemCreatedByUserMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateDriveListItemCreatedByUserMailboxSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateDriveListItemCreatedByUserMailboxSettingOperationOptions() UpdateDriveListItemCreatedByUserMailboxSettingOperationOptions {
	return UpdateDriveListItemCreatedByUserMailboxSettingOperationOptions{}
}

func (o UpdateDriveListItemCreatedByUserMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateDriveListItemCreatedByUserMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateDriveListItemCreatedByUserMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateDriveListItemCreatedByUserMailboxSetting - Update property mailboxSettings value.
func (c DriveListItemCreatedByUserMailboxSettingClient) UpdateDriveListItemCreatedByUserMailboxSetting(ctx context.Context, id beta.GroupIdDriveIdListItemId, input beta.MailboxSettings, options UpdateDriveListItemCreatedByUserMailboxSettingOperationOptions) (result UpdateDriveListItemCreatedByUserMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/createdByUser/mailboxSettings", id.ID()),
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
