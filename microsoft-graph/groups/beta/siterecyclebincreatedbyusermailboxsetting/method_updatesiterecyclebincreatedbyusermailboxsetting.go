package siterecyclebincreatedbyusermailboxsetting

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

type UpdateSiteRecycleBinCreatedByUserMailboxSettingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateSiteRecycleBinCreatedByUserMailboxSettingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateSiteRecycleBinCreatedByUserMailboxSettingOperationOptions() UpdateSiteRecycleBinCreatedByUserMailboxSettingOperationOptions {
	return UpdateSiteRecycleBinCreatedByUserMailboxSettingOperationOptions{}
}

func (o UpdateSiteRecycleBinCreatedByUserMailboxSettingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateSiteRecycleBinCreatedByUserMailboxSettingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateSiteRecycleBinCreatedByUserMailboxSettingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateSiteRecycleBinCreatedByUserMailboxSetting - Update property mailboxSettings value.
func (c SiteRecycleBinCreatedByUserMailboxSettingClient) UpdateSiteRecycleBinCreatedByUserMailboxSetting(ctx context.Context, id beta.GroupIdSiteId, input beta.MailboxSettings, options UpdateSiteRecycleBinCreatedByUserMailboxSettingOperationOptions) (result UpdateSiteRecycleBinCreatedByUserMailboxSettingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/recycleBin/createdByUser/mailboxSettings", id.ID()),
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
