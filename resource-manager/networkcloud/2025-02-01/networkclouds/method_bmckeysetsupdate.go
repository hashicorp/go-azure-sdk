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

type BmcKeySetsUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *BmcKeySet
}

type BmcKeySetsUpdateOperationOptions struct {
	IfMatch     *string
	IfNoneMatch *string
}

func DefaultBmcKeySetsUpdateOperationOptions() BmcKeySetsUpdateOperationOptions {
	return BmcKeySetsUpdateOperationOptions{}
}

func (o BmcKeySetsUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o BmcKeySetsUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o BmcKeySetsUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// BmcKeySetsUpdate ...
func (c NetworkcloudsClient) BmcKeySetsUpdate(ctx context.Context, id BmcKeySetId, input BmcKeySetPatchParameters, options BmcKeySetsUpdateOperationOptions) (result BmcKeySetsUpdateOperationResponse, err error) {
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

// BmcKeySetsUpdateThenPoll performs BmcKeySetsUpdate then polls until it's completed
func (c NetworkcloudsClient) BmcKeySetsUpdateThenPoll(ctx context.Context, id BmcKeySetId, input BmcKeySetPatchParameters, options BmcKeySetsUpdateOperationOptions) error {
	result, err := c.BmcKeySetsUpdate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing BmcKeySetsUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after BmcKeySetsUpdate: %+v", err)
	}

	return nil
}
