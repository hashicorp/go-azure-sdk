package failoverdatabases

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

type DatabasesFailoverOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type DatabasesFailoverOperationOptions struct {
	ReplicaType *ReplicaType
}

func DefaultDatabasesFailoverOperationOptions() DatabasesFailoverOperationOptions {
	return DatabasesFailoverOperationOptions{}
}

func (o DatabasesFailoverOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o DatabasesFailoverOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	return &out
}

func (o DatabasesFailoverOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.ReplicaType != nil {
		out.Append("replicaType", fmt.Sprintf("%v", *o.ReplicaType))
	}
	return &out
}

// DatabasesFailover ...
func (c FailoverDatabasesClient) DatabasesFailover(ctx context.Context, id DatabaseId, options DatabasesFailoverOperationOptions) (result DatabasesFailoverOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		Path:          fmt.Sprintf("%s/failover", id.ID()),
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

// DatabasesFailoverThenPoll performs DatabasesFailover then polls until it's completed
func (c FailoverDatabasesClient) DatabasesFailoverThenPoll(ctx context.Context, id DatabaseId, options DatabasesFailoverOperationOptions) error {
	result, err := c.DatabasesFailover(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing DatabasesFailover: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after DatabasesFailover: %+v", err)
	}

	return nil
}
