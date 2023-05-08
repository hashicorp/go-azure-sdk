package policydefinitions

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetBuiltInOperationResponse struct {
	HttpResponse *http.Response
	Model        *PolicyDefinition
}

// GetBuiltIn ...
func (c PolicyDefinitionsClient) GetBuiltIn(ctx context.Context, id PolicyDefinitionId) (result GetBuiltInOperationResponse, err error) {
	req, err := c.preparerForGetBuiltIn(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "GetBuiltIn", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "GetBuiltIn", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetBuiltIn(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "policydefinitions.PolicyDefinitionsClient", "GetBuiltIn", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetBuiltIn prepares the GetBuiltIn request.
func (c PolicyDefinitionsClient) preparerForGetBuiltIn(ctx context.Context, id PolicyDefinitionId) (*http.Request, error) {
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

// responderForGetBuiltIn handles the response to the GetBuiltIn request. The method always
// closes the http.Response Body.
func (c PolicyDefinitionsClient) responderForGetBuiltIn(resp *http.Response) (result GetBuiltInOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
