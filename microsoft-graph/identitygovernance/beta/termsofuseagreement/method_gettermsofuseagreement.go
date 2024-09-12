package termsofuseagreement

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTermsOfUseAgreementOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.Agreement
}

type GetTermsOfUseAgreementOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetTermsOfUseAgreementOperationOptions() GetTermsOfUseAgreementOperationOptions {
	return GetTermsOfUseAgreementOperationOptions{}
}

func (o GetTermsOfUseAgreementOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTermsOfUseAgreementOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTermsOfUseAgreementOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTermsOfUseAgreement - Get agreement. Retrieve the properties and relationships of an agreement object.
func (c TermsOfUseAgreementClient) GetTermsOfUseAgreement(ctx context.Context, id beta.IdentityGovernanceTermsOfUseAgreementId, options GetTermsOfUseAgreementOperationOptions) (result GetTermsOfUseAgreementOperationResponse, err error) {
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

	var model beta.Agreement
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
