package post

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckMigrationNameAvailabilityOperationResponse struct {
	HttpResponse *http.Response
	Model        *MigrationNameAvailabilityResource
}

// CheckMigrationNameAvailability ...
func (c POSTClient) CheckMigrationNameAvailability(ctx context.Context, id FlexibleServerId, input MigrationNameAvailabilityResource) (result CheckMigrationNameAvailabilityOperationResponse, err error) {
	req, err := c.preparerForCheckMigrationNameAvailability(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "CheckMigrationNameAvailability", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "CheckMigrationNameAvailability", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForCheckMigrationNameAvailability(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "post.POSTClient", "CheckMigrationNameAvailability", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForCheckMigrationNameAvailability prepares the CheckMigrationNameAvailability request.
func (c POSTClient) preparerForCheckMigrationNameAvailability(ctx context.Context, id FlexibleServerId, input MigrationNameAvailabilityResource) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/checkMigrationNameAvailability", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForCheckMigrationNameAvailability handles the response to the CheckMigrationNameAvailability request. The method always
// closes the http.Response Body.
func (c POSTClient) responderForCheckMigrationNameAvailability(resp *http.Response) (result CheckMigrationNameAvailabilityOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
