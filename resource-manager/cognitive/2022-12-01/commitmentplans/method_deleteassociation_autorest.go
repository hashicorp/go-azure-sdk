package commitmentplans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/hashicorp/go-azure-helpers/polling"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeleteAssociationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// DeleteAssociation ...
func (c CommitmentPlansClient) DeleteAssociation(ctx context.Context, id AccountAssociationId) (result DeleteAssociationOperationResponse, err error) {
	req, err := c.preparerForDeleteAssociation(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentPlansClient", "DeleteAssociation", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForDeleteAssociation(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentPlansClient", "DeleteAssociation", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// DeleteAssociationThenPoll performs DeleteAssociation then polls until it's completed
func (c CommitmentPlansClient) DeleteAssociationThenPoll(ctx context.Context, id AccountAssociationId) error {
	result, err := c.DeleteAssociation(ctx, id)
	if err != nil {
		return fmt.Errorf("performing DeleteAssociation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after DeleteAssociation: %+v", err)
	}

	return nil
}

// preparerForDeleteAssociation prepares the DeleteAssociation request.
func (c CommitmentPlansClient) preparerForDeleteAssociation(ctx context.Context, id AccountAssociationId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsDelete(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForDeleteAssociation sends the DeleteAssociation request. The method will close the
// http.Response Body if it receives an error.
func (c CommitmentPlansClient) senderForDeleteAssociation(ctx context.Context, req *http.Request) (future DeleteAssociationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
