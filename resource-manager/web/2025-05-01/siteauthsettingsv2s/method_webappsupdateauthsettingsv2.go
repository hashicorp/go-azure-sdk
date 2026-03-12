package siteauthsettingsv2s

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAppsUpdateAuthSettingsV2OperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SiteAuthSettingsV2
}

// WebAppsUpdateAuthSettingsV2 ...
func (c SiteAuthSettingsV2sClient) WebAppsUpdateAuthSettingsV2(ctx context.Context, id commonids.AppServiceId, input SiteAuthSettingsV2) (result WebAppsUpdateAuthSettingsV2OperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/config/authsettingsV2", id.ID()),
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

	var model SiteAuthSettingsV2
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
