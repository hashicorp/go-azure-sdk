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

type SetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions() SetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions {
	return SetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions{}
}

func (o SetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRef - Update the ref of navigation property
// postAttributeCollection in identity
func (c B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient) SetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRef(ctx context.Context, id stable.IdentityB2xUserFlowId, input stable.ReferenceUpdate, options SetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions) (result SetB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration/postAttributeCollection/$ref", id.ID()),
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
