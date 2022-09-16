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

type SqlPoolMetadataSyncConfigsCreateOperationResponse struct {
	HttpResponse *http.Response
	Model        *MetadataSyncConfig
}

// SqlPoolMetadataSyncConfigsCreate ...
func (c SqlPoolsClient) SqlPoolMetadataSyncConfigsCreate(ctx context.Context, id SqlPoolId, input MetadataSyncConfig) (result SqlPoolMetadataSyncConfigsCreateOperationResponse, err error) {
	req, err := c.preparerForSqlPoolMetadataSyncConfigsCreate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpools.SqlPoolsClient", "SqlPoolMetadataSyncConfigsCreate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpools.SqlPoolsClient", "SqlPoolMetadataSyncConfigsCreate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolMetadataSyncConfigsCreate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpools.SqlPoolsClient", "SqlPoolMetadataSyncConfigsCreate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolMetadataSyncConfigsCreate prepares the SqlPoolMetadataSyncConfigsCreate request.
func (c SqlPoolsClient) preparerForSqlPoolMetadataSyncConfigsCreate(ctx context.Context, id SqlPoolId, input MetadataSyncConfig) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/metadataSync/config", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolMetadataSyncConfigsCreate handles the response to the SqlPoolMetadataSyncConfigsCreate request. The method always
// closes the http.Response Body.
func (c SqlPoolsClient) responderForSqlPoolMetadataSyncConfigsCreate(resp *http.Response) (result SqlPoolMetadataSyncConfigsCreateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
