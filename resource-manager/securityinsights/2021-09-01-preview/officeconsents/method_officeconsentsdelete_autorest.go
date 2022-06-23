package officeconsents

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OfficeConsentsDeleteOperationResponse struct {
	HttpResponse *http.Response
}

// OfficeConsentsDelete ...
func (c OfficeConsentsClient) OfficeConsentsDelete(ctx context.Context, id OfficeConsentId) (result OfficeConsentsDeleteOperationResponse, err error) {
	req, err := c.preparerForOfficeConsentsDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForOfficeConsentsDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "officeconsents.OfficeConsentsClient", "OfficeConsentsDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForOfficeConsentsDelete prepares the OfficeConsentsDelete request.
func (c OfficeConsentsClient) preparerForOfficeConsentsDelete(ctx context.Context, id OfficeConsentId) (*http.Request, error) {
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

// responderForOfficeConsentsDelete handles the response to the OfficeConsentsDelete request. The method always
// closes the http.Response Body.
func (c OfficeConsentsClient) responderForOfficeConsentsDelete(resp *http.Response) (result OfficeConsentsDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
