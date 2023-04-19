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

type RegistriesDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistriesDelete ...
func (c RegistryManagementClient) RegistriesDelete(ctx context.Context, id RegistryId) (result RegistriesDeleteOperationResponse, err error) {
	req, err := c.preparerForRegistriesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistriesDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "registrymanagement.RegistryManagementClient", "RegistriesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistriesDeleteThenPoll performs RegistriesDelete then polls until it's completed
func (c RegistryManagementClient) RegistriesDeleteThenPoll(ctx context.Context, id RegistryId) error {
	result, err := c.RegistriesDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RegistriesDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistriesDelete: %+v", err)
	}

	return nil
}

// preparerForRegistriesDelete prepares the RegistriesDelete request.
func (c RegistryManagementClient) preparerForRegistriesDelete(ctx context.Context, id RegistryId) (*http.Request, error) {
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

// senderForRegistriesDelete sends the RegistriesDelete request. The method will close the
// http.Response Body if it receives an error.
func (c RegistryManagementClient) senderForRegistriesDelete(ctx context.Context, req *http.Request) (future RegistriesDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
