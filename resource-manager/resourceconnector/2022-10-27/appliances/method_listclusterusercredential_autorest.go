package appliances

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListClusterUserCredentialOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApplianceListCredentialResults
}

// ListClusterUserCredential ...
func (c AppliancesClient) ListClusterUserCredential(ctx context.Context, id ApplianceId) (result ListClusterUserCredentialOperationResponse, err error) {
	req, err := c.preparerForListClusterUserCredential(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appliances.AppliancesClient", "ListClusterUserCredential", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "appliances.AppliancesClient", "ListClusterUserCredential", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListClusterUserCredential(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "appliances.AppliancesClient", "ListClusterUserCredential", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListClusterUserCredential prepares the ListClusterUserCredential request.
func (c AppliancesClient) preparerForListClusterUserCredential(ctx context.Context, id ApplianceId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/listClusterUserCredential", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListClusterUserCredential handles the response to the ListClusterUserCredential request. The method always
// closes the http.Response Body.
func (c AppliancesClient) responderForListClusterUserCredential(resp *http.Response) (result ListClusterUserCredentialOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
