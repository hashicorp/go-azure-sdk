package openapis

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

type VaultsCreateOrUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *Vault
}

type VaultsCreateOrUpdateOperationOptions struct {
	XMsAuthorizationAuxiliary *string
}

func DefaultVaultsCreateOrUpdateOperationOptions() VaultsCreateOrUpdateOperationOptions {
	return VaultsCreateOrUpdateOperationOptions{}
}

func (o VaultsCreateOrUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.XMsAuthorizationAuxiliary != nil {
		out.Append("x-ms-authorization-auxiliary", fmt.Sprintf("%v", *o.XMsAuthorizationAuxiliary))
	}
	return &out
}

func (o VaultsCreateOrUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o VaultsCreateOrUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// VaultsCreateOrUpdate ...
func (c OpenapisClient) VaultsCreateOrUpdate(ctx context.Context, id VaultId, input Vault, options VaultsCreateOrUpdateOperationOptions) (result VaultsCreateOrUpdateOperationResponse, err error) {
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

// VaultsCreateOrUpdateThenPoll performs VaultsCreateOrUpdate then polls until it's completed
func (c OpenapisClient) VaultsCreateOrUpdateThenPoll(ctx context.Context, id VaultId, input Vault, options VaultsCreateOrUpdateOperationOptions) error {
	result, err := c.VaultsCreateOrUpdate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing VaultsCreateOrUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after VaultsCreateOrUpdate: %+v", err)
	}

	return nil
}
