package b2xuserflowlanguageoverridespage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateB2xUserFlowLanguageOverridesPageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateB2xUserFlowLanguageOverridesPageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateB2xUserFlowLanguageOverridesPageOperationOptions() UpdateB2xUserFlowLanguageOverridesPageOperationOptions {
	return UpdateB2xUserFlowLanguageOverridesPageOperationOptions{}
}

func (o UpdateB2xUserFlowLanguageOverridesPageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateB2xUserFlowLanguageOverridesPageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateB2xUserFlowLanguageOverridesPageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateB2xUserFlowLanguageOverridesPage - Update userFlowLanguagePage. Update the values in an userFlowLanguagePage
// object. You may only update the values in an overridesPage, which is used to customize the values shown to a user
// during a user journey defined by a user flow.
func (c B2xUserFlowLanguageOverridesPageClient) UpdateB2xUserFlowLanguageOverridesPage(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageIdOverridesPageId, input stable.UserFlowLanguagePage, options UpdateB2xUserFlowLanguageOverridesPageOperationOptions) (result UpdateB2xUserFlowLanguageOverridesPageOperationResponse, err error) {
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
