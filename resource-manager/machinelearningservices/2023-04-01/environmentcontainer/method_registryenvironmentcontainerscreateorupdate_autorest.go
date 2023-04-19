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

type RegistryEnvironmentContainersCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryEnvironmentContainersCreateOrUpdate ...
func (c EnvironmentContainerClient) RegistryEnvironmentContainersCreateOrUpdate(ctx context.Context, id RegistryEnvironmentId, input EnvironmentContainerResource) (result RegistryEnvironmentContainersCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRegistryEnvironmentContainersCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryEnvironmentContainersCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentcontainer.EnvironmentContainerClient", "RegistryEnvironmentContainersCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryEnvironmentContainersCreateOrUpdateThenPoll performs RegistryEnvironmentContainersCreateOrUpdate then polls until it's completed
func (c EnvironmentContainerClient) RegistryEnvironmentContainersCreateOrUpdateThenPoll(ctx context.Context, id RegistryEnvironmentId, input EnvironmentContainerResource) error {
	result, err := c.RegistryEnvironmentContainersCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegistryEnvironmentContainersCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryEnvironmentContainersCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForRegistryEnvironmentContainersCreateOrUpdate prepares the RegistryEnvironmentContainersCreateOrUpdate request.
func (c EnvironmentContainerClient) preparerForRegistryEnvironmentContainersCreateOrUpdate(ctx context.Context, id RegistryEnvironmentId, input EnvironmentContainerResource) (*http.Request, error) {
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

// senderForRegistryEnvironmentContainersCreateOrUpdate sends the RegistryEnvironmentContainersCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c EnvironmentContainerClient) senderForRegistryEnvironmentContainersCreateOrUpdate(ctx context.Context, req *http.Request) (future RegistryEnvironmentContainersCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
