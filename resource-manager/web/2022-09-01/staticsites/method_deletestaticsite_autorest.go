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

type DeleteStaticSiteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteStaticSite ...
func (c StaticSitesClient) DeleteStaticSite(ctx context.Context, id StaticSiteId) (result DeleteStaticSiteOperationResponse, err error) {
	req, err := c.preparerForDeleteStaticSite(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteStaticSite", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteStaticSite(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "DeleteStaticSite", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteStaticSiteThenPoll performs DeleteStaticSite then polls until it's completed
func (c StaticSitesClient) DeleteStaticSiteThenPoll(ctx context.Context, id StaticSiteId) error {
	result, err := c.DeleteStaticSite(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteStaticSite: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteStaticSite: %+v", err)
	}

	return nil
}

// preparerForDeleteStaticSite prepares the DeleteStaticSite request.
func (c StaticSitesClient) preparerForDeleteStaticSite(ctx context.Context, id StaticSiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeleteStaticSite sends the DeleteStaticSite request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForDeleteStaticSite(ctx context.Context, req *http.Request) (future DeleteStaticSiteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
