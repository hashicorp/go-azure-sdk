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

type GetB2xUserFlowLanguageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserFlowLanguageConfiguration
}

type GetB2xUserFlowLanguageOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetB2xUserFlowLanguageOperationOptions() GetB2xUserFlowLanguageOperationOptions {
	return GetB2xUserFlowLanguageOperationOptions{}
}

func (o GetB2xUserFlowLanguageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetB2xUserFlowLanguageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetB2xUserFlowLanguageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetB2xUserFlowLanguage - Get userFlowLanguageConfiguration. Read the properties and relationships of a
// userFlowLanguageConfiguration object. These objects represent a language available in a user flow. Note: Language
// customization is enabled by default in Microsoft Entra user flows.
func (c B2xUserFlowLanguageClient) GetB2xUserFlowLanguage(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageId, options GetB2xUserFlowLanguageOperationOptions) (result GetB2xUserFlowLanguageOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
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

	var model stable.UserFlowLanguageConfiguration
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
