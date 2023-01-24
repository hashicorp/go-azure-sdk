package post

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

type ServicesStartOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ServicesStart ...
func (c POSTClient) ServicesStart(ctx context.Context, id ServiceId) (result ServicesStartOperationResponse, err error) {
	req, err := c.preparerForServicesStart(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "ServicesStart", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForServicesStart(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "ServicesStart", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ServicesStartThenPoll performs ServicesStart then polls until it's completed
func (c POSTClient) ServicesStartThenPoll(ctx context.Context, id ServiceId) error {
	result, err := c.ServicesStart(ctx, id)
	if err != nil {
		return fmt.Errorf("performing ServicesStart: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ServicesStart: %+v", err)
	}

	return nil
}

// preparerForServicesStart prepares the ServicesStart request.
func (c POSTClient) preparerForServicesStart(ctx context.Context, id ServiceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/start", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForServicesStart sends the ServicesStart request. The method will close the
// http.Response Body if it receives an error.
func (c POSTClient) senderForServicesStart(ctx context.Context, req *http.Request) (future ServicesStartOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
