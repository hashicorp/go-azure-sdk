package deploymentstacks

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

type DeleteAtResourceGroupOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAtResourceGroupOperationOptions struct {
	BypassStackOutOfSyncError      *bool
	UnmanageActionManagementGroups *UnmanageActionManagementGroupMode
	UnmanageActionResourceGroups   *UnmanageActionResourceGroupMode
	UnmanageActionResources        *UnmanageActionResourceMode
}

func DefaultDeleteAtResourceGroupOperationOptions() DeleteAtResourceGroupOperationOptions {
	return DeleteAtResourceGroupOperationOptions{}
}

func (o DeleteAtResourceGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeleteAtResourceGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o DeleteAtResourceGroupOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.BypassStackOutOfSyncError != nil {
		out.Append("bypassStackOutOfSyncError", fmt.Sprintf("%v", *o.BypassStackOutOfSyncError))
	}
	if o.UnmanageActionManagementGroups != nil {
		out.Append("unmanageAction.ManagementGroups", fmt.Sprintf("%v", *o.UnmanageActionManagementGroups))
	}
	if o.UnmanageActionResourceGroups != nil {
		out.Append("unmanageAction.ResourceGroups", fmt.Sprintf("%v", *o.UnmanageActionResourceGroups))
	}
	if o.UnmanageActionResources != nil {
		out.Append("unmanageAction.Resources", fmt.Sprintf("%v", *o.UnmanageActionResources))
	}
	return &out
}

// DeleteAtResourceGroup ...
func (c DeploymentStacksClient) DeleteAtResourceGroup(ctx context.Context, id ProviderDeploymentStackId, options DeleteAtResourceGroupOperationOptions) (result DeleteAtResourceGroupOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodDelete,
		Path:          id.ID(),
		OptionsObject: options,
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

// DeleteAtResourceGroupThenPoll performs DeleteAtResourceGroup then polls until it's completed
func (c DeploymentStacksClient) DeleteAtResourceGroupThenPoll(ctx context.Context, id ProviderDeploymentStackId, options DeleteAtResourceGroupOperationOptions) error {
	result, err := c.DeleteAtResourceGroup(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing DeleteAtResourceGroup: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DeleteAtResourceGroup: %+v", err)
	}

	return nil
}
