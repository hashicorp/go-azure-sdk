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

// CreateZebraFotaConnectorDisconnect - Invoke action disconnect
func (c ZebraFotaConnectorClient) CreateZebraFotaConnectorDisconnect(ctx context.Context) (result CreateZebraFotaConnectorDisconnectOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       "/deviceManagement/zebraFotaConnector/disconnect",
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
