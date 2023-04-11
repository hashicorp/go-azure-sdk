package securescore

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecureScoresGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *SecureScoreItem
}

// SecureScoresGet ...
func (c SecureScoreClient) SecureScoresGet(ctx context.Context, id SecureScoreId) (result SecureScoresGetOperationResponse, err error) {
	req, err := c.preparerForSecureScoresGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoresGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoresGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSecureScoresGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securescore.SecureScoreClient", "SecureScoresGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSecureScoresGet prepares the SecureScoresGet request.
func (c SecureScoreClient) preparerForSecureScoresGet(ctx context.Context, id SecureScoreId) (*http.Request, error) {
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

// responderForSecureScoresGet handles the response to the SecureScoresGet request. The method always
// closes the http.Response Body.
func (c SecureScoreClient) responderForSecureScoresGet(resp *http.Response) (result SecureScoresGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
