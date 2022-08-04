package delete

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

type IotDpsResourceDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// IotDpsResourceDelete ...
func (c DELETEClient) IotDpsResourceDelete(ctx context.Context, id ProvisioningServiceId) (result IotDpsResourceDeleteOperationResponse, err error) {
	req, err := c.preparerForIotDpsResourceDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delete.DELETEClient", "IotDpsResourceDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForIotDpsResourceDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "delete.DELETEClient", "IotDpsResourceDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// IotDpsResourceDeleteThenPoll performs IotDpsResourceDelete then polls until it's completed
func (c DELETEClient) IotDpsResourceDeleteThenPoll(ctx context.Context, id ProvisioningServiceId) error {
	result, err := c.IotDpsResourceDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing IotDpsResourceDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after IotDpsResourceDelete: %+v", err)
	}

	return nil
}

// preparerForIotDpsResourceDelete prepares the IotDpsResourceDelete request.
func (c DELETEClient) preparerForIotDpsResourceDelete(ctx context.Context, id ProvisioningServiceId) (*http.Request, error) {
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

// senderForIotDpsResourceDelete sends the IotDpsResourceDelete request. The method will close the
// http.Response Body if it receives an error.
func (c DELETEClient) senderForIotDpsResourceDelete(ctx context.Context, req *http.Request) (future IotDpsResourceDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
