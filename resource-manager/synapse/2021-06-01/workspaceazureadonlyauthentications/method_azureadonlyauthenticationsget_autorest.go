package workspaceazureadonlyauthentications

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureADOnlyAuthenticationsGetOperationResponse struct {
	HttpResponse *http.Response
	Model        *AzureADOnlyAuthentication
}

// AzureADOnlyAuthenticationsGet ...
func (c WorkspaceAzureADOnlyAuthenticationsClient) AzureADOnlyAuthenticationsGet(ctx context.Context, id WorkspaceId) (result AzureADOnlyAuthenticationsGetOperationResponse, err error) {
	req, err := c.preparerForAzureADOnlyAuthenticationsGet(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient", "AzureADOnlyAuthenticationsGet", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient", "AzureADOnlyAuthenticationsGet", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForAzureADOnlyAuthenticationsGet(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "workspaceazureadonlyauthentications.WorkspaceAzureADOnlyAuthenticationsClient", "AzureADOnlyAuthenticationsGet", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForAzureADOnlyAuthenticationsGet prepares the AzureADOnlyAuthenticationsGet request.
func (c WorkspaceAzureADOnlyAuthenticationsClient) preparerForAzureADOnlyAuthenticationsGet(ctx context.Context, id WorkspaceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/azureADOnlyAuthentications/default", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForAzureADOnlyAuthenticationsGet handles the response to the AzureADOnlyAuthenticationsGet request. The method always
// closes the http.Response Body.
func (c WorkspaceAzureADOnlyAuthenticationsClient) responderForAzureADOnlyAuthenticationsGet(resp *http.Response) (result AzureADOnlyAuthenticationsGetOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
