package checknameavailability

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WithLocationExecuteOperationResponse struct {
	HttpResponse *http.Response
	Model        *NameAvailability
}

// WithLocationExecute ...
func (c CheckNameAvailabilityClient) WithLocationExecute(ctx context.Context, id LocationId, input CheckNameAvailabilityRequest) (result WithLocationExecuteOperationResponse, err error) {
	req, err := c.preparerForWithLocationExecute(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checknameavailability.CheckNameAvailabilityClient", "WithLocationExecute", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "checknameavailability.CheckNameAvailabilityClient", "WithLocationExecute", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForWithLocationExecute(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "checknameavailability.CheckNameAvailabilityClient", "WithLocationExecute", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForWithLocationExecute prepares the WithLocationExecute request.
func (c CheckNameAvailabilityClient) preparerForWithLocationExecute(ctx context.Context, id LocationId, input CheckNameAvailabilityRequest) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/checkNameAvailability", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForWithLocationExecute handles the response to the WithLocationExecute request. The method always
// closes the http.Response Body.
func (c CheckNameAvailabilityClient) responderForWithLocationExecute(resp *http.Response) (result WithLocationExecuteOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
