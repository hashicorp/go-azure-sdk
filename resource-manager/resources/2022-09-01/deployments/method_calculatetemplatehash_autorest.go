package deployments

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CalculateTemplateHashOperationResponse struct {
	HttpResponse *http.Response
	Model        *TemplateHashResult
}

// CalculateTemplateHash ...
func (c DeploymentsClient) CalculateTemplateHash(ctx context.Context, input interface{}) (result CalculateTemplateHashOperationResponse, err error) {
	req, err := c.preparerForCalculateTemplateHash(ctx, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CalculateTemplateHash", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CalculateTemplateHash", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCalculateTemplateHash(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "deployments.DeploymentsClient", "CalculateTemplateHash", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCalculateTemplateHash prepares the CalculateTemplateHash request.
func (c DeploymentsClient) preparerForCalculateTemplateHash(ctx context.Context, input interface{}) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.Resources/calculateTemplateHash"),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCalculateTemplateHash handles the response to the CalculateTemplateHash request. The method always
// closes the http.Response Body.
func (c DeploymentsClient) responderForCalculateTemplateHash(resp *http.Response) (result CalculateTemplateHashOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
