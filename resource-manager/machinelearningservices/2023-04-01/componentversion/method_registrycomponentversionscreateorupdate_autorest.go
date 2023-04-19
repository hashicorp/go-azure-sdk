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

type RegistryComponentVersionsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryComponentVersionsCreateOrUpdate ...
func (c ComponentVersionClient) RegistryComponentVersionsCreateOrUpdate(ctx context.Context, id RegistryComponentVersionId, input ComponentVersionResource) (result RegistryComponentVersionsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRegistryComponentVersionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryComponentVersionsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "componentversion.ComponentVersionClient", "RegistryComponentVersionsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryComponentVersionsCreateOrUpdateThenPoll performs RegistryComponentVersionsCreateOrUpdate then polls until it's completed
func (c ComponentVersionClient) RegistryComponentVersionsCreateOrUpdateThenPoll(ctx context.Context, id RegistryComponentVersionId, input ComponentVersionResource) error {
	result, err := c.RegistryComponentVersionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegistryComponentVersionsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryComponentVersionsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForRegistryComponentVersionsCreateOrUpdate prepares the RegistryComponentVersionsCreateOrUpdate request.
func (c ComponentVersionClient) preparerForRegistryComponentVersionsCreateOrUpdate(ctx context.Context, id RegistryComponentVersionId, input ComponentVersionResource) (*http.Request, error) {
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

// senderForRegistryComponentVersionsCreateOrUpdate sends the RegistryComponentVersionsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c ComponentVersionClient) senderForRegistryComponentVersionsCreateOrUpdate(ctx context.Context, req *http.Request) (future RegistryComponentVersionsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
