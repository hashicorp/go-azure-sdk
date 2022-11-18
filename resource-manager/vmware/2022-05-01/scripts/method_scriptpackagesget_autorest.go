package scripts

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScriptPackagesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *ScriptPackage
}

// ScriptPackagesGet ...
func (c ScriptsClient) ScriptPackagesGet(ctx context.Context, id ScriptPackageId) (result ScriptPackagesGetOperationResponse, err error) {
	req, err := c.preparerForScriptPackagesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptPackagesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptPackagesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForScriptPackagesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "scripts.ScriptsClient", "ScriptPackagesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForScriptPackagesGet prepares the ScriptPackagesGet request.
func (c ScriptsClient) preparerForScriptPackagesGet(ctx context.Context, id ScriptPackageId) (*http.Request, error) {
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

// responderForScriptPackagesGet handles the response to the ScriptPackagesGet request. The method always
// closes the http.Response Body.
func (c ScriptsClient) responderForScriptPackagesGet(resp *http.Response) (result ScriptPackagesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
