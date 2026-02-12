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

type SparkJobDefinitionDeleteSparkJobDefinitionOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// SparkJobDefinitionDeleteSparkJobDefinition ...
func (c SparkJobDefinitionsClient) SparkJobDefinitionDeleteSparkJobDefinition(ctx context.Context, id SparkJobDefinitionId) (result SparkJobDefinitionDeleteSparkJobDefinitionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod: http.MethodDelete,
		Path:       id.Path(),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
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

// SparkJobDefinitionDeleteSparkJobDefinitionThenPoll performs SparkJobDefinitionDeleteSparkJobDefinition then polls until it's completed
func (c SparkJobDefinitionsClient) SparkJobDefinitionDeleteSparkJobDefinitionThenPoll(ctx context.Context, id SparkJobDefinitionId) error {
	result, err := c.SparkJobDefinitionDeleteSparkJobDefinition(ctx, id)
	if err != nil {
		return fmt.Errorf("performing SparkJobDefinitionDeleteSparkJobDefinition: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after SparkJobDefinitionDeleteSparkJobDefinition: %+v", err)
	}

	return nil
}
