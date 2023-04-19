package componentversion

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

type RegistryComponentVersionsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryComponentVersionsDelete ...
func (c ComponentVersionClient) RegistryComponentVersionsDelete(ctx context.Context, id RegistryComponentVersionId) (result RegistryComponentVersionsDeleteOperationResponse, err error) {
	req, err := c.preparerForRegistryComponentVersionsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryComponentVersionsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryComponentVersionsDeleteThenPoll performs RegistryComponentVersionsDelete then polls until it's completed
func (c ComponentVersionClient) RegistryComponentVersionsDeleteThenPoll(ctx context.Context, id RegistryComponentVersionId) error {
	result, err := c.RegistryComponentVersionsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RegistryComponentVersionsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryComponentVersionsDelete: %+v", err)
	}

	return nil
}

// preparerForRegistryComponentVersionsDelete prepares the RegistryComponentVersionsDelete request.
func (c ComponentVersionClient) preparerForRegistryComponentVersionsDelete(ctx context.Context, id RegistryComponentVersionId) (*http.Request, error) {
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

// senderForRegistryComponentVersionsDelete sends the RegistryComponentVersionsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c ComponentVersionClient) senderForRegistryComponentVersionsDelete(ctx context.Context, req *http.Request) (future RegistryComponentVersionsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
