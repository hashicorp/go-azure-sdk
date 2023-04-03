package alertrulesapis

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AlertRulesUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *AlertRuleResource
}

// AlertRulesUpdate ...
func (c AlertRulesAPIsClient) AlertRulesUpdate(ctx context.Context, id AlertRuleId, input AlertRuleResourcePatch) (result AlertRulesUpdateOperationResponse, err error) {
	req, err := c.preparerForAlertRulesUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertrulesapis.AlertRulesAPIsClient", "AlertRulesUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertrulesapis.AlertRulesAPIsClient", "AlertRulesUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAlertRulesUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "alertrulesapis.AlertRulesAPIsClient", "AlertRulesUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAlertRulesUpdate prepares the AlertRulesUpdate request.
func (c AlertRulesAPIsClient) preparerForAlertRulesUpdate(ctx context.Context, id AlertRuleId, input AlertRuleResourcePatch) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPatch(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAlertRulesUpdate handles the response to the AlertRulesUpdate request. The method always
// closes the http.Response Body.
func (c AlertRulesAPIsClient) responderForAlertRulesUpdate(resp *http.Response) (result AlertRulesUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
