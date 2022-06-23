package automationrules

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationRulesCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *AutomationRule
}

// AutomationRulesCreateOrUpdate ...
func (c AutomationRulesClient) AutomationRulesCreateOrUpdate(ctx context.Context, id AutomationRuleId, input AutomationRule) (result AutomationRulesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForAutomationRulesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAutomationRulesCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAutomationRulesCreateOrUpdate prepares the AutomationRulesCreateOrUpdate request.
func (c AutomationRulesClient) preparerForAutomationRulesCreateOrUpdate(ctx context.Context, id AutomationRuleId, input AutomationRule) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAutomationRulesCreateOrUpdate handles the response to the AutomationRulesCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c AutomationRulesClient) responderForAutomationRulesCreateOrUpdate(resp *http.Response) (result AutomationRulesCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
