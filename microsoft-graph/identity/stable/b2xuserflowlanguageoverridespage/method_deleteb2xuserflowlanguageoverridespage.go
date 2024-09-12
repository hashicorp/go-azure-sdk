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

type DeleteB2xUserFlowLanguageOverridesPageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteB2xUserFlowLanguageOverridesPageOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteB2xUserFlowLanguageOverridesPageOperationOptions() DeleteB2xUserFlowLanguageOverridesPageOperationOptions {
	return DeleteB2xUserFlowLanguageOverridesPageOperationOptions{}
}

func (o DeleteB2xUserFlowLanguageOverridesPageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteB2xUserFlowLanguageOverridesPageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteB2xUserFlowLanguageOverridesPageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteB2xUserFlowLanguageOverridesPage - Delete userFlowLanguagePage. Deletes the values in an userFlowLanguagePage
// object. You may only delete the values in an overridesPage, which is used to customize the values shown to a user
// during a user journey defined by a user flow.
func (c B2xUserFlowLanguageOverridesPageClient) DeleteB2xUserFlowLanguageOverridesPage(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageIdOverridesPageId, options DeleteB2xUserFlowLanguageOverridesPageOperationOptions) (result DeleteB2xUserFlowLanguageOverridesPageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          id.ID(),
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
