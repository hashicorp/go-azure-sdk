package restores

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

type TriggerOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

type TriggerOperationOptions struct {
	XMsAuthorizationAuxiliary *string
}

func DefaultTriggerOperationOptions() TriggerOperationOptions {
	return TriggerOperationOptions{}
}

func (o TriggerOperationOptions) toHeaders() map[string]interface{} {
	out := make(map[string]interface{})

	if o.XMsAuthorizationAuxiliary != nil {
		out["x-ms-authorization-auxiliary"] = *o.XMsAuthorizationAuxiliary
	}

	return out
}

func (o TriggerOperationOptions) toQueryString() map[string]interface{} {
	out := make(map[string]interface{})

	return out
}

// Trigger ...
func (c RestoresClient) Trigger(ctx context.Context, id RecoveryPointId, input RestoreRequestResource, options TriggerOperationOptions) (result TriggerOperationResponse, err error) {
	req, err := c.preparerForTrigger(ctx, id, input, options)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restores.RestoresClient", "Trigger", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForTrigger(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "restores.RestoresClient", "Trigger", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// TriggerThenPoll performs Trigger then polls until it's completed
func (c RestoresClient) TriggerThenPoll(ctx context.Context, id RecoveryPointId, input RestoreRequestResource, options TriggerOperationOptions) error {
	result, err := c.Trigger(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing Trigger: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after Trigger: %+v", err)
	}

	return nil
}

// preparerForTrigger prepares the Trigger request.
func (c RestoresClient) preparerForTrigger(ctx context.Context, id RecoveryPointId, input RestoreRequestResource, options TriggerOperationOptions) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	for k, v := range options.toQueryString() {
		queryParameters[k] = autorest.Encode("query", v)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithHeaders(options.toHeaders()),
		autorest.WithPath(fmt.Sprintf("%s/restore", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForTrigger sends the Trigger request. The method will close the
// http.Response Body if it receives an error.
func (c RestoresClient) senderForTrigger(ctx context.Context, req *http.Request) (future TriggerOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
