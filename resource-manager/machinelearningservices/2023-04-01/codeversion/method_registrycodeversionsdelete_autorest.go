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

type RegistryCodeVersionsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryCodeVersionsDelete ...
func (c CodeVersionClient) RegistryCodeVersionsDelete(ctx context.Context, id RegistryCodeVersionId) (result RegistryCodeVersionsDeleteOperationResponse, err error) {
	req, err := c.preparerForRegistryCodeVersionsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryCodeVersionsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "codeversion.CodeVersionClient", "RegistryCodeVersionsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryCodeVersionsDeleteThenPoll performs RegistryCodeVersionsDelete then polls until it's completed
func (c CodeVersionClient) RegistryCodeVersionsDeleteThenPoll(ctx context.Context, id RegistryCodeVersionId) error {
	result, err := c.RegistryCodeVersionsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RegistryCodeVersionsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryCodeVersionsDelete: %+v", err)
	}

	return nil
}

// preparerForRegistryCodeVersionsDelete prepares the RegistryCodeVersionsDelete request.
func (c CodeVersionClient) preparerForRegistryCodeVersionsDelete(ctx context.Context, id RegistryCodeVersionId) (*http.Request, error) {
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

// senderForRegistryCodeVersionsDelete sends the RegistryCodeVersionsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c CodeVersionClient) senderForRegistryCodeVersionsDelete(ctx context.Context, req *http.Request) (future RegistryCodeVersionsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
