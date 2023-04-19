package modelversion

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

type RegistryModelVersionsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryModelVersionsDelete ...
func (c ModelVersionClient) RegistryModelVersionsDelete(ctx context.Context, id RegistryModelVersionId) (result RegistryModelVersionsDeleteOperationResponse, err error) {
	req, err := c.preparerForRegistryModelVersionsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryModelVersionsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryModelVersionsDeleteThenPoll performs RegistryModelVersionsDelete then polls until it's completed
func (c ModelVersionClient) RegistryModelVersionsDeleteThenPoll(ctx context.Context, id RegistryModelVersionId) error {
	result, err := c.RegistryModelVersionsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RegistryModelVersionsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryModelVersionsDelete: %+v", err)
	}

	return nil
}

// preparerForRegistryModelVersionsDelete prepares the RegistryModelVersionsDelete request.
func (c ModelVersionClient) preparerForRegistryModelVersionsDelete(ctx context.Context, id RegistryModelVersionId) (*http.Request, error) {
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

// senderForRegistryModelVersionsDelete sends the RegistryModelVersionsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c ModelVersionClient) senderForRegistryModelVersionsDelete(ctx context.Context, req *http.Request) (future RegistryModelVersionsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
