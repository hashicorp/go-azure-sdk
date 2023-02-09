package consumerinvitation

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RejectInvitationOperationResponse struct {
	HttpResponse *http.Response
	Model        *ConsumerInvitation
}

// RejectInvitation ...
func (c ConsumerInvitationClient) RejectInvitation(ctx context.Context, id LocationId, input ConsumerInvitation) (result RejectInvitationOperationResponse, err error) {
	req, err := c.preparerForRejectInvitation(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumerinvitation.ConsumerInvitationClient", "RejectInvitation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumerinvitation.ConsumerInvitationClient", "RejectInvitation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForRejectInvitation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "consumerinvitation.ConsumerInvitationClient", "RejectInvitation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForRejectInvitation prepares the RejectInvitation request.
func (c ConsumerInvitationClient) preparerForRejectInvitation(ctx context.Context, id LocationId, input ConsumerInvitation) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/rejectInvitation", id.ID())),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForRejectInvitation handles the response to the RejectInvitation request. The method always
// closes the http.Response Body.
func (c ConsumerInvitationClient) responderForRejectInvitation(resp *http.Response) (result RejectInvitationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
