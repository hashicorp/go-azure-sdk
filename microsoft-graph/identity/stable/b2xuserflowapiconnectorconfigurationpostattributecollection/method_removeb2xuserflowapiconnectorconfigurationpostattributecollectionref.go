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

type RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions struct {
	IfMatch *string
}

func DefaultRemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions() RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions {
	return RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions{}
}

func (o RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRef - Delete ref of navigation property
// postAttributeCollection for identity
func (c B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient) RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRef(ctx context.Context, id stable.IdentityB2xUserFlowId, options RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationOptions) (result RemoveB2xUserFlowApiConnectorConfigurationPostAttributeCollectionRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration/postAttributeCollection/$ref", id.ID()),
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

	return
}
