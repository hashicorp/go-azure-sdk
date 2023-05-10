package tenants

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckResourceNameOperationResponse struct {
	HttpResponse *http.Response
	Model        *CheckResourceNameResult
}

// CheckResourceName ...
func (c TenantsClient) CheckResourceName(ctx context.Context, input ResourceName) (result CheckResourceNameOperationResponse, err error) {
	req, err := c.preparerForCheckResourceName(ctx, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenants.TenantsClient", "CheckResourceName", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenants.TenantsClient", "CheckResourceName", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckResourceName(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenants.TenantsClient", "CheckResourceName", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckResourceName prepares the CheckResourceName request.
func (c TenantsClient) preparerForCheckResourceName(ctx context.Context, input ResourceName) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.Resources/checkResourceName"),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCheckResourceName handles the response to the CheckResourceName request. The method always
// closes the http.Response Body.
func (c TenantsClient) responderForCheckResourceName(resp *http.Response) (result CheckResourceNameOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
