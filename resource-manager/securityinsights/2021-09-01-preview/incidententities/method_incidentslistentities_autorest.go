package incidententities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentsListEntitiesOperationResponse struct {
	HttpResponse *http.Response
	Model        *IncidentEntitiesResponse
}

// IncidentsListEntities ...
func (c IncidentEntitiesClient) IncidentsListEntities(ctx context.Context, id IncidentId) (result IncidentsListEntitiesOperationResponse, err error) {
	req, err := c.preparerForIncidentsListEntities(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidententities.IncidentEntitiesClient", "IncidentsListEntities", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidententities.IncidentEntitiesClient", "IncidentsListEntities", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIncidentsListEntities(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidententities.IncidentEntitiesClient", "IncidentsListEntities", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIncidentsListEntities prepares the IncidentsListEntities request.
func (c IncidentEntitiesClient) preparerForIncidentsListEntities(ctx context.Context, id IncidentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/entities", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIncidentsListEntities handles the response to the IncidentsListEntities request. The method always
// closes the http.Response Body.
func (c IncidentEntitiesClient) responderForIncidentsListEntities(resp *http.Response) (result IncidentsListEntitiesOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
