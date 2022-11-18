package scripts

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptExecutionsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ScriptExecution
}

// ScriptExecutionsGet ...
func (c ScriptsClient) ScriptExecutionsGet(ctx context.Context, id ScriptExecutionId) (result ScriptExecutionsGetOperationResponse, err error) {
	req, err := c.preparerForScriptExecutionsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptExecutionsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptExecutionsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForScriptExecutionsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptExecutionsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForScriptExecutionsGet prepares the ScriptExecutionsGet request.
func (c ScriptsClient) preparerForScriptExecutionsGet(ctx context.Context, id ScriptExecutionId) (*http.Request, error) {
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

// responderForScriptExecutionsGet handles the response to the ScriptExecutionsGet request. The method always
// closes the http.Response Body.
func (c ScriptsClient) responderForScriptExecutionsGet(resp *http.Response) (result ScriptExecutionsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
