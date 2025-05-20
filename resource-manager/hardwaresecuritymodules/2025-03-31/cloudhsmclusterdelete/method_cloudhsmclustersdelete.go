package cloudhsmclusterdelete

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

type CloudHsmClustersDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

// CloudHsmClustersDelete ...
func (c CloudHSMClusterDeleteClient) CloudHsmClustersDelete(ctx context.Context, id CloudHsmClusterId) (result CloudHsmClustersDeleteOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
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

// CloudHsmClustersDeleteThenPoll performs CloudHsmClustersDelete then polls until it's completed
func (c CloudHSMClusterDeleteClient) CloudHsmClustersDeleteThenPoll(ctx context.Context, id CloudHsmClusterId) error {
	result, err := c.CloudHsmClustersDelete(ctx, id)
	if err != nil {
		return fmt.Errorf("performing CloudHsmClustersDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after CloudHsmClustersDelete: %+v", err)
	}

	return nil
}
