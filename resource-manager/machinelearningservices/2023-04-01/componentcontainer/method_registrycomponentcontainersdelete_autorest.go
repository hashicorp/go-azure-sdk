package componentcontainer

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

type RegistryComponentContainersDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryComponentContainersDelete ...
func (c ComponentContainerClient) RegistryComponentContainersDelete(ctx context.Context, id RegistryComponentId) (result RegistryComponentContainersDeleteOperationResponse, err error) {
	req, err := c.preparerForRegistryComponentContainersDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentcontainer.ComponentContainerClient", "RegistryComponentContainersDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryComponentContainersDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentcontainer.ComponentContainerClient", "RegistryComponentContainersDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryComponentContainersDeleteThenPoll performs RegistryComponentContainersDelete then polls until it's completed
func (c ComponentContainerClient) RegistryComponentContainersDeleteThenPoll(ctx context.Context, id RegistryComponentId) error {
	result, err := c.RegistryComponentContainersDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RegistryComponentContainersDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryComponentContainersDelete: %+v", err)
	}

	return nil
}

// preparerForRegistryComponentContainersDelete prepares the RegistryComponentContainersDelete request.
func (c ComponentContainerClient) preparerForRegistryComponentContainersDelete(ctx context.Context, id RegistryComponentId) (*http.Request, error) {
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

// senderForRegistryComponentContainersDelete sends the RegistryComponentContainersDelete request. The method will close the
// http.Response Body if it receives an error.
func (c ComponentContainerClient) senderForRegistryComponentContainersDelete(ctx context.Context, req *http.Request) (future RegistryComponentContainersDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
