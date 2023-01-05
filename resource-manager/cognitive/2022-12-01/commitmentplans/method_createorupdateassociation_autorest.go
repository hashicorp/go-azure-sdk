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

type CreateOrUpdateAssociationOperationResponse struct {
	Poller       polling.LongRunningPoller
	HttpResponse *http.Response
}

// CreateOrUpdateAssociation ...
func (c CommitmentPlansClient) CreateOrUpdateAssociation(ctx context.Context, id AccountAssociationId, input CommitmentPlanAccountAssociation) (result CreateOrUpdateAssociationOperationResponse, err error) {
	req, err := c.preparerForCreateOrUpdateAssociation(ctx, id, input)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentPlansClient", "CreateOrUpdateAssociation", nil, "Failure preparing request")
		return
	}

	result, err = c.senderForCreateOrUpdateAssociation(ctx, req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "commitmentplans.CommitmentPlansClient", "CreateOrUpdateAssociation", result.HttpResponse, "Failure sending request")
		return
	}

	return
}

// CreateOrUpdateAssociationThenPoll performs CreateOrUpdateAssociation then polls until it's completed
func (c CommitmentPlansClient) CreateOrUpdateAssociationThenPoll(ctx context.Context, id AccountAssociationId, input CommitmentPlanAccountAssociation) error {
	result, err := c.CreateOrUpdateAssociation(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CreateOrUpdateAssociation: %+v", err)
	}

	if err := result.Poller.PollUntilDone(); err != nil {
		return fmt.Errorf("polling after CreateOrUpdateAssociation: %+v", err)
	}

	return nil
}

// preparerForCreateOrUpdateAssociation prepares the CreateOrUpdateAssociation request.
func (c CommitmentPlansClient) preparerForCreateOrUpdateAssociation(ctx context.Context, id AccountAssociationId, input CommitmentPlanAccountAssociation) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPut(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(id.ID()),
		autorest.WithJSON(input),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// senderForCreateOrUpdateAssociation sends the CreateOrUpdateAssociation request. The method will close the
// http.Response Body if it receives an error.
func (c CommitmentPlansClient) senderForCreateOrUpdateAssociation(ctx context.Context, req *http.Request) (future CreateOrUpdateAssociationOperationResponse, err error) {
	var resp *http.Response
	resp, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		return
	}

	future.Poller, err = polling.NewPollerFromResponse(ctx, resp, c.Client, req.Method)
	return
}
