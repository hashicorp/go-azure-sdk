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

type AppsValidateDomainOperationResponse struct {
	HttpResponse *http.Response
	Model        *CustomDomainValidateResult
}

// AppsValidateDomain ...
func (c AppPlatformClient) AppsValidateDomain(ctx context.Context, id AppId, input CustomDomainValidatePayload) (result AppsValidateDomainOperationResponse, err error) {
	req, err := c.preparerForAppsValidateDomain(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsValidateDomain", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsValidateDomain", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAppsValidateDomain(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appplatform.AppPlatformClient", "AppsValidateDomain", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAppsValidateDomain prepares the AppsValidateDomain request.
func (c AppPlatformClient) preparerForAppsValidateDomain(ctx context.Context, id AppId, input CustomDomainValidatePayload) (*http.Request, error) {
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

// responderForAppsValidateDomain handles the response to the AppsValidateDomain request. The method always
// closes the http.Response Body.
func (c AppPlatformClient) responderForAppsValidateDomain(resp *http.Response) (result AppsValidateDomainOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
