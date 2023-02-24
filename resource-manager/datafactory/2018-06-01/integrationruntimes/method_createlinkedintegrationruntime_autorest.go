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

type CreateLinkedIntegrationRuntimeOperationResponse struct {
	HttpResponse *http.Response
	Model        *IntegrationRuntimeStatusResponse
}

// CreateLinkedIntegrationRuntime ...
func (c IntegrationRuntimesClient) CreateLinkedIntegrationRuntime(ctx context.Context, id IntegrationRuntimeId, input CreateLinkedIntegrationRuntimeRequest) (result CreateLinkedIntegrationRuntimeOperationResponse, err error) {
	req, err := c.preparerForCreateLinkedIntegrationRuntime(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "CreateLinkedIntegrationRuntime", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "CreateLinkedIntegrationRuntime", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateLinkedIntegrationRuntime(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "CreateLinkedIntegrationRuntime", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateLinkedIntegrationRuntime prepares the CreateLinkedIntegrationRuntime request.
func (c IntegrationRuntimesClient) preparerForCreateLinkedIntegrationRuntime(ctx context.Context, id IntegrationRuntimeId, input CreateLinkedIntegrationRuntimeRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/linkedIntegrationRuntime", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateLinkedIntegrationRuntime handles the response to the CreateLinkedIntegrationRuntime request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimesClient) responderForCreateLinkedIntegrationRuntime(resp *http.Response) (result CreateLinkedIntegrationRuntimeOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
