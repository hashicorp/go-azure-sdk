package tenantaccess

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegenerateSecondaryKeyOperationResponse struct {
	HttpResponse *http.Response
}

// RegenerateSecondaryKey ...
func (c TenantAccessClient) RegenerateSecondaryKey(ctx context.Context, id AccesId) (result RegenerateSecondaryKeyOperationResponse, err error) {
	req, err := c.preparerForRegenerateSecondaryKey(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantaccess.TenantAccessClient", "RegenerateSecondaryKey", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantaccess.TenantAccessClient", "RegenerateSecondaryKey", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegenerateSecondaryKey(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantaccess.TenantAccessClient", "RegenerateSecondaryKey", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegenerateSecondaryKey prepares the RegenerateSecondaryKey request.
func (c TenantAccessClient) preparerForRegenerateSecondaryKey(ctx context.Context, id AccesId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/regenerateSecondaryKey", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRegenerateSecondaryKey handles the response to the RegenerateSecondaryKey request. The method always
// closes the http.Response Body.
func (c TenantAccessClient) responderForRegenerateSecondaryKey(resp *http.Response) (result RegenerateSecondaryKeyOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
