package integrationruntime

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthKeysRegenerateOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeAuthKeys
}

// AuthKeysRegenerate ...
func (c IntegrationRuntimeClient) AuthKeysRegenerate(ctx context.Context, id IntegrationRuntimeId, input IntegrationRuntimeRegenerateKeyParameters) (result AuthKeysRegenerateOperationResponse, err error) {
	req, err := c.preparerForAuthKeysRegenerate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "AuthKeysRegenerate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "AuthKeysRegenerate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAuthKeysRegenerate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntime.IntegrationRuntimeClient", "AuthKeysRegenerate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAuthKeysRegenerate prepares the AuthKeysRegenerate request.
func (c IntegrationRuntimeClient) preparerForAuthKeysRegenerate(ctx context.Context, id IntegrationRuntimeId, input IntegrationRuntimeRegenerateKeyParameters) (*http.Request, error) {
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

// responderForAuthKeysRegenerate handles the response to the AuthKeysRegenerate request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimeClient) responderForAuthKeysRegenerate(resp *http.Response) (result AuthKeysRegenerateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
