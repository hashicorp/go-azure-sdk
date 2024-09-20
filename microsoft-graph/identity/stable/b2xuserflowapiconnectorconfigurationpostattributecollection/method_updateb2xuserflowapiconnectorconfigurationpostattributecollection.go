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

type UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions() UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions {
	return UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions{}
}

func (o UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollection - Update the navigation property
// postAttributeCollection in identity
func (c B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient) UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollection(ctx context.Context, id stable.IdentityB2xUserFlowId, input stable.IdentityApiConnector, options UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) (result UpdateB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration/postAttributeCollection", id.ID()),
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

	return
}
