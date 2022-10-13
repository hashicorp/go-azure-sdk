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

type BuildServiceBuilderDeleteOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// BuildServiceBuilderDelete ...
func (c AppPlatformClient) BuildServiceBuilderDelete(ctx context.Context, id BuilderId) (result BuildServiceBuilderDeleteOperationResponse, err error) {
	req, err := c.preparerForBuildServiceBuilderDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderDelete", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForBuildServiceBuilderDelete(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "BuildServiceBuilderDelete", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// BuildServiceBuilderDeleteThenPoll performs BuildServiceBuilderDelete then polls until it's completed
func (c AppPlatformClient) BuildServiceBuilderDeleteThenPoll(ctx context.Context, id BuilderId) error {
	result, err := c.BuildServiceBuilderDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing BuildServiceBuilderDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after BuildServiceBuilderDelete: %+v", err)
	}

	return nil
}

// preparerForBuildServiceBuilderDelete prepares the BuildServiceBuilderDelete request.
func (c AppPlatformClient) preparerForBuildServiceBuilderDelete(ctx context.Context, id BuilderId) (*http.Request, error) {
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

// senderForBuildServiceBuilderDelete sends the BuildServiceBuilderDelete request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForBuildServiceBuilderDelete(ctx context.Context, req *http.Request) (future BuildServiceBuilderDeleteOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
