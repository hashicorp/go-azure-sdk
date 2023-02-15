package groupuser

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckEntityExistsOperationResponse struct {
	HttpResponse *http.Response
}

// CheckEntityExists ...
func (c GroupUserClient) CheckEntityExists(ctx context.Context, id UserId) (result CheckEntityExistsOperationResponse, err error) {
	req, err := c.preparerForCheckEntityExists(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "groupuser.GroupUserClient", "CheckEntityExists", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "groupuser.GroupUserClient", "CheckEntityExists", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckEntityExists(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "groupuser.GroupUserClient", "CheckEntityExists", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckEntityExists prepares the CheckEntityExists request.
func (c GroupUserClient) preparerForCheckEntityExists(ctx context.Context, id UserId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsHead(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCheckEntityExists handles the response to the CheckEntityExists request. The method always
// closes the http.Response Body.
func (c GroupUserClient) responderForCheckEntityExists(resp *http.Response) (result CheckEntityExistsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
