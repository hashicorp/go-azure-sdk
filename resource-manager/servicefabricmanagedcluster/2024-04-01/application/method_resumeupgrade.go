package application

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

type ResumeUpgradeOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// ResumeUpgrade ...
func (c ApplicationClient) ResumeUpgrade(ctx context.Context, id ApplicationId, input RuntimeResumeApplicationUpgradeParameters) (result ResumeUpgradeOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
		},
		HttpMethod: http.MethodPost,

		Path: fmt.Sprintf("%s/resumeUpgrade", id.ID()),
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

// ResumeUpgradeThenPoll performs ResumeUpgrade then polls until it's completed
func (c ApplicationClient) ResumeUpgradeThenPoll(ctx context.Context, id ApplicationId, input RuntimeResumeApplicationUpgradeParameters) error {
	result, err := c.ResumeUpgrade(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing ResumeUpgrade: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ResumeUpgrade: %+v", err)
	}

	return nil
}
