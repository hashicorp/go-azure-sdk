package b2xuserflowlanguagedefaultpage

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetB2xUserFlowLanguageDefaultPageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.UserFlowLanguagePage
}

type GetB2xUserFlowLanguageDefaultPageOperationOptions struct {
	Expand   *odata.Expand
	Metadata *odata.Metadata
	Select   *[]string
}

func DefaultGetB2xUserFlowLanguageDefaultPageOperationOptions() GetB2xUserFlowLanguageDefaultPageOperationOptions {
	return GetB2xUserFlowLanguageDefaultPageOperationOptions{}
}

func (o GetB2xUserFlowLanguageDefaultPageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetB2xUserFlowLanguageDefaultPageOperationOptions) ToOData() *odata.Query {
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

func (o GetB2xUserFlowLanguageDefaultPageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetB2xUserFlowLanguageDefaultPage - Get defaultPages from identity. Collection of pages with the default content to
// display in a user flow for a specified language. This collection doesn't allow any kind of modification.
func (c B2xUserFlowLanguageDefaultPageClient) GetB2xUserFlowLanguageDefaultPage(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageIdDefaultPageId, options GetB2xUserFlowLanguageDefaultPageOperationOptions) (result GetB2xUserFlowLanguageDefaultPageOperationResponse, err error) {
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

	var model stable.UserFlowLanguagePage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
