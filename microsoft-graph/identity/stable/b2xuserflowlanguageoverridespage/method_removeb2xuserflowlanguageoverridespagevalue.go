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

type RemoveB2xUserFlowLanguageOverridesPageValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type RemoveB2xUserFlowLanguageOverridesPageValueOperationOptions struct {
	IfMatch  *string
	Metadata *odata.Metadata
}

func DefaultRemoveB2xUserFlowLanguageOverridesPageValueOperationOptions() RemoveB2xUserFlowLanguageOverridesPageValueOperationOptions {
	return RemoveB2xUserFlowLanguageOverridesPageValueOperationOptions{}
}

func (o RemoveB2xUserFlowLanguageOverridesPageValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o RemoveB2xUserFlowLanguageOverridesPageValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o RemoveB2xUserFlowLanguageOverridesPageValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// RemoveB2xUserFlowLanguageOverridesPageValue - Delete userFlowLanguagePage. Deletes the values in an
// userFlowLanguagePage object. You may only delete the values in an overridesPage, which is used to customize the
// values shown to a user during a user journey defined by a user flow.
func (c B2xUserFlowLanguageOverridesPageClient) RemoveB2xUserFlowLanguageOverridesPageValue(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageIdOverridesPageId, options RemoveB2xUserFlowLanguageOverridesPageValueOperationOptions) (result RemoveB2xUserFlowLanguageOverridesPageValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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
