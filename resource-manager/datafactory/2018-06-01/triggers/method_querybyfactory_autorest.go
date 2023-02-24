package triggers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryByFactoryOperationResponse struct {
	HttpResponse *http.Response
	Model        *TriggerQueryResponse
}

// QueryByFactory ...
func (c TriggersClient) QueryByFactory(ctx context.Context, id FactoryId, input TriggerFilterParameters) (result QueryByFactoryOperationResponse, err error) {
	req, err := c.preparerForQueryByFactory(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "QueryByFactory", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "QueryByFactory", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForQueryByFactory(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggers.TriggersClient", "QueryByFactory", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForQueryByFactory prepares the QueryByFactory request.
func (c TriggersClient) preparerForQueryByFactory(ctx context.Context, id FactoryId, input TriggerFilterParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/querytriggers", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForQueryByFactory handles the response to the QueryByFactory request. The method always
// closes the http.Response Body.
func (c TriggersClient) responderForQueryByFactory(resp *http.Response) (result QueryByFactoryOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
