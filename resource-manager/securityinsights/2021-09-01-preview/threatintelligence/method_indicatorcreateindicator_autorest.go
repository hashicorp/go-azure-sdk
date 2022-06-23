package threatintelligence

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndicatorCreateIndicatorOperationResponse struct {
	HttpResponse *http.Response
	Model        *ThreatIntelligenceInformation
}

// IndicatorCreateIndicator ...
func (c ThreatIntelligenceClient) IndicatorCreateIndicator(ctx context.Context, id WorkspaceId, input ThreatIntelligenceIndicatorModelForRequestBody) (result IndicatorCreateIndicatorOperationResponse, err error) {
	req, err := c.preparerForIndicatorCreateIndicator(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "threatintelligence.ThreatIntelligenceClient", "IndicatorCreateIndicator", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "threatintelligence.ThreatIntelligenceClient", "IndicatorCreateIndicator", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIndicatorCreateIndicator(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "threatintelligence.ThreatIntelligenceClient", "IndicatorCreateIndicator", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIndicatorCreateIndicator prepares the IndicatorCreateIndicator request.
func (c ThreatIntelligenceClient) preparerForIndicatorCreateIndicator(ctx context.Context, id WorkspaceId, input ThreatIntelligenceIndicatorModelForRequestBody) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/providers/Microsoft.SecurityInsights/threatIntelligence/main/createIndicator", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIndicatorCreateIndicator handles the response to the IndicatorCreateIndicator request. The method always
// closes the http.Response Body.
func (c ThreatIntelligenceClient) responderForIndicatorCreateIndicator(resp *http.Response) (result IndicatorCreateIndicatorOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
