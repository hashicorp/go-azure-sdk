package experiments

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartOperationResponse struct {
	HttpResponse *http.Response
	Model        *ExperimentStartOperationResult
}

// Start ...
func (c ExperimentsClient) Start(ctx context.Context, id ExperimentId) (result StartOperationResponse, err error) {
	req, err := c.preparerForStart(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "Start", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "Start", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForStart(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "experiments.ExperimentsClient", "Start", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForStart prepares the Start request.
func (c ExperimentsClient) preparerForStart(ctx context.Context, id ExperimentId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/start", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForStart handles the response to the Start request. The method always
// closes the http.Response Body.
func (c ExperimentsClient) responderForStart(resp *http.Response) (result StartOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
