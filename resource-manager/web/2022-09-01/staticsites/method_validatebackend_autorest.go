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

type ValidateBackendOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// ValidateBackend ...
func (c StaticSitesClient) ValidateBackend(ctx context.Context, id LinkedBackendId, input StaticSiteLinkedBackendARMResource) (result ValidateBackendOperationResponse, err error) {
	req, err := c.preparerForValidateBackend(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ValidateBackend", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForValidateBackend(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "staticsites.StaticSitesClient", "ValidateBackend", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// ValidateBackendThenPoll performs ValidateBackend then polls until it's completed
func (c StaticSitesClient) ValidateBackendThenPoll(ctx context.Context, id LinkedBackendId, input StaticSiteLinkedBackendARMResource) error {
	result, err := c.ValidateBackend(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ValidateBackend: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after ValidateBackend: %+v", err)
	}

	return nil
}

// preparerForValidateBackend prepares the ValidateBackend request.
func (c StaticSitesClient) preparerForValidateBackend(ctx context.Context, id LinkedBackendId, input StaticSiteLinkedBackendARMResource) (*http.Request, error) {
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

// senderForValidateBackend sends the ValidateBackend request. The method will close the
// http.Response Body if it receives an error.
func (c StaticSitesClient) senderForValidateBackend(ctx context.Context, req *http.Request) (future ValidateBackendOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
