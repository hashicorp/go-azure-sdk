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

type BuildServiceBuilderCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BuildServiceBuilderCreateOrUpdate ...
func (c AppPlatformClient) BuildServiceBuilderCreateOrUpdate(ctx context.Context, id BuilderId, input BuilderResource) (result BuildServiceBuilderCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForBuildServiceBuilderCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBuildServiceBuilderCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BuildServiceBuilderCreateOrUpdateThenPoll performs BuildServiceBuilderCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) BuildServiceBuilderCreateOrUpdateThenPoll(ctx context.Context, id BuilderId, input BuilderResource) error {
	result, err := c.BuildServiceBuilderCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing BuildServiceBuilderCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BuildServiceBuilderCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForBuildServiceBuilderCreateOrUpdate prepares the BuildServiceBuilderCreateOrUpdate request.
func (c AppPlatformClient) preparerForBuildServiceBuilderCreateOrUpdate(ctx context.Context, id BuilderId, input BuilderResource) (*http.Request, error) {
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

// senderForBuildServiceBuilderCreateOrUpdate sends the BuildServiceBuilderCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForBuildServiceBuilderCreateOrUpdate(ctx context.Context, req *http.Request) (future BuildServiceBuilderCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
