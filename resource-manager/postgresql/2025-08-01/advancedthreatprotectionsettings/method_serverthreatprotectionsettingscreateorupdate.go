package advancedthreatprotectionsettings

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerThreatProtectionSettingsCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ServerThreatProtectionSettingsCreateOrUpdate ...
func (c AdvancedThreatProtectionSettingsClient) ServerThreatProtectionSettingsCreateOrUpdate(ctx context.Context, id FlexibleServerId, input AdvancedThreatProtectionSettingsModel) (result ServerThreatProtectionSettingsCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/advancedThreatProtectionSettings/default", id.ID()),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// ServerThreatProtectionSettingsCreateOrUpdateThenPoll performs ServerThreatProtectionSettingsCreateOrUpdate then polls until it's completed
func (c AdvancedThreatProtectionSettingsClient) ServerThreatProtectionSettingsCreateOrUpdateThenPoll(ctx context.Context, id FlexibleServerId, input AdvancedThreatProtectionSettingsModel) error {
	result, err := c.ServerThreatProtectionSettingsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ServerThreatProtectionSettingsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ServerThreatProtectionSettingsCreateOrUpdate: %+v", err)
	}

	return nil
}
