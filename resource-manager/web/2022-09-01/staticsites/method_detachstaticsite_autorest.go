package staticsites

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

type DetachStaticSiteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DetachStaticSite ...
func (c StaticSitesClient) DetachStaticSite(ctx context.Context, id StaticSiteId) (result DetachStaticSiteOperationResponse, err error) {
	req, err := c.preparerForDetachStaticSite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DetachStaticSite", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDetachStaticSite(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DetachStaticSite", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DetachStaticSiteThenPoll performs DetachStaticSite then polls until it's completed
func (c StaticSitesClient) DetachStaticSiteThenPoll(ctx context.Context, id StaticSiteId) error {
	result, err := c.DetachStaticSite(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DetachStaticSite: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DetachStaticSite: %+v", err)
	}

	return nil
}

// preparerForDetachStaticSite prepares the DetachStaticSite request.
func (c StaticSitesClient) preparerForDetachStaticSite(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/detach", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDetachStaticSite sends the DetachStaticSite request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForDetachStaticSite(ctx context.Context, req *http.Request) (future DetachStaticSiteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
