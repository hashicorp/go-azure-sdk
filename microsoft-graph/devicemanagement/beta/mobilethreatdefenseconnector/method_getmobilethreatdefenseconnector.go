package mobilethreatdefenseconnector

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetMobileThreatDefenseConnectorOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.MobileThreatDefenseConnector
}

type GetMobileThreatDefenseConnectorOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetMobileThreatDefenseConnectorOperationOptions() GetMobileThreatDefenseConnectorOperationOptions {
	return GetMobileThreatDefenseConnectorOperationOptions{}
}

func (o GetMobileThreatDefenseConnectorOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetMobileThreatDefenseConnectorOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetMobileThreatDefenseConnectorOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetMobileThreatDefenseConnector - Get mobileThreatDefenseConnectors from deviceManagement. The list of Mobile threat
// Defense connectors configured by the tenant.
func (c MobileThreatDefenseConnectorClient) GetMobileThreatDefenseConnector(ctx context.Context, id beta.DeviceManagementMobileThreatDefenseConnectorId, options GetMobileThreatDefenseConnectorOperationOptions) (result GetMobileThreatDefenseConnectorOperationResponse, err error) {
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

	var model beta.MobileThreatDefenseConnector
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
