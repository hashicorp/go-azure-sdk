package proxyoperations

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspacesListStorageAccountKeysOperationResponse struct {
	HttpResponse *http.Response
	Model        *ListStorageAccountKeysResult
}

// WorkspacesListStorageAccountKeys ...
func (c ProxyOperationsClient) WorkspacesListStorageAccountKeys(ctx context.Context, id WorkspaceId) (result WorkspacesListStorageAccountKeysOperationResponse, err error) {
	req, err := c.preparerForWorkspacesListStorageAccountKeys(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxyoperations.ProxyOperationsClient", "WorkspacesListStorageAccountKeys", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxyoperations.ProxyOperationsClient", "WorkspacesListStorageAccountKeys", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWorkspacesListStorageAccountKeys(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "proxyoperations.ProxyOperationsClient", "WorkspacesListStorageAccountKeys", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWorkspacesListStorageAccountKeys prepares the WorkspacesListStorageAccountKeys request.
func (c ProxyOperationsClient) preparerForWorkspacesListStorageAccountKeys(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listStorageAccountKeys", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWorkspacesListStorageAccountKeys handles the response to the WorkspacesListStorageAccountKeys request. The method always
// closes the http.Response Body.
func (c ProxyOperationsClient) responderForWorkspacesListStorageAccountKeys(resp *http.Response) (result WorkspacesListStorageAccountKeysOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
