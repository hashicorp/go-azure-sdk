package alertruleincidents

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByAlertRuleOperationResponse struct {
	HttpResponse *http.Response
	Model        *IncidentListResult
}

// ListByAlertRule ...
func (c AlertRuleIncidentsClient) ListByAlertRule(ctx context.Context, id AlertRuleId) (result ListByAlertRuleOperationResponse, err error) {
	req, err := c.preparerForListByAlertRule(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertruleincidents.AlertRuleIncidentsClient", "ListByAlertRule", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertruleincidents.AlertRuleIncidentsClient", "ListByAlertRule", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByAlertRule(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertruleincidents.AlertRuleIncidentsClient", "ListByAlertRule", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByAlertRule prepares the ListByAlertRule request.
func (c AlertRuleIncidentsClient) preparerForListByAlertRule(ctx context.Context, id AlertRuleId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/incidents", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByAlertRule handles the response to the ListByAlertRule request. The method always
// closes the http.Response Body.
func (c AlertRuleIncidentsClient) responderForListByAlertRule(resp *http.Response) (result ListByAlertRuleOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
