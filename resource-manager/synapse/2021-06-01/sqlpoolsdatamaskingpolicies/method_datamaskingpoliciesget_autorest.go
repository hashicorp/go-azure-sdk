package sqlpoolsdatamaskingpolicies

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMaskingPoliciesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *DataMaskingPolicy
}

// DataMaskingPoliciesGet ...
func (c SqlPoolsDataMaskingPoliciesClient) DataMaskingPoliciesGet(ctx context.Context, id SqlPoolId) (result DataMaskingPoliciesGetOperationResponse, err error) {
	req, err := c.preparerForDataMaskingPoliciesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingpolicies.SqlPoolsDataMaskingPoliciesClient", "DataMaskingPoliciesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingpolicies.SqlPoolsDataMaskingPoliciesClient", "DataMaskingPoliciesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDataMaskingPoliciesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingpolicies.SqlPoolsDataMaskingPoliciesClient", "DataMaskingPoliciesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDataMaskingPoliciesGet prepares the DataMaskingPoliciesGet request.
func (c SqlPoolsDataMaskingPoliciesClient) preparerForDataMaskingPoliciesGet(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/dataMaskingPolicies/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDataMaskingPoliciesGet handles the response to the DataMaskingPoliciesGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsDataMaskingPoliciesClient) responderForDataMaskingPoliciesGet(resp *http.Response) (result DataMaskingPoliciesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
