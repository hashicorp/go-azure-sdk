package appserviceenvironmentresources

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/commonids"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppServiceEnvironmentsDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type AppServiceEnvironmentsDeleteOperationOptions struct {
	ForceDelete *bool
}

func DefaultAppServiceEnvironmentsDeleteOperationOptions() AppServiceEnvironmentsDeleteOperationOptions {
	return AppServiceEnvironmentsDeleteOperationOptions{}
}

func (o AppServiceEnvironmentsDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o AppServiceEnvironmentsDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o AppServiceEnvironmentsDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ForceDelete != nil {
		out.Append("forceDelete", fmt.Sprintf("%v", *o.ForceDelete))
	}
	return &out
}

// AppServiceEnvironmentsDelete ...
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsDelete(ctx context.Context, id commonids.AppServiceEnvironmentId, options AppServiceEnvironmentsDeleteOperationOptions) (result AppServiceEnvironmentsDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
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

// AppServiceEnvironmentsDeleteThenPoll performs AppServiceEnvironmentsDelete then polls until it's completed
func (c AppServiceEnvironmentResourcesClient) AppServiceEnvironmentsDeleteThenPoll(ctx context.Context, id commonids.AppServiceEnvironmentId, options AppServiceEnvironmentsDeleteOperationOptions) error {
	result, err := c.AppServiceEnvironmentsDelete(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing AppServiceEnvironmentsDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after AppServiceEnvironmentsDelete: %+v", err)
	}

	return nil
}
