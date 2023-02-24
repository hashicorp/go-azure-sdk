package manualtrigger

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncidentsRunPlaybookOperationResponse struct {
	HttpResponse *http.Response
	Model        *interface{}
}

// IncidentsRunPlaybook ...
func (c ManualTriggerClient) IncidentsRunPlaybook(ctx context.Context, id IncidentId, input ManualTriggerRequestBody) (result IncidentsRunPlaybookOperationResponse, err error) {
	req, err := c.preparerForIncidentsRunPlaybook(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "manualtrigger.ManualTriggerClient", "IncidentsRunPlaybook", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "manualtrigger.ManualTriggerClient", "IncidentsRunPlaybook", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIncidentsRunPlaybook(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "manualtrigger.ManualTriggerClient", "IncidentsRunPlaybook", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIncidentsRunPlaybook prepares the IncidentsRunPlaybook request.
func (c ManualTriggerClient) preparerForIncidentsRunPlaybook(ctx context.Context, id IncidentId, input ManualTriggerRequestBody) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/runPlaybook", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIncidentsRunPlaybook handles the response to the IncidentsRunPlaybook request. The method always
// closes the http.Response Body.
func (c ManualTriggerClient) responderForIncidentsRunPlaybook(resp *http.Response) (result IncidentsRunPlaybookOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
