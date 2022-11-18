package scripts

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptCmdletsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ScriptCmdlet
}

// ScriptCmdletsGet ...
func (c ScriptsClient) ScriptCmdletsGet(ctx context.Context, id ScriptCmdletId) (result ScriptCmdletsGetOperationResponse, err error) {
	req, err := c.preparerForScriptCmdletsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptCmdletsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptCmdletsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForScriptCmdletsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptCmdletsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForScriptCmdletsGet prepares the ScriptCmdletsGet request.
func (c ScriptsClient) preparerForScriptCmdletsGet(ctx context.Context, id ScriptCmdletId) (*http.Request, error) {
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

// responderForScriptCmdletsGet handles the response to the ScriptCmdletsGet request. The method always
// closes the http.Response Body.
func (c ScriptsClient) responderForScriptCmdletsGet(resp *http.Response) (result ScriptCmdletsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
