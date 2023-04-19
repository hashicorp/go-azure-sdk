package modelcontainer

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

type RegistryModelContainersCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryModelContainersCreateOrUpdate ...
func (c ModelContainerClient) RegistryModelContainersCreateOrUpdate(ctx context.Context, id RegistryModelId, input ModelContainerResource) (result RegistryModelContainersCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRegistryModelContainersCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryModelContainersCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelcontainer.ModelContainerClient", "RegistryModelContainersCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryModelContainersCreateOrUpdateThenPoll performs RegistryModelContainersCreateOrUpdate then polls until it's completed
func (c ModelContainerClient) RegistryModelContainersCreateOrUpdateThenPoll(ctx context.Context, id RegistryModelId, input ModelContainerResource) error {
	result, err := c.RegistryModelContainersCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegistryModelContainersCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryModelContainersCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForRegistryModelContainersCreateOrUpdate prepares the RegistryModelContainersCreateOrUpdate request.
func (c ModelContainerClient) preparerForRegistryModelContainersCreateOrUpdate(ctx context.Context, id RegistryModelId, input ModelContainerResource) (*http.Request, error) {
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

// senderForRegistryModelContainersCreateOrUpdate sends the RegistryModelContainersCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c ModelContainerClient) senderForRegistryModelContainersCreateOrUpdate(ctx context.Context, req *http.Request) (future RegistryModelContainersCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
