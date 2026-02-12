package sparkjobdefinitions

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkJobDefinitionRenameSparkJobDefinitionOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// SparkJobDefinitionRenameSparkJobDefinition ...
func (c SparkJobDefinitionsClient) SparkJobDefinitionRenameSparkJobDefinition(ctx context.Context, id SparkJobDefinitionId, input ArtifactRenameRequest) (result SparkJobDefinitionRenameSparkJobDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/rename", id.Path()),
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

	result.Poller, err = dataplane.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// SparkJobDefinitionRenameSparkJobDefinitionThenPoll performs SparkJobDefinitionRenameSparkJobDefinition then polls until it's completed
func (c SparkJobDefinitionsClient) SparkJobDefinitionRenameSparkJobDefinitionThenPoll(ctx context.Context, id SparkJobDefinitionId, input ArtifactRenameRequest) error {
	result, err := c.SparkJobDefinitionRenameSparkJobDefinition(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing SparkJobDefinitionRenameSparkJobDefinition: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after SparkJobDefinitionRenameSparkJobDefinition: %+v", err)
	}

	return nil
}
