package environmentversion

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

type RegistryEnvironmentVersionsDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// RegistryEnvironmentVersionsDelete ...
func (c EnvironmentVersionClient) RegistryEnvironmentVersionsDelete(ctx context.Context, id RegistryEnvironmentVersionId) (result RegistryEnvironmentVersionsDeleteOperationResponse, err error) {
	req, err := c.preparerForRegistryEnvironmentVersionsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentversion.EnvironmentVersionClient", "RegistryEnvironmentVersionsDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForRegistryEnvironmentVersionsDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "environmentversion.EnvironmentVersionClient", "RegistryEnvironmentVersionsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// RegistryEnvironmentVersionsDeleteThenPoll performs RegistryEnvironmentVersionsDelete then polls until it's completed
func (c EnvironmentVersionClient) RegistryEnvironmentVersionsDeleteThenPoll(ctx context.Context, id RegistryEnvironmentVersionId) error {
	result, err := c.RegistryEnvironmentVersionsDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing RegistryEnvironmentVersionsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after RegistryEnvironmentVersionsDelete: %+v", err)
	}

	return nil
}

// preparerForRegistryEnvironmentVersionsDelete prepares the RegistryEnvironmentVersionsDelete request.
func (c EnvironmentVersionClient) preparerForRegistryEnvironmentVersionsDelete(ctx context.Context, id RegistryEnvironmentVersionId) (*http.Request, error) {
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

// senderForRegistryEnvironmentVersionsDelete sends the RegistryEnvironmentVersionsDelete request. The method will close the
// http.Response Body if it receives an error.
func (c EnvironmentVersionClient) senderForRegistryEnvironmentVersionsDelete(ctx context.Context, req *http.Request) (future RegistryEnvironmentVersionsDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
