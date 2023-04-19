package registrymanagement

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

type RegistriesCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistriesCreateOrUpdate ...
func (c RegistryManagementClient) RegistriesCreateOrUpdate(ctx context.Context, id RegistryId, input RegistryTrackedResource) (result RegistriesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRegistriesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistriesCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistriesCreateOrUpdateThenPoll performs RegistriesCreateOrUpdate then polls until it's completed
func (c RegistryManagementClient) RegistriesCreateOrUpdateThenPoll(ctx context.Context, id RegistryId, input RegistryTrackedResource) error {
	result, err := c.RegistriesCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegistriesCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistriesCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForRegistriesCreateOrUpdate prepares the RegistriesCreateOrUpdate request.
func (c RegistryManagementClient) preparerForRegistriesCreateOrUpdate(ctx context.Context, id RegistryId, input RegistryTrackedResource) (*http.Request, error) {
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

// senderForRegistriesCreateOrUpdate sends the RegistriesCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c RegistryManagementClient) senderForRegistriesCreateOrUpdate(ctx context.Context, req *http.Request) (future RegistriesCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
