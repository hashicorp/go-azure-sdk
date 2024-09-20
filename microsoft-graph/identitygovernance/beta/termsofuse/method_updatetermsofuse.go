package termsofuse

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTermsOfUseOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTermsOfUseOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateTermsOfUseOperationOptions() UpdateTermsOfUseOperationOptions {
	return UpdateTermsOfUseOperationOptions{}
}

func (o UpdateTermsOfUseOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTermsOfUseOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTermsOfUseOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTermsOfUse - Update the navigation property termsOfUse in identityGovernance
func (c TermsOfUseClient) UpdateTermsOfUse(ctx context.Context, input beta.TermsOfUseContainer, options UpdateTermsOfUseOperationOptions) (result UpdateTermsOfUseOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          "/identityGovernance/termsOfUse",
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
