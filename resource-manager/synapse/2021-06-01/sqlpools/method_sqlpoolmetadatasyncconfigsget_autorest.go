package sqlpools

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolMetadataSyncConfigsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *MetadataSyncConfig
}

// SqlPoolMetadataSyncConfigsGet ...
func (c SqlPoolsClient) SqlPoolMetadataSyncConfigsGet(ctx context.Context, id SqlPoolId) (result SqlPoolMetadataSyncConfigsGetOperationResponse, err error) {
	req, err := c.preparerForSqlPoolMetadataSyncConfigsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpools.SqlPoolsClient", "SqlPoolMetadataSyncConfigsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpools.SqlPoolsClient", "SqlPoolMetadataSyncConfigsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolMetadataSyncConfigsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpools.SqlPoolsClient", "SqlPoolMetadataSyncConfigsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolMetadataSyncConfigsGet prepares the SqlPoolMetadataSyncConfigsGet request.
func (c SqlPoolsClient) preparerForSqlPoolMetadataSyncConfigsGet(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/metadataSync/config", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolMetadataSyncConfigsGet handles the response to the SqlPoolMetadataSyncConfigsGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsClient) responderForSqlPoolMetadataSyncConfigsGet(resp *http.Response) (result SqlPoolMetadataSyncConfigsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
