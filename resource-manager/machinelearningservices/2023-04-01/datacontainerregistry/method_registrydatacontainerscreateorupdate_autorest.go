package datacontainerregistry

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

type RegistryDataContainersCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryDataContainersCreateOrUpdate ...
func (c DataContainerRegistryClient) RegistryDataContainersCreateOrUpdate(ctx context.Context, id DataId, input DataContainerResource) (result RegistryDataContainersCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRegistryDataContainersCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacontainerregistry.DataContainerRegistryClient", "RegistryDataContainersCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryDataContainersCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "datacontainerregistry.DataContainerRegistryClient", "RegistryDataContainersCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryDataContainersCreateOrUpdateThenPoll performs RegistryDataContainersCreateOrUpdate then polls until it's completed
func (c DataContainerRegistryClient) RegistryDataContainersCreateOrUpdateThenPoll(ctx context.Context, id DataId, input DataContainerResource) error {
	result, err := c.RegistryDataContainersCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegistryDataContainersCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryDataContainersCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForRegistryDataContainersCreateOrUpdate prepares the RegistryDataContainersCreateOrUpdate request.
func (c DataContainerRegistryClient) preparerForRegistryDataContainersCreateOrUpdate(ctx context.Context, id DataId, input DataContainerResource) (*http.Request, error) {
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

// senderForRegistryDataContainersCreateOrUpdate sends the RegistryDataContainersCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c DataContainerRegistryClient) senderForRegistryDataContainersCreateOrUpdate(ctx context.Context, req *http.Request) (future RegistryDataContainersCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
