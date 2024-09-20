package zebrafotaconnector

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateZebraFotaConnectorApproveFotaAppOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CreateZebraFotaConnectorApproveFotaAppResult
}

type CreateZebraFotaConnectorApproveFotaAppOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateZebraFotaConnectorApproveFotaAppOperationOptions() CreateZebraFotaConnectorApproveFotaAppOperationOptions {
	return CreateZebraFotaConnectorApproveFotaAppOperationOptions{}
}

func (o CreateZebraFotaConnectorApproveFotaAppOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateZebraFotaConnectorApproveFotaAppOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateZebraFotaConnectorApproveFotaAppOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateZebraFotaConnectorApproveFotaApp - Invoke action approveFotaApps
func (c ZebraFotaConnectorClient) CreateZebraFotaConnectorApproveFotaApp(ctx context.Context, options CreateZebraFotaConnectorApproveFotaAppOperationOptions) (result CreateZebraFotaConnectorApproveFotaAppOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          "/deviceManagement/zebraFotaConnector/approveFotaApps",
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

	var model CreateZebraFotaConnectorApproveFotaAppResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
