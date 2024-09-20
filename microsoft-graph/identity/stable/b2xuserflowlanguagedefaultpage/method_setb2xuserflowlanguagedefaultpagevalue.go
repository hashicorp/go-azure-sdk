package b2xuserflowlanguagedefaultpage

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

type SetB2xUserFlowLanguageDefaultPageValueOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SetB2xUserFlowLanguageDefaultPageValueOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSetB2xUserFlowLanguageDefaultPageValueOperationOptions() SetB2xUserFlowLanguageDefaultPageValueOperationOptions {
	return SetB2xUserFlowLanguageDefaultPageValueOperationOptions{}
}

func (o SetB2xUserFlowLanguageDefaultPageValueOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SetB2xUserFlowLanguageDefaultPageValueOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SetB2xUserFlowLanguageDefaultPageValueOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SetB2xUserFlowLanguageDefaultPageValue - Update media content for the navigation property defaultPages in identity.
// The unique identifier for an entity. Read-only.
func (c B2xUserFlowLanguageDefaultPageClient) SetB2xUserFlowLanguageDefaultPageValue(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageIdDefaultPageId, input []byte, options SetB2xUserFlowLanguageDefaultPageValueOperationOptions) (result SetB2xUserFlowLanguageDefaultPageValueOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPut,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/$value", id.ID()),
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
