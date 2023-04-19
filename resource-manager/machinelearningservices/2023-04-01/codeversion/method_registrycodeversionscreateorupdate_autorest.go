package codeversion

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

type RegistryCodeVersionsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryCodeVersionsCreateOrUpdate ...
func (c CodeVersionClient) RegistryCodeVersionsCreateOrUpdate(ctx context.Context, id RegistryCodeVersionId, input CodeVersionResource) (result RegistryCodeVersionsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRegistryCodeVersionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryCodeVersionsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryCodeVersionsCreateOrUpdateThenPoll performs RegistryCodeVersionsCreateOrUpdate then polls until it's completed
func (c CodeVersionClient) RegistryCodeVersionsCreateOrUpdateThenPoll(ctx context.Context, id RegistryCodeVersionId, input CodeVersionResource) error {
	result, err := c.RegistryCodeVersionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegistryCodeVersionsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryCodeVersionsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForRegistryCodeVersionsCreateOrUpdate prepares the RegistryCodeVersionsCreateOrUpdate request.
func (c CodeVersionClient) preparerForRegistryCodeVersionsCreateOrUpdate(ctx context.Context, id RegistryCodeVersionId, input CodeVersionResource) (*http.Request, error) {
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

// senderForRegistryCodeVersionsCreateOrUpdate sends the RegistryCodeVersionsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c CodeVersionClient) senderForRegistryCodeVersionsCreateOrUpdate(ctx context.Context, req *http.Request) (future RegistryCodeVersionsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
