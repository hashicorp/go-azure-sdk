package defaultaccount

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SetOperationResponse struct {
	HttpResponse *http.Response
	Model        *DefaultAccountPayload
}

// Set ...
func (c DefaultAccountClient) Set(ctx context.Context, input DefaultAccountPayload) (result SetOperationResponse, err error) {
	req, err := c.preparerForSet(ctx, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "defaultaccount.DefaultAccountClient", "Set", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "defaultaccount.DefaultAccountClient", "Set", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForSet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "defaultaccount.DefaultAccountClient", "Set", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForSet prepares the Set request.
func (c DefaultAccountClient) preparerForSet(ctx context.Context, input DefaultAccountPayload) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.Purview/setDefaultAccount"),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSet handles the response to the Set request. The method always
// closes the http.Response Body.
func (c DefaultAccountClient) responderForSet(resp *http.Response) (result SetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
