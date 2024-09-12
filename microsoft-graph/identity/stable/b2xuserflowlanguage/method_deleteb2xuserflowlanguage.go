package b2xuserflowlanguage

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

type DeleteB2xUserFlowLanguageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteB2xUserFlowLanguageOperationOptions struct {
	IfMatch *string
}

func DefaultDeleteB2xUserFlowLanguageOperationOptions() DeleteB2xUserFlowLanguageOperationOptions {
	return DeleteB2xUserFlowLanguageOperationOptions{}
}

func (o DeleteB2xUserFlowLanguageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteB2xUserFlowLanguageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}

	return &out
}

func (o DeleteB2xUserFlowLanguageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteB2xUserFlowLanguage - Delete navigation property languages for identity
func (c B2xUserFlowLanguageClient) DeleteB2xUserFlowLanguage(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageId, options DeleteB2xUserFlowLanguageOperationOptions) (result DeleteB2xUserFlowLanguageOperationResponse, err error) {
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
