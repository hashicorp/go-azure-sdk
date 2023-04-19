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

type RegistryCodeContainersCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryCodeContainersCreateOrUpdate ...
func (c CodeContainerClient) RegistryCodeContainersCreateOrUpdate(ctx context.Context, id RegistryCodeId, input CodeContainerResource) (result RegistryCodeContainersCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRegistryCodeContainersCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codecontainer.CodeContainerClient", "RegistryCodeContainersCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryCodeContainersCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codecontainer.CodeContainerClient", "RegistryCodeContainersCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryCodeContainersCreateOrUpdateThenPoll performs RegistryCodeContainersCreateOrUpdate then polls until it's completed
func (c CodeContainerClient) RegistryCodeContainersCreateOrUpdateThenPoll(ctx context.Context, id RegistryCodeId, input CodeContainerResource) error {
	result, err := c.RegistryCodeContainersCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegistryCodeContainersCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryCodeContainersCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForRegistryCodeContainersCreateOrUpdate prepares the RegistryCodeContainersCreateOrUpdate request.
func (c CodeContainerClient) preparerForRegistryCodeContainersCreateOrUpdate(ctx context.Context, id RegistryCodeId, input CodeContainerResource) (*http.Request, error) {
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

// senderForRegistryCodeContainersCreateOrUpdate sends the RegistryCodeContainersCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c CodeContainerClient) senderForRegistryCodeContainersCreateOrUpdate(ctx context.Context, req *http.Request) (future RegistryCodeContainersCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
