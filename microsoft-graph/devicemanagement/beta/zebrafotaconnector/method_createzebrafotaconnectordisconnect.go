package zebrafotaconnector

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateZebraFotaConnectorDisconnectOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CreateZebraFotaConnectorDisconnectResult
}

type CreateZebraFotaConnectorDisconnectOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateZebraFotaConnectorDisconnectOperationOptions() CreateZebraFotaConnectorDisconnectOperationOptions {
	return CreateZebraFotaConnectorDisconnectOperationOptions{}
}

func (o CreateZebraFotaConnectorDisconnectOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateZebraFotaConnectorDisconnectOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateZebraFotaConnectorDisconnectOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateZebraFotaConnectorDisconnect - Invoke action disconnect
func (c ZebraFotaConnectorClient) CreateZebraFotaConnectorDisconnect(ctx context.Context, options CreateZebraFotaConnectorDisconnectOperationOptions) (result CreateZebraFotaConnectorDisconnectOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/zebraFotaConnector/disconnect",
		RetryFunc:     options.RetryFunc,
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

	var model CreateZebraFotaConnectorDisconnectResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
