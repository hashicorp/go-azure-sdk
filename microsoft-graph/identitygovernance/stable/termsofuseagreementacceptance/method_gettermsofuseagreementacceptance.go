package termsofuseagreementacceptance

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTermsOfUseAgreementAcceptanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AgreementAcceptance
}

type GetTermsOfUseAgreementAcceptanceOperationOptions struct {
	Expand    *odata.Expand
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
	Select    *[]string
}

func DefaultGetTermsOfUseAgreementAcceptanceOperationOptions() GetTermsOfUseAgreementAcceptanceOperationOptions {
	return GetTermsOfUseAgreementAcceptanceOperationOptions{}
}

func (o GetTermsOfUseAgreementAcceptanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTermsOfUseAgreementAcceptanceOperationOptions) ToOData() *odata.Query {
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

func (o GetTermsOfUseAgreementAcceptanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTermsOfUseAgreementAcceptance - Get acceptances from identityGovernance. Read-only. Information about acceptances
// of this agreement.
func (c TermsOfUseAgreementAcceptanceClient) GetTermsOfUseAgreementAcceptance(ctx context.Context, id stable.IdentityGovernanceTermsOfUseAgreementIdAcceptanceId, options GetTermsOfUseAgreementAcceptanceOperationOptions) (result GetTermsOfUseAgreementAcceptanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model stable.AgreementAcceptance
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
