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

type ListPublishingCredentialsOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ListPublishingCredentials ...
func (c WebAppsClient) ListPublishingCredentials(ctx context.Context, id SiteId) (result ListPublishingCredentialsOperationResponse, err error) {
	req, err := c.preparerForListPublishingCredentials(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublishingCredentials", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForListPublishingCredentials(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "ListPublishingCredentials", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ListPublishingCredentialsThenPoll performs ListPublishingCredentials then polls until it's completed
func (c WebAppsClient) ListPublishingCredentialsThenPoll(ctx context.Context, id SiteId) error {
	result, err := c.ListPublishingCredentials(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ListPublishingCredentials: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ListPublishingCredentials: %+v", err)
	}

	return nil
}

// preparerForListPublishingCredentials prepares the ListPublishingCredentials request.
func (c WebAppsClient) preparerForListPublishingCredentials(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/config/publishingcredentials/list", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForListPublishingCredentials sends the ListPublishingCredentials request. The method will close the
// http.Response Body if it receives an error.
func (c WebAppsClient) senderForListPublishingCredentials(ctx context.Context, req *http.Request) (future ListPublishingCredentialsOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
