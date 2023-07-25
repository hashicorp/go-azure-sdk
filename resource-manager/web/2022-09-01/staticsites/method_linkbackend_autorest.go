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

type LinkBackendOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// LinkBackend ...
func (c StaticSitesClient) LinkBackend(ctx context.Context, id LinkedBackendId, input StaticSiteLinkedBackendARMResource) (result LinkBackendOperationResponse, err error) {
	req, err := c.preparerForLinkBackend(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "LinkBackend", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForLinkBackend(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "LinkBackend", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// LinkBackendThenPoll performs LinkBackend then polls until it's completed
func (c StaticSitesClient) LinkBackendThenPoll(ctx context.Context, id LinkedBackendId, input StaticSiteLinkedBackendARMResource) error {
	result, err := c.LinkBackend(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing LinkBackend: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after LinkBackend: %+v", err)
	}

	return nil
}

// preparerForLinkBackend prepares the LinkBackend request.
func (c StaticSitesClient) preparerForLinkBackend(ctx context.Context, id LinkedBackendId, input StaticSiteLinkedBackendARMResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForLinkBackend sends the LinkBackend request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForLinkBackend(ctx context.Context, req *http.Request) (future LinkBackendOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
