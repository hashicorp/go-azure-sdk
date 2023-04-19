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

type RegistryModelVersionsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryModelVersionsCreateOrUpdate ...
func (c ModelVersionClient) RegistryModelVersionsCreateOrUpdate(ctx context.Context, id RegistryModelVersionId, input ModelVersionResource) (result RegistryModelVersionsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForRegistryModelVersionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryModelVersionsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "modelversion.ModelVersionClient", "RegistryModelVersionsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryModelVersionsCreateOrUpdateThenPoll performs RegistryModelVersionsCreateOrUpdate then polls until it's completed
func (c ModelVersionClient) RegistryModelVersionsCreateOrUpdateThenPoll(ctx context.Context, id RegistryModelVersionId, input ModelVersionResource) error {
	result, err := c.RegistryModelVersionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing RegistryModelVersionsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryModelVersionsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForRegistryModelVersionsCreateOrUpdate prepares the RegistryModelVersionsCreateOrUpdate request.
func (c ModelVersionClient) preparerForRegistryModelVersionsCreateOrUpdate(ctx context.Context, id RegistryModelVersionId, input ModelVersionResource) (*http.Request, error) {
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

// senderForRegistryModelVersionsCreateOrUpdate sends the RegistryModelVersionsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c ModelVersionClient) senderForRegistryModelVersionsCreateOrUpdate(ctx context.Context, req *http.Request) (future RegistryModelVersionsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
