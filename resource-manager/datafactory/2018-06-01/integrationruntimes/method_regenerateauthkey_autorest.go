package integrationruntimes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegenerateAuthKeyOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeAuthKeys
}

// RegenerateAuthKey ...
func (c IntegrationRuntimesClient) RegenerateAuthKey(ctx context.Context, id IntegrationRuntimeId, input IntegrationRuntimeRegenerateKeyParameters) (result RegenerateAuthKeyOperationResponse, err error) {
	req, err := c.preparerForRegenerateAuthKey(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "RegenerateAuthKey", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "RegenerateAuthKey", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegenerateAuthKey(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "RegenerateAuthKey", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegenerateAuthKey prepares the RegenerateAuthKey request.
func (c IntegrationRuntimesClient) preparerForRegenerateAuthKey(ctx context.Context, id IntegrationRuntimeId, input IntegrationRuntimeRegenerateKeyParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/regenerateAuthKey", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRegenerateAuthKey handles the response to the RegenerateAuthKey request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimesClient) responderForRegenerateAuthKey(resp *http.Response) (result RegenerateAuthKeyOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
