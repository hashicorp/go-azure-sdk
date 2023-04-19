package environmentcontainer

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

type RegistryEnvironmentContainersDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryEnvironmentContainersDelete ...
func (c EnvironmentContainerClient) RegistryEnvironmentContainersDelete(ctx context.Context, id RegistryEnvironmentId) (result RegistryEnvironmentContainersDeleteOperationResponse, err error) {
	req, err := c.preparerForRegistryEnvironmentContainersDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryEnvironmentContainersDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryEnvironmentContainersDeleteThenPoll performs RegistryEnvironmentContainersDelete then polls until it's completed
func (c EnvironmentContainerClient) RegistryEnvironmentContainersDeleteThenPoll(ctx context.Context, id RegistryEnvironmentId) error {
	result, err := c.RegistryEnvironmentContainersDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RegistryEnvironmentContainersDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryEnvironmentContainersDelete: %+v", err)
	}

	return nil
}

// preparerForRegistryEnvironmentContainersDelete prepares the RegistryEnvironmentContainersDelete request.
func (c EnvironmentContainerClient) preparerForRegistryEnvironmentContainersDelete(ctx context.Context, id RegistryEnvironmentId) (*http.Request, error) {
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

// senderForRegistryEnvironmentContainersDelete sends the RegistryEnvironmentContainersDelete request. The method will close the
// http.Response Body if it receives an error.
func (c EnvironmentContainerClient) senderForRegistryEnvironmentContainersDelete(ctx context.Context, req *http.Request) (future RegistryEnvironmentContainersDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
