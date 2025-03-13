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

type L3NetworksCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *L3Network
}

type L3NetworksCreateOrUpdateOperationOptions struct {
	IfMatch     *string
	IfNoneMatch *string
}

func DefaultL3NetworksCreateOrUpdateOperationOptions() L3NetworksCreateOrUpdateOperationOptions {
	return L3NetworksCreateOrUpdateOperationOptions{}
}

func (o L3NetworksCreateOrUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o L3NetworksCreateOrUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o L3NetworksCreateOrUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// L3NetworksCreateOrUpdate ...
func (c NetworkcloudsClient) L3NetworksCreateOrUpdate(ctx context.Context, id L3NetworkId, input L3Network, options L3NetworksCreateOrUpdateOperationOptions) (result L3NetworksCreateOrUpdateOperationResponse, err error) {
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

// L3NetworksCreateOrUpdateThenPoll performs L3NetworksCreateOrUpdate then polls until it's completed
func (c NetworkcloudsClient) L3NetworksCreateOrUpdateThenPoll(ctx context.Context, id L3NetworkId, input L3Network, options L3NetworksCreateOrUpdateOperationOptions) error {
	result, err := c.L3NetworksCreateOrUpdate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing L3NetworksCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after L3NetworksCreateOrUpdate: %+v", err)
	}

	return nil
}
