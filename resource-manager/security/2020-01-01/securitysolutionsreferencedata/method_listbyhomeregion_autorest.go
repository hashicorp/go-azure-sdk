package securitysolutionsreferencedata

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListByHomeRegionOperationResponse struct {
	HttpResponse *http.Response
	Model        *SecuritySolutionsReferenceDataList
}

// ListByHomeRegion ...
func (c SecuritySolutionsReferenceDataClient) ListByHomeRegion(ctx context.Context, id LocationId) (result ListByHomeRegionOperationResponse, err error) {
	req, err := c.preparerForListByHomeRegion(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitysolutionsreferencedata.SecuritySolutionsReferenceDataClient", "ListByHomeRegion", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitysolutionsreferencedata.SecuritySolutionsReferenceDataClient", "ListByHomeRegion", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForListByHomeRegion(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "securitysolutionsreferencedata.SecuritySolutionsReferenceDataClient", "ListByHomeRegion", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForListByHomeRegion prepares the ListByHomeRegion request.
func (c SecuritySolutionsReferenceDataClient) preparerForListByHomeRegion(ctx context.Context, id LocationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/securitySolutionsReferenceData", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForListByHomeRegion handles the response to the ListByHomeRegion request. The method always
// closes the http.Response Body.
func (c SecuritySolutionsReferenceDataClient) responderForListByHomeRegion(resp *http.Response) (result ListByHomeRegionOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
