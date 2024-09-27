package exchangeresourcenamespaceresourceaction

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateExchangeResourceNamespaceResourceActionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateExchangeResourceNamespaceResourceActionOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateExchangeResourceNamespaceResourceActionOperationOptions() UpdateExchangeResourceNamespaceResourceActionOperationOptions {
	return UpdateExchangeResourceNamespaceResourceActionOperationOptions{}
}

func (o UpdateExchangeResourceNamespaceResourceActionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateExchangeResourceNamespaceResourceActionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateExchangeResourceNamespaceResourceActionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateExchangeResourceNamespaceResourceAction - Update the navigation property resourceActions in roleManagement
func (c ExchangeResourceNamespaceResourceActionClient) UpdateExchangeResourceNamespaceResourceAction(ctx context.Context, id beta.RoleManagementExchangeResourceNamespaceIdResourceActionId, input beta.UnifiedRbacResourceAction, options UpdateExchangeResourceNamespaceResourceActionOperationOptions) (result UpdateExchangeResourceNamespaceResourceActionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
		RetryFunc:     options.RetryFunc,
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
