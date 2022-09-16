package sqlpoolsdatamaskingrules

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataMaskingRulesCreateOrUpdateOperationResponse struct {
	HttpResponse *http.Response
	Model        *DataMaskingRule
}

// DataMaskingRulesCreateOrUpdate ...
func (c SqlPoolsDataMaskingRulesClient) DataMaskingRulesCreateOrUpdate(ctx context.Context, id RuleId, input DataMaskingRule) (result DataMaskingRulesCreateOrUpdateOperationResponse, err error) {
	req, err := c.preparerForDataMaskingRulesCreateOrUpdate(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingrules.SqlPoolsDataMaskingRulesClient", "DataMaskingRulesCreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingrules.SqlPoolsDataMaskingRulesClient", "DataMaskingRulesCreateOrUpdate", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDataMaskingRulesCreateOrUpdate(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolsdatamaskingrules.SqlPoolsDataMaskingRulesClient", "DataMaskingRulesCreateOrUpdate", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDataMaskingRulesCreateOrUpdate prepares the DataMaskingRulesCreateOrUpdate request.
func (c SqlPoolsDataMaskingRulesClient) preparerForDataMaskingRulesCreateOrUpdate(ctx context.Context, id RuleId, input DataMaskingRule) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDataMaskingRulesCreateOrUpdate handles the response to the DataMaskingRulesCreateOrUpdate request. The method always
// closes the http.Response Body.
func (c SqlPoolsDataMaskingRulesClient) responderForDataMaskingRulesCreateOrUpdate(resp *http.Response) (result DataMaskingRulesCreateOrUpdateOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusCreated, http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
