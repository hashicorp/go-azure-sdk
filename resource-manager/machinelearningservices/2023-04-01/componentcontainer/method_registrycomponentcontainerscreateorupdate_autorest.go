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

type RegistryComponentContainersCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryComponentContainersCreateOrUpdate ...
func (c ComponentContainerClient) RegistryComponentContainersCreateOrUpdate(ctx context.Context, id RegistryComponentId, input ComponentContainerResource) (result RegistryComponentContainersCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRegistryComponentContainersCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentcontainer.ComponentContainerClient", "RegistryComponentContainersCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryComponentContainersCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentcontainer.ComponentContainerClient", "RegistryComponentContainersCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryComponentContainersCreateOrUpdateThenPoll performs RegistryComponentContainersCreateOrUpdate then polls until it's completed
func (c ComponentContainerClient) RegistryComponentContainersCreateOrUpdateThenPoll(ctx context.Context, id RegistryComponentId, input ComponentContainerResource) error {
	result, err := c.RegistryComponentContainersCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegistryComponentContainersCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryComponentContainersCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForRegistryComponentContainersCreateOrUpdate prepares the RegistryComponentContainersCreateOrUpdate request.
func (c ComponentContainerClient) preparerForRegistryComponentContainersCreateOrUpdate(ctx context.Context, id RegistryComponentId, input ComponentContainerResource) (*http.Request, error) {
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

// senderForRegistryComponentContainersCreateOrUpdate sends the RegistryComponentContainersCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c ComponentContainerClient) senderForRegistryComponentContainersCreateOrUpdate(ctx context.Context, req *http.Request) (future RegistryComponentContainersCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
