package entityrelations

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRelationOperationResponse struct {
	HttpResponse *http.Response
	Model        *Relation
}

// GetRelation ...
func (c EntityRelationsClient) GetRelation(ctx context.Context, id RelationId) (result GetRelationOperationResponse, err error) {
	req, err := c.preparerForGetRelation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entityrelations.EntityRelationsClient", "GetRelation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "entityrelations.EntityRelationsClient", "GetRelation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetRelation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entityrelations.EntityRelationsClient", "GetRelation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetRelation prepares the GetRelation request.
func (c EntityRelationsClient) preparerForGetRelation(ctx context.Context, id RelationId) (*http.Request, error) {
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

// responderForGetRelation handles the response to the GetRelation request. The method always
// closes the http.Response Body.
func (c EntityRelationsClient) responderForGetRelation(resp *http.Response) (result GetRelationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
