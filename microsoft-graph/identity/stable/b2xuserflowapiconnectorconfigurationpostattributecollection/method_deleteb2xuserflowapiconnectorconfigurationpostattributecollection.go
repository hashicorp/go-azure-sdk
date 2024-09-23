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

type DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions() DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions {
	return DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions{}
}

func (o DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollection - Delete navigation property
// postAttributeCollection for identity
func (c B2xUserFlowApiConnectorConfigurationPostAttributeCollectionClient) DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollection(ctx context.Context, id stable.IdentityB2xUserFlowId, options DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationOptions) (result DeleteB2xUserFlowApiConnectorConfigurationPostAttributeCollectionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/apiConnectorConfiguration/postAttributeCollection", id.ID()),
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

	return
}
