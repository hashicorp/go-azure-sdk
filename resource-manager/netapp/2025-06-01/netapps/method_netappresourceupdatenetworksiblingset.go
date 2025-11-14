package netapps

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

type NetAppResourceUpdateNetworkSiblingSetOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *NetworkSiblingSet
}

// NetAppResourceUpdateNetworkSiblingSet ...
func (c NetappsClient) NetAppResourceUpdateNetworkSiblingSet(ctx context.Context, id LocationId, input UpdateNetworkSiblingSetRequest) (result NetAppResourceUpdateNetworkSiblingSetOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/updateNetworkSiblingSet", id.ID()),
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

// NetAppResourceUpdateNetworkSiblingSetThenPoll performs NetAppResourceUpdateNetworkSiblingSet then polls until it's completed
func (c NetappsClient) NetAppResourceUpdateNetworkSiblingSetThenPoll(ctx context.Context, id LocationId, input UpdateNetworkSiblingSetRequest) error {
	result, err := c.NetAppResourceUpdateNetworkSiblingSet(ctx, id, input)
	if err != nil {
		return fmt.Errorf("performing NetAppResourceUpdateNetworkSiblingSet: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after NetAppResourceUpdateNetworkSiblingSet: %+v", err)
	}

	return nil
}
