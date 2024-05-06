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

type DeleteAtManagementGroupOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAtManagementGroupOperationOptions struct {
	BypassStackOutOfSyncError      *bool
	UnmanageActionManagementGroups *UnmanageActionManagementGroupMode
	UnmanageActionResourceGroups   *UnmanageActionResourceGroupMode
	UnmanageActionResources        *UnmanageActionResourceMode
}

func DefaultDeleteAtManagementGroupOperationOptions() DeleteAtManagementGroupOperationOptions {
	return DeleteAtManagementGroupOperationOptions{}
}

func (o DeleteAtManagementGroupOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DeleteAtManagementGroupOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o DeleteAtManagementGroupOperationOptions) ToQuery() *client.QueryParams {
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

// DeleteAtManagementGroup ...
func (c DeploymentStacksClient) DeleteAtManagementGroup(ctx context.Context, id Providers2DeploymentStackId, options DeleteAtManagementGroupOperationOptions) (result DeleteAtManagementGroupOperationResponse, err error) {
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

// DeleteAtManagementGroupThenPoll performs DeleteAtManagementGroup then polls until it's completed
func (c DeploymentStacksClient) DeleteAtManagementGroupThenPoll(ctx context.Context, id Providers2DeploymentStackId, options DeleteAtManagementGroupOperationOptions) error {
	result, err := c.DeleteAtManagementGroup(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing DeleteAtManagementGroup: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DeleteAtManagementGroup: %+v", err)
	}

	return nil
}
