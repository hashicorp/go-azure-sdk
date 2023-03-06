package keys

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetVersionOperationResponse struct {
	HttpResponse *http.Response
	Model        *Key
}

// GetVersion ...
func (c KeysClient) GetVersion(ctx context.Context, id commonids.KeyVaultKeyVersionId) (result GetVersionOperationResponse, err error) {
	req, err := c.preparerForGetVersion(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keys.KeysClient", "GetVersion", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "keys.KeysClient", "GetVersion", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetVersion(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "keys.KeysClient", "GetVersion", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetVersion prepares the GetVersion request.
func (c KeysClient) preparerForGetVersion(ctx context.Context, id commonids.KeyVaultKeyVersionId) (*http.Request, error) {
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
func (c KeysClient) responderForGetVersion(resp *http.Response) (result GetVersionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
