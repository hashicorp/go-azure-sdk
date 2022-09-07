package entities

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTimelinelistOperationResponse struct {
	HttpResponse *http.Response
	Model        *EntityTimelineResponse
}

// GetTimelinelist ...
func (c EntitiesClient) GetTimelinelist(ctx context.Context, id EntityId, input EntityTimelineParameters) (result GetTimelinelistOperationResponse, err error) {
	req, err := c.preparerForGetTimelinelist(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entities.EntitiesClient", "GetTimelinelist", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "entities.EntitiesClient", "GetTimelinelist", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetTimelinelist(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "entities.EntitiesClient", "GetTimelinelist", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetTimelinelist prepares the GetTimelinelist request.
func (c EntitiesClient) preparerForGetTimelinelist(ctx context.Context, id EntityId, input EntityTimelineParameters) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/getTimeline", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetTimelinelist handles the response to the GetTimelinelist request. The method always
// closes the http.Response Body.
func (c EntitiesClient) responderForGetTimelinelist(resp *http.Response) (result GetTimelinelistOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
