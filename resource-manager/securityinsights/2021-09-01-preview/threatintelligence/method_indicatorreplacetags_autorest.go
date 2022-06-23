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

type IndicatorReplaceTagsOperationResponse struct {
	HttpResponse *http.Response
	Model        *ThreatIntelligenceInformation
}

// IndicatorReplaceTags ...
func (c ThreatIntelligenceClient) IndicatorReplaceTags(ctx context.Context, id IndicatorId, input ThreatIntelligenceIndicatorModelForRequestBody) (result IndicatorReplaceTagsOperationResponse, err error) {
	req, err := c.preparerForIndicatorReplaceTags(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "threatintelligence.ThreatIntelligenceClient", "IndicatorReplaceTags", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "threatintelligence.ThreatIntelligenceClient", "IndicatorReplaceTags", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForIndicatorReplaceTags(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "threatintelligence.ThreatIntelligenceClient", "IndicatorReplaceTags", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForIndicatorReplaceTags prepares the IndicatorReplaceTags request.
func (c ThreatIntelligenceClient) preparerForIndicatorReplaceTags(ctx context.Context, id IndicatorId, input ThreatIntelligenceIndicatorModelForRequestBody) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/replaceTags", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForIndicatorReplaceTags handles the response to the IndicatorReplaceTags request. The method always
// closes the http.Response Body.
func (c ThreatIntelligenceClient) responderForIndicatorReplaceTags(resp *http.Response) (result IndicatorReplaceTagsOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp
	return
}
