package tenantaccessgit

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegeneratePrimaryKeyOperationResponse struct {
	HttpResponse *http.Response
}

// RegeneratePrimaryKey ...
func (c TenantAccessGitClient) RegeneratePrimaryKey(ctx context.Context, id AccesId) (result RegeneratePrimaryKeyOperationResponse, err error) {
	req, err := c.preparerForRegeneratePrimaryKey(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantaccessgit.TenantAccessGitClient", "RegeneratePrimaryKey", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantaccessgit.TenantAccessGitClient", "RegeneratePrimaryKey", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRegeneratePrimaryKey(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantaccessgit.TenantAccessGitClient", "RegeneratePrimaryKey", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRegeneratePrimaryKey prepares the RegeneratePrimaryKey request.
func (c TenantAccessGitClient) preparerForRegeneratePrimaryKey(ctx context.Context, id AccesId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/git/regeneratePrimaryKey", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRegeneratePrimaryKey handles the response to the RegeneratePrimaryKey request. The method always
// closes the http.Response Body.
func (c TenantAccessGitClient) responderForRegeneratePrimaryKey(resp *http.Response) (result RegeneratePrimaryKeyOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
