package databasemigrations

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

type MongoToCosmosDbRUMongoDeleteOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
}

type MongoToCosmosDbRUMongoDeleteOperationOptions struct {
	Force *bool
}

func DefaultMongoToCosmosDbRUMongoDeleteOperationOptions() MongoToCosmosDbRUMongoDeleteOperationOptions {
	return MongoToCosmosDbRUMongoDeleteOperationOptions{}
}

func (o MongoToCosmosDbRUMongoDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o MongoToCosmosDbRUMongoDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o MongoToCosmosDbRUMongoDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.Force != nil {
		out.Append("force", fmt.Sprintf("%v", *o.Force))
	}
	return &out
}

// MongoToCosmosDbRUMongoDelete ...
func (c DatabaseMigrationsClient) MongoToCosmosDbRUMongoDelete(ctx context.Context, id DatabaseAccountProviders2DatabaseMigrationId, options MongoToCosmosDbRUMongoDeleteOperationOptions) (result MongoToCosmosDbRUMongoDeleteOperationResponse, err error) {
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

// MongoToCosmosDbRUMongoDeleteThenPoll performs MongoToCosmosDbRUMongoDelete then polls until it's completed
func (c DatabaseMigrationsClient) MongoToCosmosDbRUMongoDeleteThenPoll(ctx context.Context, id DatabaseAccountProviders2DatabaseMigrationId, options MongoToCosmosDbRUMongoDeleteOperationOptions) error {
	result, err := c.MongoToCosmosDbRUMongoDelete(ctx, id, options)
	if err != nil {
		return fmt.Errorf("performing MongoToCosmosDbRUMongoDelete: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after MongoToCosmosDbRUMongoDelete: %+v", err)
	}

	return nil
}
