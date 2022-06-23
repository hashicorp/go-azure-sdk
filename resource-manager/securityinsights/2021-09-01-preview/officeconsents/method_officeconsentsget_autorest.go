package officeconsents

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeConsentsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *OfficeConsent
}

// OfficeConsentsGet ...
func (c OfficeConsentsClient) OfficeConsentsGet(ctx context.Context, id OfficeConsentId) (result OfficeConsentsGetOperationResponse, err error) {
	req, err := c.preparerForOfficeConsentsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForOfficeConsentsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForOfficeConsentsGet prepares the OfficeConsentsGet request.
func (c OfficeConsentsClient) preparerForOfficeConsentsGet(ctx context.Context, id OfficeConsentId) (*http.Request, error) {
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

// responderForOfficeConsentsGet handles the response to the OfficeConsentsGet request. The method always
// closes the http.Response Body.
func (c OfficeConsentsClient) responderForOfficeConsentsGet(resp *http.Response) (result OfficeConsentsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
