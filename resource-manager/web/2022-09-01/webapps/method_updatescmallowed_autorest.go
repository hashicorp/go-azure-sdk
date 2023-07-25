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

type UpdateScmAllowedOperationResponse struct {
	HttpResponse *http.Response
	Model        *CsmPublishingCredentialsPoliciesEntity
}

// UpdateScmAllowed ...
func (c WebAppsClient) UpdateScmAllowed(ctx context.Context, id SiteId, input CsmPublishingCredentialsPoliciesEntity) (result UpdateScmAllowedOperationResponse, err error) {
	req, err := c.preparerForUpdateScmAllowed(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateScmAllowed", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateScmAllowed", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForUpdateScmAllowed(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "webapps.WebAppsClient", "UpdateScmAllowed", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForUpdateScmAllowed prepares the UpdateScmAllowed request.
func (c WebAppsClient) preparerForUpdateScmAllowed(ctx context.Context, id SiteId, input CsmPublishingCredentialsPoliciesEntity) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/basicPublishingCredentialsPolicies/scm", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForUpdateScmAllowed handles the response to the UpdateScmAllowed request. The method always
// closes the http.Response Body.
func (c WebAppsClient) responderForUpdateScmAllowed(resp *http.Response) (result UpdateScmAllowedOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
