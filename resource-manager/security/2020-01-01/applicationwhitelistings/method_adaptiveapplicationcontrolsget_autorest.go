package applicationwhitelistings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdaptiveApplicationControlsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *AdaptiveApplicationControlGroup
}

// AdaptiveApplicationControlsGet ...
func (c ApplicationWhitelistingsClient) AdaptiveApplicationControlsGet(ctx context.Context, id ApplicationWhitelistingId) (result AdaptiveApplicationControlsGetOperationResponse, err error) {
	req, err := c.preparerForAdaptiveApplicationControlsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAdaptiveApplicationControlsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAdaptiveApplicationControlsGet prepares the AdaptiveApplicationControlsGet request.
func (c ApplicationWhitelistingsClient) preparerForAdaptiveApplicationControlsGet(ctx context.Context, id ApplicationWhitelistingId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAdaptiveApplicationControlsGet handles the response to the AdaptiveApplicationControlsGet request. The method always
// closes the http.Response Body.
func (c ApplicationWhitelistingsClient) responderForAdaptiveApplicationControlsGet(resp *http.Response) (result AdaptiveApplicationControlsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
