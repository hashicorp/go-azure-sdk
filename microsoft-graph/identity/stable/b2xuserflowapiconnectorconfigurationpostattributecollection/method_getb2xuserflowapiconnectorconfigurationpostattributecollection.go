package b2xuserflowapiconnectorconfigurationpostattributecollection

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.IdentityApiConnector
}

type GetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions() GetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions {
	return GetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions{}
}

func (o GetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetB2xUserFlowApiConnectorConfigurationPostAttributeCollection - Get postAttributeCollection from identity
func (c B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient) GetB2xUserFlowApiConnectorConfigurationPostAttributeCollection(ctx context.Context, id stable.IdentityB2xUserFlowId, options GetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) (result GetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration/postAttributeCollection", id.ID()),
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

	var model stable.IdentityApiConnector
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
