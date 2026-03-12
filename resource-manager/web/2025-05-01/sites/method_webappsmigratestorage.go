package sites

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

type WebAppsMigrateStorageOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *StorageMigrationResponse
}

type WebAppsMigrateStorageOperationOptions struct {
	SubscriptionName *string
}

func DefaultWebAppsMigrateStorageOperationOptions() WebAppsMigrateStorageOperationOptions {
	return WebAppsMigrateStorageOperationOptions{}
}

func (o WebAppsMigrateStorageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o WebAppsMigrateStorageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o WebAppsMigrateStorageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}
	if o.SubscriptionName != nil {
		out.Append("subscriptionName", fmt.Sprintf("%v", *o.SubscriptionName))
	}
	return &out
}

// WebAppsMigrateStorage ...
func (c SitesClient) WebAppsMigrateStorage(ctx context.Context, id commonids.AppServiceId, input StorageMigrationOptions, options WebAppsMigrateStorageOperationOptions) (result WebAppsMigrateStorageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/migrate", id.ID()),
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

// WebAppsMigrateStorageThenPoll performs WebAppsMigrateStorage then polls until it's completed
func (c SitesClient) WebAppsMigrateStorageThenPoll(ctx context.Context, id commonids.AppServiceId, input StorageMigrationOptions, options WebAppsMigrateStorageOperationOptions) error {
	result, err := c.WebAppsMigrateStorage(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing WebAppsMigrateStorage: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after WebAppsMigrateStorage: %+v", err)
	}

	return nil
}
