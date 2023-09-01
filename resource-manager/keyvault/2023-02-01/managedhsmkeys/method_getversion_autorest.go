package managedhsmkeys

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVersionOperationResponse struct {
	HttpResponse *http.Response
	Model        *ManagedHsmKey
}

// GetVersion ...
func (c ManagedHsmKeysClient) GetVersion(ctx context.Context, id VersionId) (result GetVersionOperationResponse, err error) {
	req, err := c.preparerForGetVersion(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedhsmkeys.ManagedHsmKeysClient", "GetVersion", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedhsmkeys.ManagedHsmKeysClient", "GetVersion", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetVersion(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "managedhsmkeys.ManagedHsmKeysClient", "GetVersion", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetVersion prepares the GetVersion request.
func (c ManagedHsmKeysClient) preparerForGetVersion(ctx context.Context, id VersionId) (*http.Request, error) {
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

// responderForGetVersion handles the response to the GetVersion request. The method always
// closes the http.Response Body.
func (c ManagedHsmKeysClient) responderForGetVersion(resp *http.Response) (result GetVersionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
