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

type VaultsUpdateOperationResponse struct {
	Poller       pollers.Poller
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *Vault
}

type VaultsUpdateOperationOptions struct {
	XMsAuthorizationAuxiliary *string
}

func DefaultVaultsUpdateOperationOptions() VaultsUpdateOperationOptions {
	return VaultsUpdateOperationOptions{}
}

func (o VaultsUpdateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.XMsAuthorizationAuxiliary != nil {
		out.Append("x-ms-authorization-auxiliary", fmt.Sprintf("%v", *o.XMsAuthorizationAuxiliary))
	}
	return &out
}

func (o VaultsUpdateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o VaultsUpdateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// VaultsUpdate ...
func (c OpenapisClient) VaultsUpdate(ctx context.Context, id VaultId, input PatchVault, options VaultsUpdateOperationOptions) (result VaultsUpdateOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
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

// VaultsUpdateThenPoll performs VaultsUpdate then polls until it's completed
func (c OpenapisClient) VaultsUpdateThenPoll(ctx context.Context, id VaultId, input PatchVault, options VaultsUpdateOperationOptions) error {
	result, err := c.VaultsUpdate(ctx, id, input, options)
	if err != nil {
		return fmt.Errorf("performing VaultsUpdate: %+v", err)
	}

	if err := result.Poller.PollUntilDone(ctx); err != nil {
		return fmt.Errorf("polling after VaultsUpdate: %+v", err)
	}

	return nil
}
