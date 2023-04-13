package operations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConvertGraphRunbookContentOperationResponse struct {
	HttpResponse *http.Response
	Model        *GraphicalRunbookContent
}

// ConvertGraphRunbookContent ...
func (c OperationsClient) ConvertGraphRunbookContent(ctx context.Context, id AutomationAccountId, input GraphicalRunbookContent) (result ConvertGraphRunbookContentOperationResponse, err error) {
	req, err := c.preparerForConvertGraphRunbookContent(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operations.OperationsClient", "ConvertGraphRunbookContent", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "operations.OperationsClient", "ConvertGraphRunbookContent", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForConvertGraphRunbookContent(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "operations.OperationsClient", "ConvertGraphRunbookContent", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForConvertGraphRunbookContent prepares the ConvertGraphRunbookContent request.
func (c OperationsClient) preparerForConvertGraphRunbookContent(ctx context.Context, id AutomationAccountId, input GraphicalRunbookContent) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/convertGraphRunbookContent", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForConvertGraphRunbookContent handles the response to the ConvertGraphRunbookContent request. The method always
// closes the http.Response Body.
func (c OperationsClient) responderForConvertGraphRunbookContent(resp *http.Response) (result ConvertGraphRunbookContentOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
