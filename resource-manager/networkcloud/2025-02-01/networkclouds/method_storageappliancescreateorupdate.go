package networkclouds

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

type StorageAppliancesCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *StorageAppliance
}

type StorageAppliancesCreateOrUpdateOperationOptions struct {
	IfMatch     *string
	IfNoneMatch *string
}

func DefaultStorageAppliancesCreateOrUpdateOperationOptions() StorageAppliancesCreateOrUpdateOperationOptions {
	return StorageAppliancesCreateOrUpdateOperationOptions{}
}

func (o StorageAppliancesCreateOrUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	if o.IfNoneMatch != nil {
		out.Append("If-None-Match", fmt.Sprintf("%v", *o.IfNoneMatch))
	}
	return &out
}

func (o StorageAppliancesCreateOrUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o StorageAppliancesCreateOrUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// StorageAppliancesCreateOrUpdate ...
func (c NetworkcloudsClient) StorageAppliancesCreateOrUpdate(ctx context.Context, id StorageApplianceId, input StorageAppliance, options StorageAppliancesCreateOrUpdateOperationOptions) (result StorageAppliancesCreateOrUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          id.ID(),
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

// StorageAppliancesCreateOrUpdateThenPoll performs StorageAppliancesCreateOrUpdate then polls until it's completed
func (c NetworkcloudsClient) StorageAppliancesCreateOrUpdateThenPoll(ctx context.Context, id StorageApplianceId, input StorageAppliance, options StorageAppliancesCreateOrUpdateOperationOptions) error {
	result, err := c.StorageAppliancesCreateOrUpdate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing StorageAppliancesCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after StorageAppliancesCreateOrUpdate: %+v", err)
	}

	return nil
}
