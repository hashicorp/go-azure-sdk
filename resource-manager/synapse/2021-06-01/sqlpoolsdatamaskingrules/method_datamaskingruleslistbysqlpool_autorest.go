package sqlpoolsdatamaskingrules

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMaskingRulesListBySqlPoolOperationResponse struct {
	HttpResponse *http.Response
	Model        *DataMaskingRuleListResult
}

// DataMaskingRulesListBySqlPool ...
func (c SqlPoolsDataMaskingRulesClient) DataMaskingRulesListBySqlPool(ctx context.Context, id SqlPoolId) (result DataMaskingRulesListBySqlPoolOperationResponse, err error) {
	req, err := c.preparerForDataMaskingRulesListBySqlPool(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingrules.SqlPoolsDataMaskingRulesClient", "DataMaskingRulesListBySqlPool", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingrules.SqlPoolsDataMaskingRulesClient", "DataMaskingRulesListBySqlPool", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDataMaskingRulesListBySqlPool(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingrules.SqlPoolsDataMaskingRulesClient", "DataMaskingRulesListBySqlPool", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDataMaskingRulesListBySqlPool prepares the DataMaskingRulesListBySqlPool request.
func (c SqlPoolsDataMaskingRulesClient) preparerForDataMaskingRulesListBySqlPool(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/dataMaskingPolicies/default/rules", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDataMaskingRulesListBySqlPool handles the response to the DataMaskingRulesListBySqlPool request. The method always
// closes the http.Response Body.
func (c SqlPoolsDataMaskingRulesClient) responderForDataMaskingRulesListBySqlPool(resp *http.Response) (result DataMaskingRulesListBySqlPoolOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
