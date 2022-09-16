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

type SqlPoolTransparentDataEncryptionsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *TransparentDataEncryption
}

// SqlPoolTransparentDataEncryptionsGet ...
func (c SqlPoolsTransparentDataEncryptionClient) SqlPoolTransparentDataEncryptionsGet(ctx context.Context, id SqlPoolId) (result SqlPoolTransparentDataEncryptionsGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolTransparentDataEncryptionsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolTransparentDataEncryptionsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolTransparentDataEncryptionsGet prepares the SqlPoolTransparentDataEncryptionsGet request.
func (c SqlPoolsTransparentDataEncryptionClient) preparerForSqlPoolTransparentDataEncryptionsGet(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/transparentDataEncryption/current", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolTransparentDataEncryptionsGet handles the response to the SqlPoolTransparentDataEncryptionsGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsTransparentDataEncryptionClient) responderForSqlPoolTransparentDataEncryptionsGet(resp *http.Response) (result SqlPoolTransparentDataEncryptionsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
