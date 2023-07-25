package staticsites

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

type ValidateBackendForBuildOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ValidateBackendForBuild ...
func (c StaticSitesClient) ValidateBackendForBuild(ctx context.Context, id BuildLinkedBackendId, input StaticSiteLinkedBackendARMResource) (result ValidateBackendForBuildOperationResponse, err error) {
	req, err := c.preparerForValidateBackendForBuild(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ValidateBackendForBuild", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForValidateBackendForBuild(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ValidateBackendForBuild", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ValidateBackendForBuildThenPoll performs ValidateBackendForBuild then polls until it's completed
func (c StaticSitesClient) ValidateBackendForBuildThenPoll(ctx context.Context, id BuildLinkedBackendId, input StaticSiteLinkedBackendARMResource) error {
	result, err := c.ValidateBackendForBuild(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ValidateBackendForBuild: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ValidateBackendForBuild: %+v", err)
	}

	return nil
}

// preparerForValidateBackendForBuild prepares the ValidateBackendForBuild request.
func (c StaticSitesClient) preparerForValidateBackendForBuild(ctx context.Context, id BuildLinkedBackendId, input StaticSiteLinkedBackendARMResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/validate", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForValidateBackendForBuild sends the ValidateBackendForBuild request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForValidateBackendForBuild(ctx context.Context, req *http.Request) (future ValidateBackendForBuildOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
