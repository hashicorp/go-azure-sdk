package domainjoinconnector

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetDomainJoinConnectorOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.DeviceManagementDomainJoinConnector
}

type GetDomainJoinConnectorOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetDomainJoinConnectorOperationOptions() GetDomainJoinConnectorOperationOptions {
	return GetDomainJoinConnectorOperationOptions{}
}

func (o GetDomainJoinConnectorOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetDomainJoinConnectorOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetDomainJoinConnectorOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetDomainJoinConnector - Get domainJoinConnectors from deviceManagement. A list of connector objects.
func (c DomainJoinConnectorClient) GetDomainJoinConnector(ctx context.Context, id beta.DeviceManagementDomainJoinConnectorId, options GetDomainJoinConnectorOperationOptions) (result GetDomainJoinConnectorOperationResponse, err error) {
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

	var model beta.DeviceManagementDomainJoinConnector
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
