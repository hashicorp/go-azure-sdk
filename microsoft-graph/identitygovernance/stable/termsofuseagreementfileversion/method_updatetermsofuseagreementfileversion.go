package termsofuseagreementfileversion

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateTermsOfUseAgreementFileVersionOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateTermsOfUseAgreementFileVersionOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdateTermsOfUseAgreementFileVersionOperationOptions() UpdateTermsOfUseAgreementFileVersionOperationOptions {
	return UpdateTermsOfUseAgreementFileVersionOperationOptions{}
}

func (o UpdateTermsOfUseAgreementFileVersionOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateTermsOfUseAgreementFileVersionOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateTermsOfUseAgreementFileVersionOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateTermsOfUseAgreementFileVersion - Update the navigation property versions in identityGovernance
func (c TermsOfUseAgreementFileVersionClient) UpdateTermsOfUseAgreementFileVersion(ctx context.Context, id stable.IdentityGovernanceTermsOfUseAgreementIdFileIdVersionId, input stable.AgreementFileVersion, options UpdateTermsOfUseAgreementFileVersionOperationOptions) (result UpdateTermsOfUseAgreementFileVersionOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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
