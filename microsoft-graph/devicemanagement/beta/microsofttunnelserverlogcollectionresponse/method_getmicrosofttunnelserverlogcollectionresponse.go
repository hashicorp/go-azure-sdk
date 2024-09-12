package microsofttunnelserverlogcollectionresponse

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMicrosoftTunnelServerLogCollectionResponseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MicrosoftTunnelServerLogCollectionResponse
}

type GetMicrosoftTunnelServerLogCollectionResponseOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetMicrosoftTunnelServerLogCollectionResponseOperationOptions() GetMicrosoftTunnelServerLogCollectionResponseOperationOptions {
	return GetMicrosoftTunnelServerLogCollectionResponseOperationOptions{}
}

func (o GetMicrosoftTunnelServerLogCollectionResponseOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMicrosoftTunnelServerLogCollectionResponseOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetMicrosoftTunnelServerLogCollectionResponseOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetMicrosoftTunnelServerLogCollectionResponse - Get microsoftTunnelServerLogCollectionResponses from
// deviceManagement. Collection of MicrosoftTunnelServerLogCollectionResponse settings associated with account.
func (c MicrosoftTunnelServerLogCollectionResponseClient) GetMicrosoftTunnelServerLogCollectionResponse(ctx context.Context, id beta.DeviceManagementMicrosoftTunnelServerLogCollectionResponseId, options GetMicrosoftTunnelServerLogCollectionResponseOperationOptions) (result GetMicrosoftTunnelServerLogCollectionResponseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model beta.MicrosoftTunnelServerLogCollectionResponse
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
