package b2xuserflowlanguage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateB2xUserFlowLanguageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateB2xUserFlowLanguageOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateB2xUserFlowLanguageOperationOptions() UpdateB2xUserFlowLanguageOperationOptions {
	return UpdateB2xUserFlowLanguageOperationOptions{}
}

func (o UpdateB2xUserFlowLanguageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateB2xUserFlowLanguageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateB2xUserFlowLanguageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateB2xUserFlowLanguage - Update the navigation property languages in identity
func (c B2xUserFlowLanguageClient) UpdateB2xUserFlowLanguage(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageId, input stable.UserFlowLanguageConfiguration, options UpdateB2xUserFlowLanguageOperationOptions) (result UpdateB2xUserFlowLanguageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
