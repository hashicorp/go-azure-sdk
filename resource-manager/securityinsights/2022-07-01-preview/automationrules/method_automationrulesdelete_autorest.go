package automationrules

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AutomationRulesDeleteOperationResponse struct {
	HttpResponse *http.Response
	Model        *interface{}
}

// AutomationRulesDelete ...
func (c AutomationRulesClient) AutomationRulesDelete(ctx context.Context, id AutomationRuleId) (result AutomationRulesDeleteOperationResponse, err error) {
	req, err := c.preparerForAutomationRulesDelete(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesDelete", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesDelete", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAutomationRulesDelete(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "automationrules.AutomationRulesClient", "AutomationRulesDelete", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAutomationRulesDelete prepares the AutomationRulesDelete request.
func (c AutomationRulesClient) preparerForAutomationRulesDelete(ctx context.Context, id AutomationRuleId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAutomationRulesDelete handles the response to the AutomationRulesDelete request. The method always
// closes the http.Response Body.
func (c AutomationRulesClient) responderForAutomationRulesDelete(resp *http.Response) (result AutomationRulesDeleteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
