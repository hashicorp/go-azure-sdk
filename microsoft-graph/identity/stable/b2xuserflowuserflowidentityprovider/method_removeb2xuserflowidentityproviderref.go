package b2xuserflowuserflowidentityprovider

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

type RemoveB2xUserFlowIdentityProviderRefOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveB2xUserFlowIdentityProviderRefOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultRemoveB2xUserFlowIdentityProviderRefOperationOptions() RemoveB2xUserFlowIdentityProviderRefOperationOptions {
	return RemoveB2xUserFlowIdentityProviderRefOperationOptions{}
}

func (o RemoveB2xUserFlowIdentityProviderRefOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveB2xUserFlowIdentityProviderRefOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveB2xUserFlowIdentityProviderRefOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveB2xUserFlowIdentityProviderRef - Delete ref of navigation property userFlowIdentityProviders for identity
func (c B2xUserFlowUserFlowIdentityProviderClient) RemoveB2xUserFlowIdentityProviderRef(ctx context.Context, id stable.IdentityB2xUserFlowIdUserFlowIdentityProviderId, options RemoveB2xUserFlowIdentityProviderRefOperationOptions) (result RemoveB2xUserFlowIdentityProviderRefOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$ref", id.ID()),
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
