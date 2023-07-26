package cognitiveservicescommitmentplans

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommitmentPlansUpdatePlanOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// CommitmentPlansUpdatePlan ...
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansUpdatePlan(ctx context.Context, id CommitmentPlanId, input PatchResourceTagsAndSku) (result CommitmentPlansUpdatePlanOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPatch,
		Path:       id.ID(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// CommitmentPlansUpdatePlanThenPoll performs CommitmentPlansUpdatePlan then polls until it's completed
func (c CognitiveServicesCommitmentPlansClient) CommitmentPlansUpdatePlanThenPoll(ctx context.Context, id CommitmentPlanId, input PatchResourceTagsAndSku) error {
	result, err := c.CommitmentPlansUpdatePlan(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing CommitmentPlansUpdatePlan: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after CommitmentPlansUpdatePlan: %+v", err)
	}

	return nil
}
