package appserviceenvironmentresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppServiceEnvironmentsUpgradeOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// AppServiceEnvironmentsUpgrade ...
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsUpgrade(ctx context.Context, id commonids.AppServiceEnvironmentId) (result AppServiceEnvironmentsUpgradeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/upgrade", id.ID()),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// AppServiceEnvironmentsUpgradeThenPoll performs AppServiceEnvironmentsUpgrade then polls until it's completed
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsUpgradeThenPoll(ctx context.Context, id commonids.AppServiceEnvironmentId) error {
	result, err := c.AppServiceEnvironmentsUpgrade(ctx, id)
	if err != nil {
		return fmt.Errorf("performing AppServiceEnvironmentsUpgrade: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after AppServiceEnvironmentsUpgrade: %+v", err)
	}

	return nil
}
