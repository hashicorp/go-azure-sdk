package webapps

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetFtpAllowedOperationResponse struct {
	HttpResponse *http.Response
	Model        *CsmPublishingCredentialsPoliciesEntity
}

// GetFtpAllowed ...
func (c WebAppsClient) GetFtpAllowed(ctx context.Context, id SiteId) (result GetFtpAllowedOperationResponse, err error) {
	req, err := c.preparerForGetFtpAllowed(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFtpAllowed", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFtpAllowed", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetFtpAllowed(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "GetFtpAllowed", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetFtpAllowed prepares the GetFtpAllowed request.
func (c WebAppsClient) preparerForGetFtpAllowed(ctx context.Context, id SiteId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/basicPublishingCredentialsPolicies/ftp", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetFtpAllowed handles the response to the GetFtpAllowed request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForGetFtpAllowed(resp *http.Response) (result GetFtpAllowedOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
