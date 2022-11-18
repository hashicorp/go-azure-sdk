package location

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetQuotasOperationResponse struct {
	HttpResponse *http.Response
	Model        *BatchLocationQuota
}

// GetQuotas ...
func (c LocationClient) GetQuotas(ctx context.Context, id LocationId) (result GetQuotasOperationResponse, err error) {
	req, err := c.preparerForGetQuotas(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.LocationClient", "GetQuotas", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.LocationClient", "GetQuotas", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetQuotas(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "location.LocationClient", "GetQuotas", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetQuotas prepares the GetQuotas request.
func (c LocationClient) preparerForGetQuotas(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/quotas", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetQuotas handles the response to the GetQuotas request. The method always
// closes the http.Response Body.
func (c LocationClient) responderForGetQuotas(resp *http.Response) (result GetQuotasOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
