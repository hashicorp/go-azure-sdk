package namedvalue

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListValueOperationResponse struct {
	HttpResponse *http.Response
	Model        *NamedValueSecretContract
}

// ListValue ...
func (c NamedValueClient) ListValue(ctx context.Context, id NamedValueId) (result ListValueOperationResponse, err error) {
	req, err := c.preparerForListValue(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "namedvalue.NamedValueClient", "ListValue", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "namedvalue.NamedValueClient", "ListValue", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListValue(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "namedvalue.NamedValueClient", "ListValue", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListValue prepares the ListValue request.
func (c NamedValueClient) preparerForListValue(ctx context.Context, id NamedValueId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listValue", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListValue handles the response to the ListValue request. The method always
// closes the http.Response Body.
func (c NamedValueClient) responderForListValue(resp *http.Response) (result ListValueOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
