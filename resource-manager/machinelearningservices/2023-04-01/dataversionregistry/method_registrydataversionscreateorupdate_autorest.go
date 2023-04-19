package dataversionregistry

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

type RegistryDataVersionsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryDataVersionsCreateOrUpdate ...
func (c DataVersionRegistryClient) RegistryDataVersionsCreateOrUpdate(ctx context.Context, id VersionId, input DataVersionBaseResource) (result RegistryDataVersionsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRegistryDataVersionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryDataVersionsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "dataversionregistry.DataVersionRegistryClient", "RegistryDataVersionsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryDataVersionsCreateOrUpdateThenPoll performs RegistryDataVersionsCreateOrUpdate then polls until it's completed
func (c DataVersionRegistryClient) RegistryDataVersionsCreateOrUpdateThenPoll(ctx context.Context, id VersionId, input DataVersionBaseResource) error {
	result, err := c.RegistryDataVersionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegistryDataVersionsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryDataVersionsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForRegistryDataVersionsCreateOrUpdate prepares the RegistryDataVersionsCreateOrUpdate request.
func (c DataVersionRegistryClient) preparerForRegistryDataVersionsCreateOrUpdate(ctx context.Context, id VersionId, input DataVersionBaseResource) (*http.Request, error) {
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

// senderForRegistryDataVersionsCreateOrUpdate sends the RegistryDataVersionsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c DataVersionRegistryClient) senderForRegistryDataVersionsCreateOrUpdate(ctx context.Context, req *http.Request) (future RegistryDataVersionsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
