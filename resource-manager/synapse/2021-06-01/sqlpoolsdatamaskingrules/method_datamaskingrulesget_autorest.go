package sqlpoolsdatamaskingrules

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMaskingRulesGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *DataMaskingRule
}

// DataMaskingRulesGet ...
func (c SqlPoolsDataMaskingRulesClient) DataMaskingRulesGet(ctx context.Context, id RuleId) (result DataMaskingRulesGetOperationResponse, err error) {
	req, err := c.preparerForDataMaskingRulesGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingrules.SqlPoolsDataMaskingRulesClient", "DataMaskingRulesGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingrules.SqlPoolsDataMaskingRulesClient", "DataMaskingRulesGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDataMaskingRulesGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingrules.SqlPoolsDataMaskingRulesClient", "DataMaskingRulesGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDataMaskingRulesGet prepares the DataMaskingRulesGet request.
func (c SqlPoolsDataMaskingRulesClient) preparerForDataMaskingRulesGet(ctx context.Context, id RuleId) (*http.Request, error) {
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

// responderForDataMaskingRulesGet handles the response to the DataMaskingRulesGet request. The method always
// closes the http.Response Body.
func (c SqlPoolsDataMaskingRulesClient) responderForDataMaskingRulesGet(resp *http.Response) (result DataMaskingRulesGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
