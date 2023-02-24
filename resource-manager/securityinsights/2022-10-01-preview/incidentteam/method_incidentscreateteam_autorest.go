package incidentteam

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentsCreateTeamOperationResponse struct {
	HttpResponse *http.Response
	Model        *TeamInformation
}

// IncidentsCreateTeam ...
func (c IncidentTeamClient) IncidentsCreateTeam(ctx context.Context, id IncidentId, input TeamProperties) (result IncidentsCreateTeamOperationResponse, err error) {
	req, err := c.preparerForIncidentsCreateTeam(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidentteam.IncidentTeamClient", "IncidentsCreateTeam", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidentteam.IncidentTeamClient", "IncidentsCreateTeam", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIncidentsCreateTeam(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidentteam.IncidentTeamClient", "IncidentsCreateTeam", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIncidentsCreateTeam prepares the IncidentsCreateTeam request.
func (c IncidentTeamClient) preparerForIncidentsCreateTeam(ctx context.Context, id IncidentId, input TeamProperties) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/createTeam", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIncidentsCreateTeam handles the response to the IncidentsCreateTeam request. The method always
// closes the http.Response Body.
func (c IncidentTeamClient) responderForIncidentsCreateTeam(resp *http.Response) (result IncidentsCreateTeamOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
