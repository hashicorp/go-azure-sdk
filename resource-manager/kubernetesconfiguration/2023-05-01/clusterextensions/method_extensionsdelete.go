package clusterextensions

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

type ExtensionsDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type ExtensionsDeleteOperationOptions struct {
	ForceDelete *bool
}

func DefaultExtensionsDeleteOperationOptions() ExtensionsDeleteOperationOptions {
	return ExtensionsDeleteOperationOptions{}
}

func (o ExtensionsDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ExtensionsDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o ExtensionsDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ForceDelete != nil {
		out.Append("forceDelete", fmt.Sprintf("%v", *o.ForceDelete))
	}
	return &out
}

// ExtensionsDelete ...
func (c ClusterExtensionsClient) ExtensionsDelete(ctx context.Context, id ScopedExtensionId, options ExtensionsDeleteOperationOptions) (result ExtensionsDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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

	result.Poller, err = resourcemanager.PollerFromResponse(resp, c.Client)
	if err != nil {
		return
	}

	return
}

// ExtensionsDeleteThenPoll performs ExtensionsDelete then polls until it's completed
func (c ClusterExtensionsClient) ExtensionsDeleteThenPoll(ctx context.Context, id ScopedExtensionId, options ExtensionsDeleteOperationOptions) error {
	result, err := c.ExtensionsDelete(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing ExtensionsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after ExtensionsDelete: %+v", err)
	}

	return nil
}
