package incidentalerts

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentsListAlertsOperationResponse struct {
	HttpResponse *http.Response
	Model        *IncidentAlertList
}

// IncidentsListAlerts ...
func (c IncidentAlertsClient) IncidentsListAlerts(ctx context.Context, id IncidentId) (result IncidentsListAlertsOperationResponse, err error) {
	req, err := c.preparerForIncidentsListAlerts(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidentalerts.IncidentAlertsClient", "IncidentsListAlerts", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidentalerts.IncidentAlertsClient", "IncidentsListAlerts", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIncidentsListAlerts(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "incidentalerts.IncidentAlertsClient", "IncidentsListAlerts", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIncidentsListAlerts prepares the IncidentsListAlerts request.
func (c IncidentAlertsClient) preparerForIncidentsListAlerts(ctx context.Context, id IncidentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/alerts", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIncidentsListAlerts handles the response to the IncidentsListAlerts request. The method always
// closes the http.Response Body.
func (c IncidentAlertsClient) responderForIncidentsListAlerts(resp *http.Response) (result IncidentsListAlertsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
