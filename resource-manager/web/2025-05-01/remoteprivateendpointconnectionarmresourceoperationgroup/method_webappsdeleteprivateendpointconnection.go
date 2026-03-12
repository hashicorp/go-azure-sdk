package remoteprivateendpointconnectionarmresourceoperationgroup

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

type WebAppsDeletePrivateEndpointConnectionOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *interface{}
}

// WebAppsDeletePrivateEndpointConnection ...
func (c RemotePrivateEndpointConnectionARMResourceOperationGroupClient) WebAppsDeletePrivateEndpointConnection(ctx context.Context, id PrivateEndpointConnectionId) (result WebAppsDeletePrivateEndpointConnectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod: http.MethodDelete,
		Path:       id.ID(),
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

// WebAppsDeletePrivateEndpointConnectionThenPoll performs WebAppsDeletePrivateEndpointConnection then polls until it's completed
func (c RemotePrivateEndpointConnectionARMResourceOperationGroupClient) WebAppsDeletePrivateEndpointConnectionThenPoll(ctx context.Context, id PrivateEndpointConnectionId) error {
	result, err := c.WebAppsDeletePrivateEndpointConnection(ctx, id)
	if err != nil {
		return fmt.Errorf("performing WebAppsDeletePrivateEndpointConnection: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after WebAppsDeletePrivateEndpointConnection: %+v", err)
	}

	return nil
}
