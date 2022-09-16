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

type DataMaskingPoliciesCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *DataMaskingPolicy
}

// DataMaskingPoliciesCreateOrUpdate ...
func (c SqlPoolsDataMaskingPoliciesClient) DataMaskingPoliciesCreateOrUpdate(ctx context.Context, id SqlPoolId, input DataMaskingPolicy) (result DataMaskingPoliciesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForDataMaskingPoliciesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingpolicies.SqlPoolsDataMaskingPoliciesClient", "DataMaskingPoliciesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingpolicies.SqlPoolsDataMaskingPoliciesClient", "DataMaskingPoliciesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDataMaskingPoliciesCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingpolicies.SqlPoolsDataMaskingPoliciesClient", "DataMaskingPoliciesCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDataMaskingPoliciesCreateOrUpdate prepares the DataMaskingPoliciesCreateOrUpdate request.
func (c SqlPoolsDataMaskingPoliciesClient) preparerForDataMaskingPoliciesCreateOrUpdate(ctx context.Context, id SqlPoolId, input DataMaskingPolicy) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/dataMaskingPolicies/default", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDataMaskingPoliciesCreateOrUpdate handles the response to the DataMaskingPoliciesCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c SqlPoolsDataMaskingPoliciesClient) responderForDataMaskingPoliciesCreateOrUpdate(resp *http.Response) (result DataMaskingPoliciesCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
