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

type ApiPortalsValidateDomainOperationResponse struct {
	HttpResponse *http.Response
	Model        *CustomDomainValidateResult
}

// ApiPortalsValidateDomain ...
func (c AppPlatformClient) ApiPortalsValidateDomain(ctx context.Context, id ApiPortalId, input CustomDomainValidatePayload) (result ApiPortalsValidateDomainOperationResponse, err error) {
	req, err := c.preparerForApiPortalsValidateDomain(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsValidateDomain", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsValidateDomain", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForApiPortalsValidateDomain(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "ApiPortalsValidateDomain", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForApiPortalsValidateDomain prepares the ApiPortalsValidateDomain request.
func (c AppPlatformClient) preparerForApiPortalsValidateDomain(ctx context.Context, id ApiPortalId, input CustomDomainValidatePayload) (*http.Request, error) {
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

// responderForApiPortalsValidateDomain handles the response to the ApiPortalsValidateDomain request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForApiPortalsValidateDomain(resp *http.Response) (result ApiPortalsValidateDomainOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
