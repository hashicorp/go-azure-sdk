package networkclouds

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

type ConsolesUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *Console
}

type ConsolesUpdateOperationOptions struct {
	IfMatch     *string
	IfNoneMatch *string
}

func DefaultConsolesUpdateOperationOptions() ConsolesUpdateOperationOptions {
	return ConsolesUpdateOperationOptions{}
}

func (o ConsolesUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o ConsolesUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ConsolesUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ConsolesUpdate ...
func (c NetworkcloudsClient) ConsolesUpdate(ctx context.Context, id ConsoleId, input ConsolePatchParameters, options ConsolesUpdateOperationOptions) (result ConsolesUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

// ConsolesUpdateThenPoll performs ConsolesUpdate then polls until it's completed
func (c NetworkcloudsClient) ConsolesUpdateThenPoll(ctx context.Context, id ConsoleId, input ConsolePatchParameters, options ConsolesUpdateOperationOptions) error {
	result, err := c.ConsolesUpdate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing ConsolesUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ConsolesUpdate: %+v", err)
	}

	return nil
}
