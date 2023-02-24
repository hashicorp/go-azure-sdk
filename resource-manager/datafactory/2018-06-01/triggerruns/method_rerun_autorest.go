package triggerruns

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RerunOperationResponse struct {
	HttpResponse *http.Response
}

// Rerun ...
func (c TriggerrunsClient) Rerun(ctx context.Context, id TriggerRunId) (result RerunOperationResponse, err error) {
	req, err := c.preparerForRerun(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggerruns.TriggerrunsClient", "Rerun", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggerruns.TriggerrunsClient", "Rerun", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRerun(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "triggerruns.TriggerrunsClient", "Rerun", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRerun prepares the Rerun request.
func (c TriggerrunsClient) preparerForRerun(ctx context.Context, id TriggerRunId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/rerun", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRerun handles the response to the Rerun request. The method always
// closes the http.Response Body.
func (c TriggerrunsClient) responderForRerun(resp *http.Response) (result RerunOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
