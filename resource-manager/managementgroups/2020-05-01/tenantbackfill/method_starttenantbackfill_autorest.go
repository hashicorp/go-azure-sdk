package tenantbackfill

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartTenantBackfillOperationResponse struct {
	HttpResponse *http.Response
	Model        *TenantBackfillStatusResult
}

// StartTenantBackfill ...
func (c TenantBackfillClient) StartTenantBackfill(ctx context.Context) (result StartTenantBackfillOperationResponse, err error) {
	req, err := c.preparerForStartTenantBackfill(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantbackfill.TenantBackfillClient", "StartTenantBackfill", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantbackfill.TenantBackfillClient", "StartTenantBackfill", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStartTenantBackfill(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "tenantbackfill.TenantBackfillClient", "StartTenantBackfill", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStartTenantBackfill prepares the StartTenantBackfill request.
func (c TenantBackfillClient) preparerForStartTenantBackfill(ctx context.Context) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath("/providers/Microsoft.Management/startTenantBackfill"),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStartTenantBackfill handles the response to the StartTenantBackfill request. The method always
// closes the http.Response Body.
func (c TenantBackfillClient) responderForStartTenantBackfill(resp *http.Response) (result StartTenantBackfillOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
