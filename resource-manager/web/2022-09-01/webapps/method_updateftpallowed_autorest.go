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

type UpdateFtpAllowedOperationResponse struct {
	HttpResponse *http.Response
	Model        *CsmPublishingCredentialsPoliciesEntity
}

// UpdateFtpAllowed ...
func (c WebAppsClient) UpdateFtpAllowed(ctx context.Context, id SiteId, input CsmPublishingCredentialsPoliciesEntity) (result UpdateFtpAllowedOperationResponse, err error) {
	req, err := c.preparerForUpdateFtpAllowed(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateFtpAllowed", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateFtpAllowed", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateFtpAllowed(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateFtpAllowed", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateFtpAllowed prepares the UpdateFtpAllowed request.
func (c WebAppsClient) preparerForUpdateFtpAllowed(ctx context.Context, id SiteId, input CsmPublishingCredentialsPoliciesEntity) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/basicPublishingCredentialsPolicies/ftp", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateFtpAllowed handles the response to the UpdateFtpAllowed request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateFtpAllowed(resp *http.Response) (result UpdateFtpAllowedOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
