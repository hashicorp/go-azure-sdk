package tenantbackfill

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StatusOperationResponse struct {
	HttpResponse *http.Response
	Model        *TenantBackfillStatusResult
}

// Status ...
func (c TenantBackfillClient) Status(ctx context.Context) (result StatusOperationResponse, err error) {
	req, err := c.preparerForStatus(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantbackfill.TenantBackfillClient", "Status", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantbackfill.TenantBackfillClient", "Status", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStatus(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantbackfill.TenantBackfillClient", "Status", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStatus prepares the Status request.
func (c TenantBackfillClient) preparerForStatus(ctx context.Context) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.Management/tenantBackfillStatus"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStatus handles the response to the Status request. The method always
// closes the http.Response Body.
func (c TenantBackfillClient) responderForStatus(resp *http.Response) (result StatusOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
