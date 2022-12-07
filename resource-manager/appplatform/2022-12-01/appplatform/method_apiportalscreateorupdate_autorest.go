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

type ApiPortalsCreateOrUpdateOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ApiPortalsCreateOrUpdate ...
func (c AppPlatformClient) ApiPortalsCreateOrUpdate(ctx context.Context, id ApiPortalId, input ApiPortalResource) (result ApiPortalsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForApiPortalsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForApiPortalsCreateOrUpdate(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ApiPortalsCreateOrUpdateThenPoll performs ApiPortalsCreateOrUpdate then polls until it's completed
func (c AppPlatformClient) ApiPortalsCreateOrUpdateThenPoll(ctx context.Context, id ApiPortalId, input ApiPortalResource) error {
	result, err := c.ApiPortalsCreateOrUpdate(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ApiPortalsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ApiPortalsCreateOrUpdate: %+v", err)
	}

	return nil
}

// preparerForApiPortalsCreateOrUpdate prepares the ApiPortalsCreateOrUpdate request.
func (c AppPlatformClient) preparerForApiPortalsCreateOrUpdate(ctx context.Context, id ApiPortalId, input ApiPortalResource) (*http.Request, error) {
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

// senderForApiPortalsCreateOrUpdate sends the ApiPortalsCreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (c AppPlatformClient) senderForApiPortalsCreateOrUpdate(ctx context.Context, req *http.Request) (future ApiPortalsCreateOrUpdateOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
