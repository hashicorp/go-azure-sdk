package recommendations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DisableAllForWebAppOperationResponse struct {
	HttpResponse *http.Response
}

// DisableAllForWebApp ...
func (c RecommendationsClient) DisableAllForWebApp(ctx context.Context, id SiteId) (result DisableAllForWebAppOperationResponse, err error) {
	req, err := c.preparerForDisableAllForWebApp(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableAllForWebApp", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableAllForWebApp", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForDisableAllForWebApp(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "recommendations.RecommendationsClient", "DisableAllForWebApp", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForDisableAllForWebApp prepares the DisableAllForWebApp request.
func (c RecommendationsClient) preparerForDisableAllForWebApp(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/recommendations/disable", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForDisableAllForWebApp handles the response to the DisableAllForWebApp request. The method always
// closes the http.Response Body.
func (c RecommendationsClient) responderForDisableAllForWebApp(resp *http.Response) (result DisableAllForWebAppOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
