package namedvalue

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

type RefreshSecretOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RefreshSecret ...
func (c NamedValueClient) RefreshSecret(ctx context.Context, id NamedValueId) (result RefreshSecretOperationResponse, err error) {
	req, err := c.preparerForRefreshSecret(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "namedvalue.NamedValueClient", "RefreshSecret", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRefreshSecret(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "namedvalue.NamedValueClient", "RefreshSecret", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RefreshSecretThenPoll performs RefreshSecret then polls until it's completed
func (c NamedValueClient) RefreshSecretThenPoll(ctx context.Context, id NamedValueId) error {
	result, err := c.RefreshSecret(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RefreshSecret: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RefreshSecret: %+v", err)
	}

	return nil
}

// preparerForRefreshSecret prepares the RefreshSecret request.
func (c NamedValueClient) preparerForRefreshSecret(ctx context.Context, id NamedValueId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/refreshSecret", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForRefreshSecret sends the RefreshSecret request. The method will close the
// http.Response Body if it receives an error.
func (c NamedValueClient) senderForRefreshSecret(ctx context.Context, req *http.Request) (future RefreshSecretOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
