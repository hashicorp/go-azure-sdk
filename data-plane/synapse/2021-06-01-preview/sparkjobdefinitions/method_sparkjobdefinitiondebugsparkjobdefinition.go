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

type SparkJobDefinitionDebugSparkJobDefinitionOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *SparkBatchJob
}

// SparkJobDefinitionDebugSparkJobDefinition ...
func (c SparkJobDefinitionsClient) SparkJobDefinitionDebugSparkJobDefinition(ctx context.Context, input SparkJobDefinitionResource) (result SparkJobDefinitionDebugSparkJobDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       "/debugSparkJobDefinition",
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

// SparkJobDefinitionDebugSparkJobDefinitionThenPoll performs SparkJobDefinitionDebugSparkJobDefinition then polls until it's completed
func (c SparkJobDefinitionsClient) SparkJobDefinitionDebugSparkJobDefinitionThenPoll(ctx context.Context, input SparkJobDefinitionResource) error {
	result, err := c.SparkJobDefinitionDebugSparkJobDefinition(ctx, input)
	if err != nil {
		return fmt.Errorf("performing SparkJobDefinitionDebugSparkJobDefinition: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after SparkJobDefinitionDebugSparkJobDefinition: %+v", err)
	}

	return nil
}
