package b2xuserflowlanguageoverridespage

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

type SetB2xUserFlowLanguageOverridesPageValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetB2xUserFlowLanguageOverridesPageValueOperationOptions struct {
	ContentType string
	Metadata    *odata.Metadata
	RetryFunc   client.RequestRetryFunc
}

func DefaultSetB2xUserFlowLanguageOverridesPageValueOperationOptions() SetB2xUserFlowLanguageOverridesPageValueOperationOptions {
	return SetB2xUserFlowLanguageOverridesPageValueOperationOptions{}
}

func (o SetB2xUserFlowLanguageOverridesPageValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetB2xUserFlowLanguageOverridesPageValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetB2xUserFlowLanguageOverridesPageValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetB2xUserFlowLanguageOverridesPageValue - Update userFlowLanguagePage. Update the values in an userFlowLanguagePage
// object. You may only update the values in an overridesPage, which is used to customize the values shown to a user
// during a user journey defined by a user flow.
func (c B2xUserFlowLanguageOverridesPageClient) SetB2xUserFlowLanguageOverridesPageValue(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageIdOverridesPageId, input []byte, options SetB2xUserFlowLanguageOverridesPageValueOperationOptions) (result SetB2xUserFlowLanguageOverridesPageValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: options.ContentType,
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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
