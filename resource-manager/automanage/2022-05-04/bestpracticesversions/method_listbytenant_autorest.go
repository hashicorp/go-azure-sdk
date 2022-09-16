package bestpracticesversions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByTenantOperationResponse struct {
	HttpResponse *http.Response
	Model        *BestPracticeList
}

// ListByTenant ...
func (c BestPracticesVersionsClient) ListByTenant(ctx context.Context, id BestPracticeId) (result ListByTenantOperationResponse, err error) {
	req, err := c.preparerForListByTenant(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bestpracticesversions.BestPracticesVersionsClient", "ListByTenant", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "bestpracticesversions.BestPracticesVersionsClient", "ListByTenant", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByTenant(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "bestpracticesversions.BestPracticesVersionsClient", "ListByTenant", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByTenant prepares the ListByTenant request.
func (c BestPracticesVersionsClient) preparerForListByTenant(ctx context.Context, id BestPracticeId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/versions", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByTenant handles the response to the ListByTenant request. The method always
// closes the http.Response Body.
func (c BestPracticesVersionsClient) responderForListByTenant(resp *http.Response) (result ListByTenantOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
