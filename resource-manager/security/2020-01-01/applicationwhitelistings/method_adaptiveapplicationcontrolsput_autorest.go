package applicationwhitelistings

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdaptiveApplicationControlsPutOperationResponse struct {
	HttpResponse *http.Response
	Model        *AdaptiveApplicationControlGroup
}

// AdaptiveApplicationControlsPut ...
func (c ApplicationWhitelistingsClient) AdaptiveApplicationControlsPut(ctx context.Context, id ApplicationWhitelistingId, input AdaptiveApplicationControlGroup) (result AdaptiveApplicationControlsPutOperationResponse, err error) {
	req, err := c.preparerForAdaptiveApplicationControlsPut(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsPut", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsPut", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAdaptiveApplicationControlsPut(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationwhitelistings.ApplicationWhitelistingsClient", "AdaptiveApplicationControlsPut", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAdaptiveApplicationControlsPut prepares the AdaptiveApplicationControlsPut request.
func (c ApplicationWhitelistingsClient) preparerForAdaptiveApplicationControlsPut(ctx context.Context, id ApplicationWhitelistingId, input AdaptiveApplicationControlGroup) (*http.Request, error) {
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

// responderForAdaptiveApplicationControlsPut handles the response to the AdaptiveApplicationControlsPut request. The method always
// closes the http.Response Body.
func (c ApplicationWhitelistingsClient) responderForAdaptiveApplicationControlsPut(resp *http.Response) (result AdaptiveApplicationControlsPutOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
