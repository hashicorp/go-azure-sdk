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

type RemoveLinksOperationResponse struct {
	HttpResponse *http.Response
}

// RemoveLinks ...
func (c IntegrationRuntimesClient) RemoveLinks(ctx context.Context, id IntegrationRuntimeId, input LinkedIntegrationRuntimeRequest) (result RemoveLinksOperationResponse, err error) {
	req, err := c.preparerForRemoveLinks(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "RemoveLinks", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "RemoveLinks", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRemoveLinks(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "integrationruntimes.IntegrationRuntimesClient", "RemoveLinks", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRemoveLinks prepares the RemoveLinks request.
func (c IntegrationRuntimesClient) preparerForRemoveLinks(ctx context.Context, id IntegrationRuntimeId, input LinkedIntegrationRuntimeRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/removeLinks", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRemoveLinks handles the response to the RemoveLinks request. The method always
// closes the http.Response Body.
func (c IntegrationRuntimesClient) responderForRemoveLinks(resp *http.Response) (result RemoveLinksOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
