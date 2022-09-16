package sqlpoolsreplicationlinks

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolReplicationLinksGetByNameOperationResponse struct {
	HttpResponse *http.Response
	Model        *ReplicationLink
}

// SqlPoolReplicationLinksGetByName ...
func (c SqlPoolsReplicationLinksClient) SqlPoolReplicationLinksGetByName(ctx context.Context, id ReplicationLinkId) (result SqlPoolReplicationLinksGetByNameOperationResponse, err error) {
	req, err := c.preparerForSqlPoolReplicationLinksGetByName(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsreplicationlinks.SqlPoolsReplicationLinksClient", "SqlPoolReplicationLinksGetByName", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsreplicationlinks.SqlPoolsReplicationLinksClient", "SqlPoolReplicationLinksGetByName", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSqlPoolReplicationLinksGetByName(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsreplicationlinks.SqlPoolsReplicationLinksClient", "SqlPoolReplicationLinksGetByName", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSqlPoolReplicationLinksGetByName prepares the SqlPoolReplicationLinksGetByName request.
func (c SqlPoolsReplicationLinksClient) preparerForSqlPoolReplicationLinksGetByName(ctx context.Context, id ReplicationLinkId) (*http.Request, error) {
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

// responderForSqlPoolReplicationLinksGetByName handles the response to the SqlPoolReplicationLinksGetByName request. The method always
// closes the http.Response Body.
func (c SqlPoolsReplicationLinksClient) responderForSqlPoolReplicationLinksGetByName(resp *http.Response) (result SqlPoolReplicationLinksGetByNameOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
