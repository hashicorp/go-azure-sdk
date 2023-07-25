package resourceproviders

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MoveOperationResponse struct {
	HttpResponse *http.Response
}

// Move ...
func (c ResourceProvidersClient) Move(ctx context.Context, id commonids.ResourceGroupId, input CsmMoveResourceEnvelope) (result MoveOperationResponse, err error) {
	req, err := c.preparerForMove(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "Move", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "Move", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForMove(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "Move", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForMove prepares the Move request.
func (c ResourceProvidersClient) preparerForMove(ctx context.Context, id commonids.ResourceGroupId, input CsmMoveResourceEnvelope) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/moveResources", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForMove handles the response to the Move request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForMove(resp *http.Response) (result MoveOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
