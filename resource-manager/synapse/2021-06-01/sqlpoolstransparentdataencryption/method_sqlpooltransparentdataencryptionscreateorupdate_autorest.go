package sqlpoolstransparentdataencryption

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolTransparentDataEncryptionsCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *TransparentDataEncryption
}

// SqlPoolTransparentDataEncryptionsCreateOrUpdate ...
func (c SqlPoolsTransparentDataEncryptionClient) SqlPoolTransparentDataEncryptionsCreateOrUpdate(ctx context.Context, id SqlPoolId, input TransparentDataEncryption) (result SqlPoolTransparentDataEncryptionsCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForSqlPoolTransparentDataEncryptionsCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolTransparentDataEncryptionsCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolTransparentDataEncryptionsCreateOrUpdate prepares the SqlPoolTransparentDataEncryptionsCreateOrUpdate request.
func (c SqlPoolsTransparentDataEncryptionClient) preparerForSqlPoolTransparentDataEncryptionsCreateOrUpdate(ctx context.Context, id SqlPoolId, input TransparentDataEncryption) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/transparentDataEncryption/current", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolTransparentDataEncryptionsCreateOrUpdate handles the response to the SqlPoolTransparentDataEncryptionsCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c SqlPoolsTransparentDataEncryptionClient) responderForSqlPoolTransparentDataEncryptionsCreateOrUpdate(resp *http.Response) (result SqlPoolTransparentDataEncryptionsCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
