package appplatform

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GatewaysValidateDomainOperationResponse struct {
	HttpResponse *http.Response
	Model        *CustomDomainValidateResult
}

// GatewaysValidateDomain ...
func (c AppPlatformClient) GatewaysValidateDomain(ctx context.Context, id GatewayId, input CustomDomainValidatePayload) (result GatewaysValidateDomainOperationResponse, err error) {
	req, err := c.preparerForGatewaysValidateDomain(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysValidateDomain", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysValidateDomain", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGatewaysValidateDomain(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "GatewaysValidateDomain", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGatewaysValidateDomain prepares the GatewaysValidateDomain request.
func (c AppPlatformClient) preparerForGatewaysValidateDomain(ctx context.Context, id GatewayId, input CustomDomainValidatePayload) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/validateDomain", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGatewaysValidateDomain handles the response to the GatewaysValidateDomain request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForGatewaysValidateDomain(resp *http.Response) (result GatewaysValidateDomainOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
