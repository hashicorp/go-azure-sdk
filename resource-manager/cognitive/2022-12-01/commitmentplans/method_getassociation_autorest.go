package commitmentplans

import (
	"context"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetAssociationOperationResponse struct {
	HttpResponse *http.Response
	Model        *CommitmentPlanAccountAssociation
}

// GetAssociation ...
func (c CommitmentPlansClient) GetAssociation(ctx context.Context, id AccountAssociationId) (result GetAssociationOperationResponse, err error) {
	req, err := c.preparerForGetAssociation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentPlansClient", "GetAssociation", nil, "Failure preparing request")
		return
	}

	result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentPlansClient", "GetAssociation", result.HttpResponse, "Failure sending request")
		return
	}

	result, err = c.responderForGetAssociation(result.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentPlansClient", "GetAssociation", result.HttpResponse, "Failure responding to request")
		return
	}

	return
}

// preparerForGetAssociation prepares the GetAssociation request.
func (c CommitmentPlansClient) preparerForGetAssociation(ctx context.Context, id AccountAssociationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForGetAssociation handles the response to the GetAssociation request. The method always
// closes the http.Response Body.
func (c CommitmentPlansClient) responderForGetAssociation(resp *http.Response) (result GetAssociationOperationResponse, err error) {
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result.Model),
		autorest.ByClosing())
	result.HttpResponse = resp

	return
}
