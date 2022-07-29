package applicationgroup

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOrUpdateApplicationGroupOperationResponse struct {
	HttpResponse *http.Response
	Model        *ApplicationGroup
}

// CreateOrUpdateApplicationGroup ...
func (c ApplicationGroupClient) CreateOrUpdateApplicationGroup(ctx context.Context, id ApplicationGroupId, input ApplicationGroup) (result CreateOrUpdateApplicationGroupOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateApplicationGroup(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationgroup.ApplicationGroupClient", "CreateOrUpdateApplicationGroup", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationgroup.ApplicationGroupClient", "CreateOrUpdateApplicationGroup", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCreateOrUpdateApplicationGroup(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "applicationgroup.ApplicationGroupClient", "CreateOrUpdateApplicationGroup", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCreateOrUpdateApplicationGroup prepares the CreateOrUpdateApplicationGroup request.
func (c ApplicationGroupClient) preparerForCreateOrUpdateApplicationGroup(ctx context.Context, id ApplicationGroupId, input ApplicationGroup) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCreateOrUpdateApplicationGroup handles the response to the CreateOrUpdateApplicationGroup request. The method always
// closes the http.Response Body.
func (c ApplicationGroupClient) responderForCreateOrUpdateApplicationGroup(resp *http.Response) (result CreateOrUpdateApplicationGroupOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
