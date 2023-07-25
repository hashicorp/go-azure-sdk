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

type LinkBackendToBuildOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// LinkBackendToBuild ...
func (c StaticSitesClient) LinkBackendToBuild(ctx context.Context, id BuildLinkedBackendId, input StaticSiteLinkedBackendARMResource) (result LinkBackendToBuildOperationResponse, err error) {
	req, err := c.preparerForLinkBackendToBuild(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "LinkBackendToBuild", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForLinkBackendToBuild(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "LinkBackendToBuild", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// LinkBackendToBuildThenPoll performs LinkBackendToBuild then polls until it's completed
func (c StaticSitesClient) LinkBackendToBuildThenPoll(ctx context.Context, id BuildLinkedBackendId, input StaticSiteLinkedBackendARMResource) error {
	result, err := c.LinkBackendToBuild(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing LinkBackendToBuild: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after LinkBackendToBuild: %+v", err)
	}

	return nil
}

// preparerForLinkBackendToBuild prepares the LinkBackendToBuild request.
func (c StaticSitesClient) preparerForLinkBackendToBuild(ctx context.Context, id BuildLinkedBackendId, input StaticSiteLinkedBackendARMResource) (*http.Request, error) {
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

// senderForLinkBackendToBuild sends the LinkBackendToBuild request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForLinkBackendToBuild(ctx context.Context, req *http.Request) (future LinkBackendToBuildOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
