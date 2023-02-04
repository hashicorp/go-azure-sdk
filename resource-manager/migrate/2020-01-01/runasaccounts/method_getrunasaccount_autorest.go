package runasaccounts

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetRunAsAccountOperationResponse struct {
	HttpResponse *http.Response
	Model        *VMwareRunAsAccount
}

// GetRunAsAccount ...
func (c RunAsAccountsClient) GetRunAsAccount(ctx context.Context, id commonids.VMwareSiteRunAsAccountId) (result GetRunAsAccountOperationResponse, err error) {
	req, err := c.preparerForGetRunAsAccount(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "runasaccounts.RunAsAccountsClient", "GetRunAsAccount", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "runasaccounts.RunAsAccountsClient", "GetRunAsAccount", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetRunAsAccount(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "runasaccounts.RunAsAccountsClient", "GetRunAsAccount", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetRunAsAccount prepares the GetRunAsAccount request.
func (c RunAsAccountsClient) preparerForGetRunAsAccount(ctx context.Context, id commonids.VMwareSiteRunAsAccountId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetRunAsAccount handles the response to the GetRunAsAccount request. The method always
// closes the http.Response Body.
func (c RunAsAccountsClient) responderForGetRunAsAccount(resp *http.Response) (result GetRunAsAccountOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
