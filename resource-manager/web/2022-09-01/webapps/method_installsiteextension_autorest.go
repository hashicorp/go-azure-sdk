package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InstallSiteExtensionOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// InstallSiteExtension ...
func (c WebAppsClient) InstallSiteExtension(ctx context.Context, id SiteExtensionId) (result InstallSiteExtensionOperationResponse, err error) {
	req, err := c.preparerForInstallSiteExtension(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "InstallSiteExtension", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForInstallSiteExtension(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "InstallSiteExtension", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// InstallSiteExtensionThenPoll performs InstallSiteExtension then polls until it's completed
func (c WebAppsClient) InstallSiteExtensionThenPoll(ctx context.Context, id SiteExtensionId) error {
	result, err := c.InstallSiteExtension(ctx, id)
	if err != nil {
		return fmt.Errorf("performing InstallSiteExtension: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after InstallSiteExtension: %+v", err)
	}

	return nil
}

// preparerForInstallSiteExtension prepares the InstallSiteExtension request.
func (c WebAppsClient) preparerForInstallSiteExtension(ctx context.Context, id SiteExtensionId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForInstallSiteExtension sends the InstallSiteExtension request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForInstallSiteExtension(ctx context.Context, req *http.Request) (future InstallSiteExtensionOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
