package actions

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

type CloudEndpointsPostRestoreOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CloudEndpointsPostRestore ...
func (c ActionsClient) CloudEndpointsPostRestore(ctx context.Context, id CloudEndpointId, input PostRestoreRequest) (result CloudEndpointsPostRestoreOperationResponse, err error) {
	req, err := c.preparerForCloudEndpointsPostRestore(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "actions.ActionsClient", "CloudEndpointsPostRestore", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCloudEndpointsPostRestore(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "actions.ActionsClient", "CloudEndpointsPostRestore", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CloudEndpointsPostRestoreThenPoll performs CloudEndpointsPostRestore then polls until it's completed
func (c ActionsClient) CloudEndpointsPostRestoreThenPoll(ctx context.Context, id CloudEndpointId, input PostRestoreRequest) error {
	result, err := c.CloudEndpointsPostRestore(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CloudEndpointsPostRestore: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CloudEndpointsPostRestore: %+v", err)
	}

	return nil
}

// preparerForCloudEndpointsPostRestore prepares the CloudEndpointsPostRestore request.
func (c ActionsClient) preparerForCloudEndpointsPostRestore(ctx context.Context, id CloudEndpointId, input PostRestoreRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/postrestore", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCloudEndpointsPostRestore sends the CloudEndpointsPostRestore request. The method will close the
// http.Response Body if it receives an error.
func (c ActionsClient) senderForCloudEndpointsPostRestore(ctx context.Context, req *http.Request) (future CloudEndpointsPostRestoreOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
