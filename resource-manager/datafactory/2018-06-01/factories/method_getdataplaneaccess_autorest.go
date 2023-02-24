package factories

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDataPlaneAccessOperationResponse struct {
	HttpResponse *http.Response
	Model        *AccessPolicyResponse
}

// GetDataPlaneAccess ...
func (c FactoriesClient) GetDataPlaneAccess(ctx context.Context, id FactoryId, input UserAccessPolicy) (result GetDataPlaneAccessOperationResponse, err error) {
	req, err := c.preparerForGetDataPlaneAccess(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "factories.FactoriesClient", "GetDataPlaneAccess", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "factories.FactoriesClient", "GetDataPlaneAccess", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetDataPlaneAccess(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "factories.FactoriesClient", "GetDataPlaneAccess", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetDataPlaneAccess prepares the GetDataPlaneAccess request.
func (c FactoriesClient) preparerForGetDataPlaneAccess(ctx context.Context, id FactoryId, input UserAccessPolicy) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getDataPlaneAccess", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetDataPlaneAccess handles the response to the GetDataPlaneAccess request. The method always
// closes the http.Response Body.
func (c FactoriesClient) responderForGetDataPlaneAccess(resp *http.Response) (result GetDataPlaneAccessOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
