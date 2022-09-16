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

type RenameOperationResponse struct {
	HttpResponse *http.Response
}

// Rename ...
func (c SqlPoolsClient) Rename(ctx context.Context, id SqlPoolId, input ResourceMoveDefinition) (result RenameOperationResponse, err error) {
	req, err := c.preparerForRename(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpools.SqlPoolsClient", "Rename", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpools.SqlPoolsClient", "Rename", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRename(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpools.SqlPoolsClient", "Rename", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRename prepares the Rename request.
func (c SqlPoolsClient) preparerForRename(ctx context.Context, id SqlPoolId, input ResourceMoveDefinition) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/move", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRename handles the response to the Rename request. The method always
// closes the http.Response Body.
func (c SqlPoolsClient) responderForRename(resp *http.Response) (result RenameOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
