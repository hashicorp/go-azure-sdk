package subscriptions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompileQueryOperationResponse struct {
	HttpResponse *http.Response
	Model        *QueryCompilationResult
}

// CompileQuery ...
func (c SubscriptionsClient) CompileQuery(ctx context.Context, id LocationId, input CompileQuery) (result CompileQueryOperationResponse, err error) {
	req, err := c.preparerForCompileQuery(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "CompileQuery", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "CompileQuery", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCompileQuery(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "subscriptions.SubscriptionsClient", "CompileQuery", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCompileQuery prepares the CompileQuery request.
func (c SubscriptionsClient) preparerForCompileQuery(ctx context.Context, id LocationId, input CompileQuery) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/compileQuery", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCompileQuery handles the response to the CompileQuery request. The method always
// closes the http.Response Body.
func (c SubscriptionsClient) responderForCompileQuery(resp *http.Response) (result CompileQueryOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
