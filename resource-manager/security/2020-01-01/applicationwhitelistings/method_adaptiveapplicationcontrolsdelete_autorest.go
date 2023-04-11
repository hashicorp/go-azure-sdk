package applicationwhitelistings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdaptiveApplicationControlsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// AdaptiveApplicationControlsDelete ...
func (c ApplicationWhitelistingsClient) AdaptiveApplicationControlsDelete(ctx context.Context, id ApplicationWhitelistingId) (result AdaptiveApplicationControlsDeleteOperationResponse, err error) {
	req, err := c.preparerForAdaptiveApplicationControlsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAdaptiveApplicationControlsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAdaptiveApplicationControlsDelete prepares the AdaptiveApplicationControlsDelete request.
func (c ApplicationWhitelistingsClient) preparerForAdaptiveApplicationControlsDelete(ctx context.Context, id ApplicationWhitelistingId) (*http.Request, error) {
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

// responderForAdaptiveApplicationControlsDelete handles the response to the AdaptiveApplicationControlsDelete request. The method always
// closes the http.Response Body.
func (c ApplicationWhitelistingsClient) responderForAdaptiveApplicationControlsDelete(resp *http.Response) (result AdaptiveApplicationControlsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted, http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
