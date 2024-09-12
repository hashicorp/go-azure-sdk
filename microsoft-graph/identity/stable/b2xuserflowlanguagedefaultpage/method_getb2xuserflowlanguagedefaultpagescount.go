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

type GetB2xUserFlowLanguageDefaultPagesCountOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]byte
}

type GetB2xUserFlowLanguageDefaultPagesCountOperationOptions struct {
	Filter *string
	Search *string
}

func DefaultGetB2xUserFlowLanguageDefaultPagesCountOperationOptions() GetB2xUserFlowLanguageDefaultPagesCountOperationOptions {
	return GetB2xUserFlowLanguageDefaultPagesCountOperationOptions{}
}

func (o GetB2xUserFlowLanguageDefaultPagesCountOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetB2xUserFlowLanguageDefaultPagesCountOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Filter != nil {
		out.Filter = *o.Filter
	}
	if o.Search != nil {
		out.Search = *o.Search
	}
	return &out
}

func (o GetB2xUserFlowLanguageDefaultPagesCountOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetB2xUserFlowLanguageDefaultPagesCount - Get the number of the resource
func (c B2xUserFlowLanguageDefaultPageClient) GetB2xUserFlowLanguageDefaultPagesCount(ctx context.Context, id stable.IdentityB2xUserFlowIdLanguageId, options GetB2xUserFlowLanguageDefaultPagesCountOperationOptions) (result GetB2xUserFlowLanguageDefaultPagesCountOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "text/plain",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/defaultPages/$count", id.ID()),
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
