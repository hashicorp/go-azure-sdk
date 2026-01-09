package b2xuserflowlanguageoverridespage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetB2xUserFlowLanguageOverridesPageValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetB2xUserFlowLanguageOverridesPageValueOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultGetB2xUserFlowLanguageOverridesPageValueOperationOptions() GetB2xUserFlowLanguageOverridesPageValueOperationOptions {
	return GetB2xUserFlowLanguageOverridesPageValueOperationOptions{}
}

func (o GetB2xUserFlowLanguageOverridesPageValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetB2xUserFlowLanguageOverridesPageValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o GetB2xUserFlowLanguageOverridesPageValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetB2xUserFlowLanguageOverridesPageValue - List overridesPages. Get the userFlowLanguagePage resources from the
// overridesPages navigation property. These pages are used to customize the values shown to the user during a user
// journey in a user flow.
func (c B2xUserFlowLanguageOverridesPageClient) GetB2xUserFlowLanguageOverridesPageValue(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageIdOverridesPageId, options GetB2xUserFlowLanguageOverridesPageValueOperationOptions) (result GetB2xUserFlowLanguageOverridesPageValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/octet-stream",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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

	var model []byte
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
