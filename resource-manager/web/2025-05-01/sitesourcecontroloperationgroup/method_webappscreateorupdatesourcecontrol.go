package sitesourcecontroloperationgroup

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

type WebAppsCreateOrUpdateSourceControlOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SiteSourceControl
}

// WebAppsCreateOrUpdateSourceControl ...
func (c SiteSourceControlOperationGroupClient) WebAppsCreateOrUpdateSourceControl(ctx context.Context, id commonids.AppServiceId, input SiteSourceControl) (result WebAppsCreateOrUpdateSourceControlOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/sourceControls/web", id.ID()),
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

// WebAppsCreateOrUpdateSourceControlThenPoll performs WebAppsCreateOrUpdateSourceControl then polls until it's completed
func (c SiteSourceControlOperationGroupClient) WebAppsCreateOrUpdateSourceControlThenPoll(ctx context.Context, id commonids.AppServiceId, input SiteSourceControl) error {
	result, err := c.WebAppsCreateOrUpdateSourceControl(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing WebAppsCreateOrUpdateSourceControl: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after WebAppsCreateOrUpdateSourceControl: %+v", err)
	}

	return nil
}
