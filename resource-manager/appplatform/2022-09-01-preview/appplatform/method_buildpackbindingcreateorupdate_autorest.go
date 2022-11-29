package appplatform

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

type BuildpackBindingCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BuildpackBindingCreateOrUpdate ...
func (c AppPlatformClient) BuildpackBindingCreateOrUpdate(ctx context.Context, id BuildPackBindingId, input BuildpackBindingResource) (result BuildpackBindingCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForBuildpackBindingCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildpackBindingCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBuildpackBindingCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildpackBindingCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BuildpackBindingCreateOrUpdateThenPoll performs BuildpackBindingCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) BuildpackBindingCreateOrUpdateThenPoll(ctx context.Context, id BuildPackBindingId, input BuildpackBindingResource) error {
	result, err := c.BuildpackBindingCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BuildpackBindingCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BuildpackBindingCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForBuildpackBindingCreateOrUpdate prepares the BuildpackBindingCreateOrUpdate request.
func (c AppPlatformClient) preparerForBuildpackBindingCreateOrUpdate(ctx context.Context, id BuildPackBindingId, input BuildpackBindingResource) (*http.Request, error) {
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

// senderForBuildpackBindingCreateOrUpdate sends the BuildpackBindingCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForBuildpackBindingCreateOrUpdate(ctx context.Context, req *http.Request) (future BuildpackBindingCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
