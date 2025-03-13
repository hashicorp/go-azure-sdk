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

type ConsolesCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *Console
}

type ConsolesCreateOrUpdateOperationOptions struct {
	IfMatch     *string
	IfNoneMatch *string
}

func DefaultConsolesCreateOrUpdateOperationOptions() ConsolesCreateOrUpdateOperationOptions {
	return ConsolesCreateOrUpdateOperationOptions{}
}

func (o ConsolesCreateOrUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o ConsolesCreateOrUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ConsolesCreateOrUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ConsolesCreateOrUpdate ...
func (c NetworkcloudsClient) ConsolesCreateOrUpdate(ctx context.Context, id ConsoleId, input Console, options ConsolesCreateOrUpdateOperationOptions) (result ConsolesCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
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

// ConsolesCreateOrUpdateThenPoll performs ConsolesCreateOrUpdate then polls until it's completed
func (c NetworkcloudsClient) ConsolesCreateOrUpdateThenPoll(ctx context.Context, id ConsoleId, input Console, options ConsolesCreateOrUpdateOperationOptions) error {
	result, err := c.ConsolesCreateOrUpdate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing ConsolesCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ConsolesCreateOrUpdate: %+v", err)
	}

	return nil
}
