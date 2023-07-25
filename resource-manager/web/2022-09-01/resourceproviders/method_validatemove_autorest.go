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

type ValidateMoveOperationResponse struct {
	HttpResponse *http.Response
}

// ValidateMove ...
func (c ResourceProvidersClient) ValidateMove(ctx context.Context, id commonids.ResourceGroupId, input CsmMoveResourceEnvelope) (result ValidateMoveOperationResponse, err error) {
	req, err := c.preparerForValidateMove(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ValidateMove", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ValidateMove", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForValidateMove(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "resourceproviders.ResourceProvidersClient", "ValidateMove", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForValidateMove prepares the ValidateMove request.
func (c ResourceProvidersClient) preparerForValidateMove(ctx context.Context, id commonids.ResourceGroupId, input CsmMoveResourceEnvelope) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/validateMoveResources", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForValidateMove handles the response to the ValidateMove request. The method always
// closes the http.Response Body.
func (c ResourceProvidersClient) responderForValidateMove(resp *http.Response) (result ValidateMoveOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusNoContent),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
