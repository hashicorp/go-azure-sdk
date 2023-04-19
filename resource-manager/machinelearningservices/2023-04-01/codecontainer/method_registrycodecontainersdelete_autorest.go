package codecontainer

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

type RegistryCodeContainersDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryCodeContainersDelete ...
func (c CodeContainerClient) RegistryCodeContainersDelete(ctx context.Context, id RegistryCodeId) (result RegistryCodeContainersDeleteOperationResponse, err error) {
	req, err := c.preparerForRegistryCodeContainersDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codecontainer.CodeContainerClient", "RegistryCodeContainersDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryCodeContainersDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codecontainer.CodeContainerClient", "RegistryCodeContainersDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryCodeContainersDeleteThenPoll performs RegistryCodeContainersDelete then polls until it's completed
func (c CodeContainerClient) RegistryCodeContainersDeleteThenPoll(ctx context.Context, id RegistryCodeId) error {
	result, err := c.RegistryCodeContainersDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RegistryCodeContainersDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryCodeContainersDelete: %+v", err)
	}

	return nil
}

// preparerForRegistryCodeContainersDelete prepares the RegistryCodeContainersDelete request.
func (c CodeContainerClient) preparerForRegistryCodeContainersDelete(ctx context.Context, id RegistryCodeId) (*http.Request, error) {
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

// senderForRegistryCodeContainersDelete sends the RegistryCodeContainersDelete request. The method will close the
// http.Response Body if it receives an error.
func (c CodeContainerClient) senderForRegistryCodeContainersDelete(ctx context.Context, req *http.Request) (future RegistryCodeContainersDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
