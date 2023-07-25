package domains

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransferOutOperationResponse struct {
	HttpResponse *http.Response
	Model        *Domain
}

// TransferOut ...
func (c DomainsClient) TransferOut(ctx context.Context, id DomainId) (result TransferOutOperationResponse, err error) {
	req, err := c.preparerForTransferOut(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "TransferOut", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "TransferOut", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForTransferOut(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "domains.DomainsClient", "TransferOut", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForTransferOut prepares the TransferOut request.
func (c DomainsClient) preparerForTransferOut(ctx context.Context, id DomainId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/transferOut", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForTransferOut handles the response to the TransferOut request. The method always
// closes the http.Response Body.
func (c DomainsClient) responderForTransferOut(resp *http.Response) (result TransferOutOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
