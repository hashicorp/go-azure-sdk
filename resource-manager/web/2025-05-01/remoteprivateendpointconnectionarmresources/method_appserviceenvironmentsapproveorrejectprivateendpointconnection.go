package remoteprivateendpointconnectionarmresources

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

type AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnectionOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *RemotePrivateEndpointConnectionARMResource
}

// AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnection ...
func (c RemotePrivateEndpointConnectionARMResourcesClient) AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnection(ctx context.Context, id HostingEnvironmentPrivateEndpointConnectionId, input RemotePrivateEndpointConnectionARMResource) (result AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       id.ID(),
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

// AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnectionThenPoll performs AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnection then polls until it's completed
func (c RemotePrivateEndpointConnectionARMResourcesClient) AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnectionThenPoll(ctx context.Context, id HostingEnvironmentPrivateEndpointConnectionId, input RemotePrivateEndpointConnectionARMResource) error {
	result, err := c.AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnection(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnection: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after AppServiceEnvironmentsApproveOrRejectPrivateEndpointConnection: %+v", err)
	}

	return nil
}
